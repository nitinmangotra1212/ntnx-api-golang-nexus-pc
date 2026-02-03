#!/usr/bin/env python

"""
Simple script to create sample item_stats records in IDF for testing $expand=itemStats functionality.
Run this AFTER running setup_nexus_idf.py to create items.

Usage:
  python3 create_item_stats.py
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
import random
import env
from google.protobuf.text_format import Merge
from insights_interface.insights_interface_pb2 import *
from insights_interface.insights_interface import *

def create_item_stats_for_items(item_ext_ids):
  """Create sample item_stats records for given item extIds"""
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  print(f"\n📝 Creating item_stats records for {len(item_ext_ids)} items...")
  stats_count = 0
  
  for item_ext_id in item_ext_ids:
    # Each item can have 1-3 stats records (to test one-to-many relationship)
    num_stats_records = random.randint(1, 3)
    
    for j in range(num_stats_records):
      stats_uuid = str(uuid.uuid4())
      
      # Generate random stats values
      age = random.randint(1, 10)
      heart_rate = random.randint(60, 100)
      food_intake = round(random.uniform(100.0, 500.0), 1)
      
      arg = UpdateEntityArg()
      query = '''
      entity_guid {
        entity_type_name: "item_stats"
        entity_id: "''' + stats_uuid + '''"
      }
      full_update: false
      attribute_data_arg_list {
        attribute_data {
          name: "stats_ext_id"
          value {
            str_value: "''' + stats_uuid + '''"
          }
        }
      }
      attribute_data_arg_list {
        attribute_data {
          name: "item_ext_id"
          value {
            str_value: "''' + item_ext_id + '''"
          }
        }
      }
      attribute_data_arg_list {
        attribute_data {
          name: "age"
          value {
            int64_value: ''' + str(age) + '''
          }
        }
      }
      attribute_data_arg_list {
        attribute_data {
          name: "heart_rate"
          value {
            int64_value: ''' + str(heart_rate) + '''
          }
        }
      }
      attribute_data_arg_list {
        attribute_data {
          name: "food_intake"
          value {
            double_value: ''' + str(food_intake) + '''
          }
        }
      }
      '''
      Merge(query, arg)
      
      try:
        ret = insights_interface.UpdateEntity(arg)
        stats_count += 1
      except InsightsInterfaceError as ex:
        print(f"❌ Error creating stats record for item {item_ext_id}: " + ex.message)
  
  print(f"  ✅ Created {stats_count} item_stats records!")
  return stats_count

if __name__ == "__main__":
  with warnings.catch_warnings():
    warnings.simplefilter("ignore")
    
    print("⚠️  This script creates item_stats records for item extIds.")
    print("   Replace the sample_ext_ids list with actual item extIds from your IDF.")
    print("   You can get item extIds by calling: GET /api/nexus/v4.1/config/items\n")
    
    # Item extIds - extract these from your API response
    # Call: GET /api/nexus/v4.1/config/items
    # Copy the 'extId' values from the response and add them here
    sample_ext_ids = [
      "7c8fa618-ab6f-40b9-b48f-760b1e5352b8",
      "10ee4d4b-3542-4677-9290-b637f03fb21b",
      "d8c2c5be-d544-49cc-ae73-15334ca74364",
      "9297d5ba-172e-4e52-8161-1806f64cc445",
      "c83420fe-d403-49f2-92ee-99c9b5371307",
      "7435ec8b-b73a-45bd-b1d6-b959f29a4420",
      "6f9a4c5a-a457-4bc0-bce0-8b6a3f8a0ebe",
      "e9330d7d-7b84-4bd7-8573-b0c6e5055870",
      "6840bfe6-840e-4721-923b-3e455334f617",
      "da867936-a16e-4529-ab6c-01a9ecca2051",
      "22c36414-ced4-40ec-86b6-f99b54bd1833",
      "d09cfa8d-7a0c-44b2-b867-d44d0987048c",
      "2a0b2573-e857-4628-9f03-27b460af2c2c",
    ]
    
    if not sample_ext_ids:
      print("❌ No item extIds provided!")
      print("\n💡 To get item extIds:")
      print("   1. Call the API: GET /api/nexus/v4.1/config/items")
      print("   2. Copy the 'extId' values from the response")
      print("   3. Add them to the 'sample_ext_ids' list in this script")
      print("   4. Run this script again")
      sys.exit(1)
    
    create_item_stats_for_items(sample_ext_ids)
    print("\n✅ Done! Now test the API with $expand=itemStats")

