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
    
    # Sample item extIds - extracted from your API response
    # These are the actual extIds from your items in IDF
    sample_ext_ids = [
      "62253556-2475-4f52-84d8-5338cd7895d1",  # item 31
      "d29433cb-f731-44fa-b1be-7ee20cf5de7b",  # item 67
      "c506a692-3879-4d30-af99-7c517b6383e7",  # item 61
      "92dc28f6-6939-4266-8765-55181d8d8594",  # item 6
      "6ca16ea2-da98-40b7-8a1e-a2249e2ff2d2",  # item 89
      "2a7ac459-aae5-4c21-84ce-1360253bf7c4",  # item 35
      "4782ddd4-8f7b-4927-bd8b-d96f34fe03a8",  # item 30
      "6f454278-ddd9-4df7-a19a-3e078e649699",  # item 79
      "318c1eda-41e6-4871-ae08-637a6c31d2fd",  # item 82
      "704379d2-eb68-47a4-b4be-5f8267c1850c",  # item 4
      "3fcc3a72-6ccc-43fc-8070-fd4957262ce7",  # item 1
      "e546c48f-83c8-4715-9bb0-1ee7f73412cc",  # item 62
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

