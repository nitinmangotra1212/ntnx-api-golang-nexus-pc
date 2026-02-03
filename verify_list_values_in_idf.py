#!/usr/bin/env python

"""
Script to verify if list values exist in IDF for items.
Run this to check if the list values were successfully added.

Usage:
  python3 verify_list_values_in_idf.py
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

def verify_list_values_in_item(item_ext_id):
  """Verify if list values exist for an item"""
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  # Query the item
  arg = GetEntitiesArg()
  query = '''
  entity_guid_list {
    entity_type_name: "item"
    entity_id: "''' + item_ext_id + '''"
  }
  '''
  Merge(query, arg)
  
  try:
    # Use GetEntities method (available in InsightsInterface)
    ret = insights_interface.GetEntities(arg)
    
    # Check different possible response structures
    entities = None
    if hasattr(ret, 'entity') and ret.entity:
      entities = ret.entity
    elif hasattr(ret, 'entities') and ret.entities:
      entities = ret.entities
    elif hasattr(ret, 'entity_list') and ret.entity_list:
      entities = ret.entity_list
    else:
      print(f"  ❌ Item {item_ext_id} not found in IDF (no entities in response)")
      return False
    
    if len(entities) == 0:
      print(f"  ❌ Item {item_ext_id} not found in IDF")
      return False
    
    entity = entities[0]
    print(f"\n📋 Item {item_ext_id}:")
    
    # Check for list attributes
    has_string_list = False
    has_int64_list = False
    has_float_list = False
    has_bool_list = False
    has_byte_list = False
    has_enum_list = False
    
    # Check different possible attribute access patterns
    attrs = None
    if hasattr(entity, 'attribute_data_map') and entity.attribute_data_map:
      attrs = entity.attribute_data_map
    elif hasattr(entity, 'AttributeDataMap') and entity.AttributeDataMap:
      attrs = entity.AttributeDataMap
    else:
      print(f"  ⚠️  No attributes found for item {item_ext_id}")
      return False
    
    for attr in attrs:
      # Handle different attribute name access patterns
      attr_name = attr.name if hasattr(attr, 'name') else (attr.Name if hasattr(attr, 'Name') else None)
      attr_value = attr.value if hasattr(attr, 'value') else (attr.Value if hasattr(attr, 'Value') else None)
      
      if attr_name == "string_list":
        has_string_list = True
        if attr_value and hasattr(attr_value, 'HasField') and attr_value.HasField("str_list"):
          values = list(attr_value.str_list.value_list)
          print(f"  ✅ string_list: {values}")
        elif attr_value and hasattr(attr_value, 'str_list') and attr_value.str_list:
          values = list(attr_value.str_list.value_list)
          print(f"  ✅ string_list: {values}")
        else:
          value_type = attr_value.WhichOneof('value_type') if hasattr(attr_value, 'WhichOneof') else "unknown"
          print(f"  ⚠️  string_list exists but value type is: {value_type}")
      
      elif attr_name == "int64_list":
        has_int64_list = True
        if attr_value and hasattr(attr_value, 'HasField') and attr_value.HasField("int64_list"):
          values = list(attr_value.int64_list.value_list)
          print(f"  ✅ int64_list: {values}")
        elif attr_value and hasattr(attr_value, 'int64_list') and attr_value.int64_list:
          values = list(attr_value.int64_list.value_list)
          print(f"  ✅ int64_list: {values}")
        else:
          value_type = attr_value.WhichOneof('value_type') if hasattr(attr_value, 'WhichOneof') else "unknown"
          print(f"  ⚠️  int64_list exists but value type is: {value_type}")
      
      elif attr_name == "float_list":
        has_float_list = True
        if attr_value and hasattr(attr_value, 'HasField') and attr_value.HasField("double_list"):
          values = list(attr_value.double_list.value_list)
          print(f"  ✅ float_list: {values}")
        elif attr_value and hasattr(attr_value, 'double_list') and attr_value.double_list:
          values = list(attr_value.double_list.value_list)
          print(f"  ✅ float_list: {values}")
        else:
          value_type = attr_value.WhichOneof('value_type') if hasattr(attr_value, 'WhichOneof') else "unknown"
          print(f"  ⚠️  float_list exists but value type is: {value_type}")
      
      elif attr_name == "bool_list":
        has_bool_list = True
        if attr_value and hasattr(attr_value, 'HasField') and attr_value.HasField("bool_list"):
          values = list(attr_value.bool_list.value_list)
          print(f"  ✅ bool_list: {values}")
        elif attr_value and hasattr(attr_value, 'bool_list') and attr_value.bool_list:
          values = list(attr_value.bool_list.value_list)
          print(f"  ✅ bool_list: {values}")
        else:
          value_type = attr_value.WhichOneof('value_type') if hasattr(attr_value, 'WhichOneof') else "unknown"
          print(f"  ⚠️  bool_list exists but value type is: {value_type}")
      
      elif attr_name == "byte_list":
        has_byte_list = True
        if attr_value and hasattr(attr_value, 'HasField') and attr_value.HasField("int64_list"):
          values = list(attr_value.int64_list.value_list)
          print(f"  ✅ byte_list: {values}")
        elif attr_value and hasattr(attr_value, 'int64_list') and attr_value.int64_list:
          values = list(attr_value.int64_list.value_list)
          print(f"  ✅ byte_list: {values}")
        else:
          value_type = attr_value.WhichOneof('value_type') if hasattr(attr_value, 'WhichOneof') else "unknown"
          print(f"  ⚠️  byte_list exists but value type is: {value_type}")
      
      elif attr_name == "enum_list":
        has_enum_list = True
        if attr_value and hasattr(attr_value, 'HasField') and attr_value.HasField("str_list"):
          values = list(attr_value.str_list.value_list)
          print(f"  ✅ enum_list: {values}")
        elif attr_value and hasattr(attr_value, 'str_list') and attr_value.str_list:
          values = list(attr_value.str_list.value_list)
          print(f"  ✅ enum_list: {values}")
        else:
          value_type = attr_value.WhichOneof('value_type') if hasattr(attr_value, 'WhichOneof') else "unknown"
          print(f"  ⚠️  enum_list exists but value type is: {value_type}")
    
    # Report missing lists
    if not has_string_list:
      print(f"  ❌ string_list: NOT FOUND")
    if not has_int64_list:
      print(f"  ❌ int64_list: NOT FOUND")
    if not has_float_list:
      print(f"  ❌ float_list: NOT FOUND")
    if not has_bool_list:
      print(f"  ❌ bool_list: NOT FOUND")
    if not has_byte_list:
      print(f"  ❌ byte_list: NOT FOUND")
    if not has_enum_list:
      print(f"  ❌ enum_list: NOT FOUND")
    
    return has_string_list or has_int64_list or has_float_list or has_bool_list or has_byte_list or has_enum_list
    
  except InsightsInterfaceError as ex:
    print(f"❌ Error querying item {item_ext_id}: " + ex.message)
    return False

if __name__ == "__main__":
  with warnings.catch_warnings():
    warnings.simplefilter("ignore")
    
    # Parse gflags to avoid UnparsedFlagAccessError
    try:
      FLAGS(sys.argv)
    except gflags.FlagsError as e:
      print('%s\nUsage: %s ARGS\n%s' % (e, sys.argv[0], FLAGS))
      sys.exit(1)
    
    print("🔍 Verifying list values in IDF for items...\n")
    
    # Sample item extIds - same as in add_list_values_to_items.py
    sample_ext_ids = [
      "7c8fa618-ab6f-40b9-b48f-760b1e5352b8",
      "10ee4d4b-3542-4677-9290-b637f03fb21b",
      "d8c2c5be-d544-49cc-ae73-15334ca74364",
    ]
    
    found_count = 0
    for ext_id in sample_ext_ids:
      if verify_list_values_in_item(ext_id):
        found_count += 1
    
    print(f"\n✅ Found list values in {found_count} out of {len(sample_ext_ids)} items")
    if found_count == 0:
      print("\n💡 If no list values found:")
      print("   1. Run add_list_values_to_items.py to add list values")
      print("   2. Make sure the list attributes are registered in IDF (run setup_nexus_idf.py)")

