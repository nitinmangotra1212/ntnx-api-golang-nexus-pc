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
      "987f4bee-4144-46bc-88a9-0d3d12f2e34e",
      "1581ca6d-5ee8-4fd2-b4f4-1dbf4130aaaa",
      "615e9e10-42f5-4253-80fa-d824aeb410e4",
      "c8dc5c4c-911d-4f03-8f79-79a888212d13",
      "c4c6571c-4ff9-4d7c-b00d-ff4d1d69b3de",
      "059d64d7-2d20-4bac-99e8-9973d8af8539",
      "05aafecf-4a87-44b2-b252-b454db3124df",
      "abdb4146-7afd-4cea-97fa-7fb888937702",
      "37580f9c-0132-4ed2-88a8-71e7c95fb1fa",
      "f0eb604d-c41c-4ffc-8356-e3ef7fc72d71",
      "6f31eec4-4fa5-4e60-83f3-83935469d80f",
      "8cb1a935-61e2-4872-82e7-7e969633060a",
      "44e7ae41-bc7d-4e45-b1a2-8d95464ae3f8",
      "c8aa2fe6-a737-47cd-b7d1-0a88404b1c68",
      "652b9fd3-e8ba-449a-9854-78a5d9077065",
      "2576e2a6-d64f-4617-b533-487fa8258ee4",
      "8b45bc54-314f-4571-9da1-6c246247f2e3",
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

