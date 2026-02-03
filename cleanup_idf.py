#!/usr/bin/env python

"""
Cleanup script to delete all items and associations from IDF
Run this before running setup_nexus_idf.py to start fresh
"""

import os
VIRTUALENV_PATH = "/home/nutanix/cluster/.venv/bin/bin/python3.9"

if os.path.exists(VIRTUALENV_PATH):
  if os.environ.get("PYTHON_TARGET_VERSION") is None:
    os.environ["PYTHON_TARGET_VERSION"] = "3.9"
  if os.environ.get("PYTHON_TARGET_PATH") is None:
    os.environ["PYTHON_TARGET_PATH"] = VIRTUALENV_PATH

import sys
sys.path.append("/home/nutanix/cluster/bin")

import warnings
import env
import gflags
from google.protobuf.text_format import Merge
from insights_interface.insights_interface_pb2 import *
from insights_interface.insights_interface import *

FLAGS = gflags.FLAGS

def delete_all_items():
  """Delete all items from IDF using QueryEntities to get entity IDs"""
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  print("\n🗑️  Deleting all items from IDF...")
  
  # Query IDF to get all item entities
  try:
    # Build query to get all items using Merge (text format)
    query_arg = GetEntitiesWithMetricsArg()
    query_text = '''
    entity_type_name: "item"
    query {
      group_by {
        raw_columns {
          column: "ext_id"
        }
        raw_columns {
          column: "item_id"
        }
        raw_limit {
          limit: 10000
          offset: 0
        }
      }
    }
    '''
    Merge(query_text, query_arg)
    
    # Query IDF
    query_ret = insights_interface.GetEntitiesWithMetricsRet(query_arg)
    
    if not query_ret or not query_ret.GroupResultsList:
      print("  ℹ️  No items found in IDF")
      return 0
    
    # Extract entity IDs from query results
    entity_ids = []
    for group_result in query_ret.GroupResultsList:
      for entity_with_metric in group_result.RawResults:
        entity = entity_with_metric.Entity
        if entity and entity.EntityGuid:
          entity_id = entity.EntityGuid.EntityId
          if entity_id:
            entity_ids.append(entity_id)
    
    if not entity_ids:
      print("  ℹ️  No items found in IDF")
      return 0
    
    print(f"  📋 Found {len(entity_ids)} items to delete...")
    
  except Exception as ex:
    print(f"  ⚠️  Error querying IDF: {ex}")
    print(f"  💡 Tip: IDF entities might not be deletable if they're non-CAS")
    return 0
  
  # Try to delete each entity
  item_count = 0
  skipped_count = 0
  
  for entity_id in entity_ids:
    try:
      # For non-CAS entities, try deleting without CAS value first
      delete_arg = DeleteEntityArg()
      delete_query = '''
      entity_guid {
        entity_type_name: "item"
        entity_id: "''' + entity_id + '''"
      }
      '''
      Merge(delete_query, delete_arg)
      
      # Try delete without CAS (for non-CAS entities)
      insights_interface.DeleteEntity(delete_arg)
      item_count += 1
      if item_count % 10 == 0:
        print(f"  ✅ Deleted {item_count} items...")
        
    except InsightsInterfaceError as ex:
      error_msg = ex.message.lower()
      # Skip if entity not found or doesn't exist
      if "not found" in error_msg or "does not exist" in error_msg:
        skipped_count += 1
        continue
      # Skip if non-CAS entity error (these can't be deleted)
      if "kcasupdatefornoncasentity" in error_msg or "noncas" in error_msg or "cas" in error_msg:
        skipped_count += 1
        if skipped_count <= 3:  # Only show first few warnings
          print(f"  ⚠️  Skipping non-CAS entity {entity_id[:8]}... (cannot be deleted)")
        continue
      # Other errors
      if skipped_count <= 3:
        print(f"  ⚠️  Error deleting item {entity_id[:8]}...: {ex.message}")
      skipped_count += 1
  
  print(f"  ✅ Deleted {item_count} items from IDF")
  if skipped_count > 0:
    print(f"  ⚠️  Skipped {skipped_count} items (non-CAS entities cannot be deleted)")
    print(f"  💡 Tip: Non-CAS entities will be overwritten when you run setup_nexus_idf.py")
  
  return item_count

def delete_all_associations():
  """Delete all item_associations from IDF using QueryEntities"""
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  print("\n🗑️  Deleting all item_associations from IDF...")
  
  # Query IDF to get all association entities
  try:
    # Build query to get all associations using Merge (text format)
    query_arg = GetEntitiesWithMetricsArg()
    query_text = '''
    entity_type_name: "item_associations"
    query {
      group_by {
        raw_columns {
          column: "item_id"
        }
        raw_columns {
          column: "entity_type"
        }
        raw_limit {
          limit: 10000
          offset: 0
        }
      }
    }
    '''
    Merge(query_text, query_arg)
    
    # Query IDF
    query_ret = insights_interface.GetEntitiesWithMetricsRet(query_arg)
    
    if not query_ret or not query_ret.GroupResultsList:
      print("  ℹ️  No associations found in IDF")
      return 0
    
    # Extract entity IDs from query results
    entity_ids = []
    for group_result in query_ret.GroupResultsList:
      for entity_with_metric in group_result.RawResults:
        entity = entity_with_metric.Entity
        if entity and entity.EntityGuid:
          entity_id = entity.EntityGuid.EntityId
          if entity_id:
            entity_ids.append(entity_id)
    
    if not entity_ids:
      print("  ℹ️  No associations found in IDF")
      return 0
    
    print(f"  📋 Found {len(entity_ids)} associations to delete...")
    
  except Exception as ex:
    print(f"  ⚠️  Error querying IDF: {ex}")
    return 0
  
  # Try to delete each association
  assoc_count = 0
  skipped_count = 0
  
  for entity_id in entity_ids:
    try:
      # For non-CAS entities, try deleting without CAS value
      delete_arg = DeleteEntityArg()
      delete_query = '''
      entity_guid {
        entity_type_name: "item_associations"
        entity_id: "''' + entity_id + '''"
      }
      '''
      Merge(delete_query, delete_arg)
      
      insights_interface.DeleteEntity(delete_arg)
      assoc_count += 1
      if assoc_count % 10 == 0:
        print(f"  ✅ Deleted {assoc_count} associations...")
        
    except InsightsInterfaceError as ex:
      error_msg = ex.message.lower()
      # Skip if entity not found or doesn't exist
      if "not found" in error_msg or "does not exist" in error_msg:
        skipped_count += 1
        continue
      # Skip if non-CAS entity error
      if "kcasupdatefornoncasentity" in error_msg or "noncas" in error_msg or "cas" in error_msg:
        skipped_count += 1
        continue
      # Other errors
      if skipped_count <= 3:
        print(f"  ⚠️  Error deleting association {entity_id[:8]}...: {ex.message}")
      skipped_count += 1
  
  print(f"  ✅ Deleted {assoc_count} associations from IDF")
  if skipped_count > 0:
    print(f"  ⚠️  Skipped {skipped_count} associations (non-CAS entities cannot be deleted)")
  
  return assoc_count

if __name__ == "__main__":
  with warnings.catch_warnings():
    warnings.simplefilter("ignore")
    
    # Parse gflags to avoid UnparsedFlagAccessError
    try:
      FLAGS(sys.argv)
    except gflags.FlagsError as e:
      print('%s\nUsage: %s ARGS\n%s' % (e, sys.argv[0], FLAGS))
      sys.exit(1)
    
    print("="*60)
    print("🧹 IDF Cleanup Script")
    print("="*60)
    print("\n⚠️  This will delete ALL items and associations from IDF!")
    print("   Press Ctrl+C to cancel, or wait 5 seconds to continue...")
    
    import time
    time.sleep(5)
    
    # Delete associations first (they reference items)
    assoc_count = delete_all_associations()
    
    # Delete item_stats (they reference items)
    print("\n🗑️  Deleting all item_stats from IDF...")
    stats_count = 0
    insights_interface = InsightsInterface("127.0.0.1", "2027")
    try:
      query_arg = GetEntitiesWithMetricsArg()
      query_text = '''
      entity_type_name: "item_stats"
      query {
        group_by {
          raw_columns {
            column: "stats_ext_id"
          }
          raw_limit {
            limit: 10000
            offset: 0
          }
        }
      }
      '''
      Merge(query_text, query_arg)
      
      query_ret = insights_interface.GetEntitiesWithMetricsRet(query_arg)
      if query_ret and query_ret.GroupResultsList:
        for group_result in query_ret.GroupResultsList:
          for entity_with_metric in group_result.RawResults:
            entity = entity_with_metric.Entity
            if entity and entity.EntityGuid:
              entity_id = entity.EntityGuid.EntityId
              if entity_id:
                try:
                  delete_arg = DeleteEntityArg()
                  delete_query = f'''
                  entity_guid {{
                    entity_type_name: "item_stats"
                    entity_id: "{entity_id}"
                  }}
                  '''
                  Merge(delete_query, delete_arg)
                  insights_interface.DeleteEntity(delete_arg)
                  stats_count += 1
                except:
                  pass
      print(f"  ✅ Deleted {stats_count} item_stats from IDF")
    except Exception as ex:
      print(f"  ⚠️  Error deleting item_stats: {ex}")
    
    # Then delete items
    item_count = delete_all_items()
    
    print("\n" + "="*60)
    print(f"✅ Cleanup complete!")
    print(f"   Deleted {item_count} items")
    print(f"   Deleted {assoc_count} associations")
    print(f"   Deleted {stats_count} item_stats")
    print("="*60)
    print("\n💡 Next steps:")
    print("   1. Run setup_nexus_idf.py to create fresh data with list attributes")
    print("   2. Run create_associations.py to create associations (optional)")
    print("   3. Run create_item_stats.py to create item stats (optional)")

