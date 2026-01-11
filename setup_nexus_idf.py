#!/usr/bin/env python

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

# Note: The default IP and port is the localhost IDF instance.
# You can configure it accordingly.

def register_item():
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  # Step 1: Register Entity Type
  arg = RegisterEntityTypesArg()
  query = '''
    entity_type_info_list {
      entity_type_name: "item"
      type_info {
        suppress_replication : true
        backup_replication_control {
          is_backup_required: false
        }
        replication_control_list {
          enable_replication_from : kNDFS
        }
        is_evictable: false
        retain_attributes_on_deletion: false
        track_attribute_changes : false
        enable_pulsehd_collection : false
      }
    }
  '''
  Merge(query, arg)
  
  try:
    ret = insights_interface.RegisterEntityTypes(arg)
    print("✅ Successfully registered entity type: item")
  except InsightsInterfaceError as ex:
    print("❌ Error registering entity type: " + ex.message + "\n")
    print(ex.ret)
    return
  
  # Step 2: Register Metric Types (Attributes)
  # Only 5 columns: item_id, item_name, item_type, description, ext_id
  arg = RegisterMetricTypesArg()
  query = '''
  metric_type_list {
    is_attribute: true
    metric_name: "item_id"
    entity_type_name: "item"
    user_metadata: "{\\"data_type\\":\\"int64\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "item_name"
    entity_type_name: "item"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "item_type"
    entity_type_name: "item"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "description"
    entity_type_name: "item"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "ext_id"
    entity_type_name: "item"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  '''
  Merge(query, arg)
  
  try:
    ret = insights_interface.RegisterMetricTypes(arg)
    print("✅ Successfully registered metric types (attributes)")
  except InsightsInterfaceError as ex:
    print("❌ Error registering metric types: " + ex.message + "\n")
    print(ex.ret)
    return
  
  # Step 3: Create sample data (optional - for testing)
  print("\n📝 Creating sample items in IDF...")
  uuid_list = []
  for _ in range(110):
    uuid_list.append(str(uuid.uuid4()))
  
  for i in range(110):
    arg = UpdateEntityArg()
    item_name = "test item " + str(i)
    description = "test item description " + str(i)
    item_uuid = uuid_list[i]
    
    query = '''
    entity_guid {
      entity_type_name: "item"
      entity_id: "''' + item_uuid + '''"
    }
    full_update: false
    attribute_data_arg_list {
      attribute_data {
        name: "item_id"
        value {
          int64_value: ''' + str(i + 1) + '''
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "item_name"
        value {
          str_value: "''' + item_name + '''"
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "item_type"
        value {
          str_value: "TYPE1"
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "description"
        value {
          str_value: "''' + description + '''"
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "ext_id"
        value {
          str_value: "''' + item_uuid + '''"
        }
      }
    }
    '''
    
    Merge(query, arg)
    
    try:
      ret = insights_interface.UpdateEntity(arg)
      if (i + 1) % 10 == 0:
        print(f"  ✅ Created {i + 1} items...")
    except InsightsInterfaceError as ex:
      print(f"❌ Error creating item {i + 1}: " + ex.message + "\n")
      print(ex.ret)
  
  print(f"\n✅ Successfully created {110} sample items in IDF!")
  print(f"\n🔍 Access IDF data at: http://<PC_IP>:2027/")

def register_item_associations():
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  # Step 1: Register Entity Type
  arg = RegisterEntityTypesArg()
  query = '''
    entity_type_info_list {
      entity_type_name: "item_associations"
      type_info {
        suppress_replication : true
        backup_replication_control {
          is_backup_required: false
        }
        replication_control_list {
          enable_replication_from : kNDFS
        }
        is_evictable: false
        retain_attributes_on_deletion: false
        track_attribute_changes : false
        enable_pulsehd_collection : false
      }
    }
  '''
  Merge(query, arg)
  
  try:
    ret = insights_interface.RegisterEntityTypes(arg)
    print("✅ Successfully registered entity type: item_associations")
  except InsightsInterfaceError as ex:
    print("❌ Error registering entity type: " + ex.message + "\n")
    print(ex.ret)
    return
  
  # Step 2: Register Metric Types (Attributes)
  # 4 columns: item_id, entity_type, entity_id, count
  arg = RegisterMetricTypesArg()
  query = '''
  metric_type_list {
    is_attribute: true
    metric_name: "item_id"
    entity_type_name: "item_associations"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "entity_type"
    entity_type_name: "item_associations"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "entity_id"
    entity_type_name: "item_associations"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "count"
    entity_type_name: "item_associations"
    user_metadata: "{\\"data_type\\":\\"int64\\"}"
  }
  '''
  Merge(query, arg)
  
  try:
    ret = insights_interface.RegisterMetricTypes(arg)
    print("✅ Successfully registered metric types (attributes) for item_associations")
  except InsightsInterfaceError as ex:
    print("❌ Error registering metric types: " + ex.message + "\n")
    print(ex.ret)
    return
  
  # Step 3: Create sample associations (optional - for testing)
  print("\n📝 Creating sample associations in IDF...")
  # Note: We'll create associations for items that were just created in register_item()
  # Since we can't easily query IDF from Python, we'll create associations for a few sample item extIds
  # In production, associations would be created by other services
  
  # For testing, create associations for the first 10 items
  # We'll use a simple approach: create associations with predictable item extIds
  # In real usage, you'd get these from IDF or from the items you just created
  
  print("  ℹ️  Note: Creating sample associations for testing.")
  print("     In production, associations are created by other services.")
  print("     For now, we'll skip automatic association creation.")
  print("     You can manually create associations using UpdateEntity with:")
  print("       - entity_type_name: 'item_associations'")
  print("       - attributes: item_id, entity_type, entity_id, count")
  
  # Skip automatic creation for now - user can create associations manually if needed
  print(f"\n✅ Successfully registered item_associations entity type in IDF!")

def register_file():
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  # Step 1: Register Entity Type
  arg = RegisterEntityTypesArg()
  query = '''
    entity_type_info_list {
      entity_type_name: "file"
      type_info {
        suppress_replication : true
        backup_replication_control {
          is_backup_required: false
        }
        replication_control_list {
          enable_replication_from : kNDFS
        }
        is_evictable: false
        retain_attributes_on_deletion: false
        track_attribute_changes : false
        enable_pulsehd_collection : false
      }
    }
  '''
  Merge(query, arg)
  
  try:
    ret = insights_interface.RegisterEntityTypes(arg)
    print("✅ Successfully registered entity type: file")
  except InsightsInterfaceError as ex:
    print("❌ Error registering entity type: " + ex.message + "\n")
    print(ex.ret)
    return
  
  # Step 2: Register Metric Types (Attributes)
  # Attributes: ext_id, file_name, file_size, content_type
  arg = RegisterMetricTypesArg()
  query = '''
  metric_type_list {
    is_attribute: true
    metric_name: "ext_id"
    entity_type_name: "file"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "file_name"
    entity_type_name: "file"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "file_size"
    entity_type_name: "file"
    user_metadata: "{\\"data_type\\":\\"int64\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "content_type"
    entity_type_name: "file"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  '''
  Merge(query, arg)
  
  try:
    ret = insights_interface.RegisterMetricTypes(arg)
    print("✅ Successfully registered metric types (attributes) for file")
  except InsightsInterfaceError as ex:
    print("❌ Error registering metric types: " + ex.message + "\n")
    print(ex.ret)
    return
  
  print(f"\n✅ Successfully registered file entity type in IDF!")

if __name__ == "__main__":
  with warnings.catch_warnings():
    warnings.simplefilter("ignore")
    register_item()
    print("\n" + "="*60)
    register_item_associations()
    print("\n" + "="*60)
    register_file()

