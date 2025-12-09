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
from google.protobuf.text_format import Merge
from insights_interface.insights_interface_pb2 import *
from insights_interface.insights_interface import *

def delete_all_items():
  """Delete all items from IDF"""
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  print("\n🗑️  Deleting all items from IDF...")
  
  # Use IDF HTTP API to get entity IDs (simpler than query builder)
  import urllib.request
  import json
  
  try:
    # Query IDF web UI API to get all item entities
    url = "http://127.0.0.1:2027/entities?type=item"
    with urllib.request.urlopen(url) as response:
      html = response.read().decode('utf-8')
      # Parse HTML to extract entity IDs (IDF web UI returns HTML table)
      # Entity IDs are in the table rows
      import re
      # Find all entity IDs in the HTML (they appear as links or in table cells)
      # Pattern: entity_id appears in URLs like /entities?type=item&id=<uuid>
      entity_ids = re.findall(r'entity_id=([a-f0-9\-]{36})', html)
      # Also try finding UUIDs directly in the page
      if not entity_ids:
        entity_ids = re.findall(r'([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})', html, re.IGNORECASE)
      
      if not entity_ids:
        print("  ℹ️  No items found in IDF (or couldn't parse entity IDs from web UI)")
        return 0
      
      print(f"  📋 Found {len(entity_ids)} items to delete...")
      
  except Exception as ex:
    print(f"  ⚠️  Error querying IDF web UI: {ex}")
    print(f"  💡 Falling back to manual deletion - check IDF web UI")
    return 0
  
  try:
    item_count = 0
    for entity_id in entity_ids:
      # Delete entity - need to get entity first to get CAS value
      get_arg = GetEntitiesArg()
      get_query = '''
      entity_guid_list {
        entity_type_name: "item"
        entity_id: "''' + entity_id + '''"
      }
      '''
      Merge(get_query, get_arg)
      
      try:
        # Try to get entity first to get CAS value
        try:
          get_ret = insights_interface.GetEntity(get_arg)
          # GetEntitiesRet has entity list
          if hasattr(get_ret, 'entity') and len(get_ret.entity) > 0:
            entity = get_ret.entity[0]
            cas_value = entity.cas_value if hasattr(entity, 'cas_value') and entity.cas_value else 1
          else:
            cas_value = 1  # Default CAS value
        except:
          # If GetEntity fails, try with default CAS
          cas_value = 1
        
        # Delete with CAS value
        delete_arg = DeleteEntityArg()
        delete_query = '''
        entity_guid {
          entity_type_name: "item"
          entity_id: "''' + entity_id + '''"
        }
        cas_value: ''' + str(cas_value) + '''
        '''
        Merge(delete_query, delete_arg)
        
        insights_interface.DeleteEntity(delete_arg)
        item_count += 1
        if item_count % 10 == 0:
          print(f"  ✅ Deleted {item_count} items...")
      except InsightsInterfaceError as ex:
        # Entity might already be deleted, continue
        if "not found" not in ex.message.lower() and "does not exist" not in ex.message.lower() and "incorrect cas" not in ex.message.lower():
          # Try with CAS = 1 if CAS error
          if "incorrect cas" in ex.message.lower():
            try:
              delete_arg = DeleteEntityArg()
              delete_query = '''
              entity_guid {
                entity_type_name: "item"
                entity_id: "''' + entity_id + '''"
              }
              cas_value: 1
              '''
              Merge(delete_query, delete_arg)
              insights_interface.DeleteEntity(delete_arg)
              item_count += 1
            except:
              pass  # Skip if still fails
          else:
            print(f"  ⚠️  Error deleting item {entity_id}: {ex.message}")
    
    print(f"  ✅ Deleted {item_count} items from IDF")
    return item_count
  except Exception as ex:
    print(f"  ⚠️  Error during deletion: {ex}")
    return item_count if 'item_count' in locals() else 0

def delete_all_associations():
  """Delete all item_associations from IDF"""
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  print("\n🗑️  Deleting all item_associations from IDF...")
  
  # Use IDF HTTP API to get entity IDs (simpler than query builder)
  import urllib.request
  import re
  
  try:
    # Query IDF web UI API to get all association entities
    url = "http://127.0.0.1:2027/entities?type=item_associations"
    with urllib.request.urlopen(url) as response:
      html = response.read().decode('utf-8')
      # Parse HTML to extract entity IDs
      entity_ids = re.findall(r'entity_id=([a-f0-9\-]{36})', html)
      if not entity_ids:
        entity_ids = re.findall(r'([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})', html, re.IGNORECASE)
      
      if not entity_ids:
        print("  ℹ️  No associations found in IDF (or couldn't parse entity IDs from web UI)")
        return 0
      
      print(f"  📋 Found {len(entity_ids)} associations to delete...")
      
  except Exception as ex:
    print(f"  ⚠️  Error querying IDF web UI: {ex}")
    return 0
  
  try:
    assoc_count = 0
    for entity_id in entity_ids:
      # Delete entity - need to get entity first to get CAS value
      get_arg = GetEntitiesArg()
      get_query = '''
      entity_guid_list {
        entity_type_name: "item_associations"
        entity_id: "''' + entity_id + '''"
      }
      '''
      Merge(get_query, get_arg)
      
      try:
        # Try to get entity first to get CAS value
        try:
          get_ret = insights_interface.GetEntity(get_arg)
          # GetEntitiesRet has entity list
          if hasattr(get_ret, 'entity') and len(get_ret.entity) > 0:
            entity = get_ret.entity[0]
            cas_value = entity.cas_value if hasattr(entity, 'cas_value') and entity.cas_value else 1
          else:
            cas_value = 1  # Default CAS value
        except:
          # If GetEntity fails, try with default CAS
          cas_value = 1
        
        # Delete with CAS value
        delete_arg = DeleteEntityArg()
        delete_query = '''
        entity_guid {
          entity_type_name: "item_associations"
          entity_id: "''' + entity_id + '''"
        }
        cas_value: ''' + str(cas_value) + '''
        '''
        Merge(delete_query, delete_arg)
        
        insights_interface.DeleteEntity(delete_arg)
        assoc_count += 1
        if assoc_count % 10 == 0:
          print(f"  ✅ Deleted {assoc_count} associations...")
      except InsightsInterfaceError as ex:
        # Entity might already be deleted, continue
        if "not found" not in ex.message.lower() and "does not exist" not in ex.message.lower() and "incorrect cas" not in ex.message.lower():
          # Try with CAS = 1 if CAS error
          if "incorrect cas" in ex.message.lower():
            try:
              delete_arg = DeleteEntityArg()
              delete_query = '''
              entity_guid {
                entity_type_name: "item_associations"
                entity_id: "''' + entity_id + '''"
              }
              cas_value: 1
              '''
              Merge(delete_query, delete_arg)
              insights_interface.DeleteEntity(delete_arg)
              assoc_count += 1
            except:
              pass  # Skip if still fails
          else:
            print(f"  ⚠️  Error deleting association {entity_id}: {ex.message}")
    
    print(f"  ✅ Deleted {assoc_count} associations from IDF")
    return assoc_count
  except Exception as ex:
    print(f"  ⚠️  Error during deletion: {ex}")
    return assoc_count if 'assoc_count' in locals() else 0

if __name__ == "__main__":
  with warnings.catch_warnings():
    warnings.simplefilter("ignore")
    
    print("="*60)
    print("🧹 IDF Cleanup Script")
    print("="*60)
    print("\n⚠️  This will delete ALL items and associations from IDF!")
    print("   Press Ctrl+C to cancel, or wait 5 seconds to continue...")
    
    import time
    time.sleep(5)
    
    # Delete associations first (they reference items)
    assoc_count = delete_all_associations()
    
    # Then delete items
    item_count = delete_all_items()
    
    print("\n" + "="*60)
    print(f"✅ Cleanup complete!")
    print(f"   Deleted {item_count} items")
    print(f"   Deleted {assoc_count} associations")
    print("="*60)
    print("\n💡 Next steps:")
    print("   1. Run setup_nexus_idf.py to create fresh data")
    print("   2. Run create_associations.py to create associations")

