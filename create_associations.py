#!/usr/bin/env python

"""
Simple script to create sample associations in IDF for testing $expand functionality.
Run this AFTER running setup_nexus_idf.py to create items.

Usage:
  python3 create_associations.py
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
import uuid
import env
from google.protobuf.text_format import Merge
from insights_interface.insights_interface_pb2 import *
from insights_interface.insights_interface import *

def create_associations_for_items(item_ext_ids):
  """Create sample associations for given item extIds"""
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  print(f"\n📝 Creating associations for {len(item_ext_ids)} items...")
  association_count = 0
  
  for item_ext_id in item_ext_ids:
    # Create 2 associations per item
    for i in range(2):
      arg = UpdateEntityArg()
      assoc_uuid = str(uuid.uuid4())
      entity_type = "vm" if i == 0 else "host"
      entity_id = str(uuid.uuid4())
      count = (i + 1) * 5
      
      query = '''
      entity_guid {
        entity_type_name: "item_associations"
        entity_id: "''' + assoc_uuid + '''"
      }
      full_update: false
      attribute_data_arg_list {
        attribute_data {
          name: "item_id"
          value {
            str_value: "''' + item_ext_id + '''"
          }
        }
      }
      attribute_data_arg_list {
        attribute_data {
          name: "entity_type"
          value {
            str_value: "''' + entity_type + '''"
          }
        }
      }
      attribute_data_arg_list {
        attribute_data {
          name: "entity_id"
          value {
            str_value: "''' + entity_id + '''"
          }
        }
      }
      attribute_data_arg_list {
        attribute_data {
          name: "count"
          value {
            int64_value: ''' + str(count) + '''
          }
        }
      }
      '''
      
      Merge(query, arg)
      
      try:
        ret = insights_interface.UpdateEntity(arg)
        association_count += 1
      except InsightsInterfaceError as ex:
        print(f"❌ Error creating association: " + ex.message)
  
  print(f"  ✅ Created {association_count} associations!")
  return association_count

if __name__ == "__main__":
  with warnings.catch_warnings():
    warnings.simplefilter("ignore")
    
    # For testing, create associations for some sample item extIds
    # In production, you'd get these from IDF or from your item creation process
    # For now, we'll use a few sample UUIDs - replace these with actual item extIds from your IDF
    
    print("⚠️  This script creates associations for sample item extIds.")
    print("   Replace the sample_ext_ids list with actual item extIds from your IDF.")
    print("   You can get item extIds by querying IDF or from your item creation logs.\n")
    
    # Item extIds - extracted from your API response
    # These are the actual extIds from your items in IDF
    sample_ext_ids = [
      "7480f4b3-e4c0-4595-928c-d39d21f90780",
      "69c61ef2-848c-4277-9d55-15229eedd0e5",
      "5837ba55-08e4-495a-80bd-d5a428a30d7d",
      "050e203e-d080-4eef-ba83-5ffd43dc3f1c",
      "c791b044-fbce-441c-bb13-c82c54781406",
      "06b0086e-d83e-46e5-b1e6-197793c9e2b5",
      "fd953777-1359-481c-bac7-cdcbef6a5903",
      "a4eeca37-1a49-4611-82de-865f07a4a202",
      "b22a96b3-80b5-4e2a-a671-947f7c7c201c",
      "76ac21d1-1b61-432f-84f5-4fb327f03fcd",
      "12a7e430-79e6-4cde-bd84-4897608b61f2",
      "f8284d6f-3f00-4ec6-8acd-e05cbbc70a23",
      "1c97914a-0953-48f0-b880-1b007d2072b1",
      "be4b2959-6b8f-4a4c-844f-9503f360b5d4",
      "21c8f004-f085-4c18-a321-2f62275c3518",
      "1f65cae5-c351-48e1-9383-a0b81e975dba",
      "0ece8f56-8622-405e-a832-495d5f4d8e3b",
    ]
    
    if not sample_ext_ids:
      print("❌ No item extIds provided!")
      print("\n💡 To get item extIds:")
      print("   1. Call the API: GET /api/nexus/v4.1/config/items")
      print("   2. Copy the 'extId' values from the response")
      print("   3. Add them to the 'sample_ext_ids' list in this script")
      print("   4. Run this script again")
      sys.exit(1)
    
    create_associations_for_items(sample_ext_ids)
    print("\n✅ Done! Now test the API with $expand=associations")

