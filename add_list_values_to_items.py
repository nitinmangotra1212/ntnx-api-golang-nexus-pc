#!/usr/bin/env python

"""
Script to add list values to existing items in IDF for testing.
Run this AFTER running setup_nexus_idf.py to create items.

Usage:
  python3 add_list_values_to_items.py
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
import random
import env
from google.protobuf.text_format import Merge
from insights_interface.insights_interface_pb2 import *
from insights_interface.insights_interface import *

def add_list_values_to_item(item_ext_id):
  """Add list values to an item"""
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  # Generate random list values
  string_list = ["value1", "value2", "value3"]
  int64_list = [1, 2, 3, 4, 5]
  float_list = [1.5, 2.7, 3.14, 4.9]
  bool_list = [True, False, True]
  byte_list = [1, 2, 3, 4, 5]
  enum_list = ["ACTIVE", "INACTIVE", "PENDING"]
  
  arg = UpdateEntityArg()
  query = '''
  entity_guid {
    entity_type_name: "item"
    entity_id: "''' + item_ext_id + '''"
  }
  full_update: false
  attribute_data_arg_list {
    attribute_data {
      name: "string_list"
      value {
        str_list {
          value_list: "value1"
          value_list: "value2"
          value_list: "value3"
        }
      }
    }
  }
  attribute_data_arg_list {
    attribute_data {
      name: "int64_list"
      value {
        int64_list {
          value_list: 1
          value_list: 2
          value_list: 3
          value_list: 4
          value_list: 5
        }
      }
    }
  }
  attribute_data_arg_list {
    attribute_data {
      name: "float_list"
      value {
        double_list {
          value_list: 1.5
          value_list: 2.7
          value_list: 3.14
          value_list: 4.9
        }
      }
    }
  }
  attribute_data_arg_list {
    attribute_data {
      name: "bool_list"
      value {
        bool_list {
          value_list: true
          value_list: false
          value_list: true
        }
      }
    }
  }
  attribute_data_arg_list {
    attribute_data {
      name: "byte_list"
      value {
        int64_list {
          value_list: 1
          value_list: 2
          value_list: 3
          value_list: 4
          value_list: 5
        }
      }
    }
  }
  attribute_data_arg_list {
    attribute_data {
      name: "enum_list"
      value {
        str_list {
          value_list: "ACTIVE"
          value_list: "INACTIVE"
          value_list: "PENDING"
        }
      }
    }
  }
  '''
  
  Merge(query, arg)
  
  try:
    ret = insights_interface.UpdateEntity(arg)
    print(f"  ✅ Added list values to item {item_ext_id}")
    return True
  except InsightsInterfaceError as ex:
    print(f"❌ Error adding list values to item {item_ext_id}: " + ex.message)
    return False

if __name__ == "__main__":
  with warnings.catch_warnings():
    warnings.simplefilter("ignore")
    
    print("⚠️  This script adds list values to existing items in IDF.")
    print("   It will update the first 5 items with sample list values.\n")
    
    # Sample item extIds - replace with actual extIds from your API response
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
    
    success_count = 0
    for ext_id in sample_ext_ids:
      if add_list_values_to_item(ext_id):
        success_count += 1
    
    print(f"\n✅ Done! Added list values to {success_count} items.")
    print("   Now test the API: GET /api/nexus/v4.1/config/items")
    print("   The list fields should appear in the response.")


