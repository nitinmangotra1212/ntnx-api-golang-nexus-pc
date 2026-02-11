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
  # Original 5 columns: item_id, item_name, item_type, description, ext_id
  # New GroupBy columns: quantity, price, is_active, priority, status
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
  metric_type_list {
    is_attribute: true
    metric_name: "quantity"
    entity_type_name: "item"
    user_metadata: "{\\"data_type\\":\\"int64\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "price"
    entity_type_name: "item"
    user_metadata: "{\\"data_type\\":\\"double\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "is_active"
    entity_type_name: "item"
    user_metadata: "{\\"data_type\\":\\"bool\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "priority"
    entity_type_name: "item"
    user_metadata: "{\\"data_type\\":\\"int64\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "status"
    entity_type_name: "item"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "int64_list"
    entity_type_name: "item"
    user_metadata: "{\\"data_type\\":\\"int64_list\\"}"
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
  import random
  uuid_list = []
  for _ in range(120):  # 100 TYPE1 + 20 TYPE2
    uuid_list.append(str(uuid.uuid4()))
  
  # Status values for variety
  status_values = ["pending", "active", "inactive"]
  
  # Create TYPE1 items (100 items)
  for i in range(100):
    arg = UpdateEntityArg()
    item_name = "test item " + str(i)
    description = "test item description " + str(i)
    item_uuid = uuid_list[i]
    
    # Generate random values for new columns
    quantity = random.randint(10, 200)
    price = round(random.uniform(10.0, 50.0), 1)
    is_active = random.choice([True, False])
    priority = random.randint(1, 5)
    status = random.choice(status_values)
    
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
    attribute_data_arg_list {
      attribute_data {
        name: "quantity"
        value {
          int64_value: ''' + str(quantity) + '''
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "price"
        value {
          double_value: ''' + str(price) + '''
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "is_active"
        value {
          bool_value: ''' + str(is_active).lower() + '''
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "priority"
        value {
          int64_value: ''' + str(priority) + '''
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "status"
        value {
          str_value: "''' + status + '''"
        }
      }
    }
    '''
    
    # Add list attributes with random values for TYPE1 items (only int64_list)
    int64_list = [random.randint(100, 200), random.randint(1, 50), random.randint(1, 100)]
    
    # Add int64_list
    query += '''
    attribute_data_arg_list {
      attribute_data {
        name: "int64_list"
        value {
          int64_list {
'''
    for val in int64_list:
      query += f'            value_list: {val}\n'
    query += '''          }
        }
      }
    }
    '''
    
    Merge(query, arg)
    
    try:
      ret = insights_interface.UpdateEntity(arg)
      if (i + 1) % 10 == 0:
        print(f"  ✅ Created {i + 1} TYPE1 items...")
    except InsightsInterfaceError as ex:
      print(f"❌ Error creating item {i + 1}: " + ex.message + "\n")
      print(ex.ret)
  
  # Create TYPE2 items (20 items) - for GroupBy testing
  print("\n📝 Creating TYPE2 items for GroupBy testing...")
  for i in range(20):
    arg = UpdateEntityArg()
    item_id = 200 + i + 1  # Start from 201
    item_name = "TYPE2 Item " + str(i + 1)
    description = "Description for TYPE2 item " + str(i + 1)
    item_uuid = uuid_list[100 + i]  # Use remaining UUIDs
    
    # Generate values for TYPE2 items
    quantity = 50 + (i * 5)
    price = round(20.0 + (i * 0.5), 1)
    is_active = (i % 2 == 0)  # Alternate true/false
    priority = 2  # Priority 2 for TYPE2
    status = status_values[i % 3]  # Rotate through statuses
    
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
          int64_value: ''' + str(item_id) + '''
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
          str_value: "TYPE2"
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
    attribute_data_arg_list {
      attribute_data {
        name: "quantity"
        value {
          int64_value: ''' + str(quantity) + '''
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "price"
        value {
          double_value: ''' + str(price) + '''
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "is_active"
        value {
          bool_value: ''' + str(is_active).lower() + '''
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "priority"
        value {
          int64_value: ''' + str(priority) + '''
        }
      }
    }
    attribute_data_arg_list {
      attribute_data {
        name: "status"
        value {
          str_value: "''' + status + '''"
        }
      }
    }
    '''
    
    # Add list attributes with random values for TYPE2 items (only int64_list)
    int64_list = [random.randint(100, 200), random.randint(1, 50), random.randint(1, 100)]
    
    # Add int64_list
    query += '''
    attribute_data_arg_list {
      attribute_data {
        name: "int64_list"
        value {
          int64_list {
'''
    for val in int64_list:
      query += f'            value_list: {val}\n'
    query += '''          }
        }
      }
    }
    '''
    
    Merge(query, arg)
    
    try:
      ret = insights_interface.UpdateEntity(arg)
      print(f"  ✅ Created TYPE2 item {i + 1}...")
    except InsightsInterfaceError as ex:
      print(f"❌ Error creating TYPE2 item {i + 1}: " + ex.message + "\n")
      print(ex.ret)
  
  print(f"\n✅ Successfully created 100 TYPE1 items + 20 TYPE2 items in IDF with all columns including list attributes!")
  print(f"💡 All items now have list attributes (int64List) with random values!")
  print(f"\n🔍 Access IDF data at: http://<PC_IP>:2027/")

def register_item_stats():
  """Register item_stats entity type and attributes in IDF"""
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  # Step 1: Register Entity Type
  arg = RegisterEntityTypesArg()
  query = '''
    entity_type_info_list {
      entity_type_name: "item_stats"
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
    print("✅ Successfully registered entity type: item_stats")
  except InsightsInterfaceError as ex:
    print("❌ Error registering entity type: " + ex.message + "\n")
    print(ex.ret)
    return
  
  # Step 2: Register Metric Types
  # Some fields are attributes (is_attribute: true) - single values stored with entity
  # Some fields are time-series metrics (is_attribute: false) - multiple timestamped values in TSDB
  arg = RegisterMetricTypesArg()
  query = '''
  metric_type_list {
    is_attribute: true
    metric_name: "stats_ext_id"
    entity_type_name: "item_stats"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "item_ext_id"
    entity_type_name: "item_stats"
    user_metadata: "{\\"data_type\\":\\"string\\"}"
  }
  metric_type_list {
    is_attribute: false
    metric_name: "age"
    entity_type_name: "item_stats"
    user_metadata: "{\\"data_type\\":\\"int64\\"}"
  }
  metric_type_list {
    is_attribute: false
    metric_name: "heart_rate"
    entity_type_name: "item_stats"
    user_metadata: "{\\"data_type\\":\\"int64\\"}"
  }
  metric_type_list {
    is_attribute: false
    metric_name: "food_intake"
    entity_type_name: "item_stats"
    user_metadata: "{\\"data_type\\":\\"double\\"}"
  }
  '''
  Merge(query, arg)
  
  try:
    ret = insights_interface.RegisterMetricTypes(arg)
    print("✅ Successfully registered metric types for item_stats")
  except InsightsInterfaceError as ex:
    print("❌ Error registering metric types: " + ex.message + "\n")
    print(ex.ret)
    return
  
  print("\n✅ Successfully registered item_stats entity type and metrics in IDF!")
  print("  📊 Attributes (is_attribute: true): stats_ext_id, item_ext_id")
  print("  📈 Time-series metrics (is_attribute: false): age, heart_rate, food_intake")
  print("  ℹ️  Note: To create item_stats records, use create_item_stats.py script with item extIds.")

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
  # Original 4 columns: item_id, entity_type, entity_id, count
  # New GroupBy columns: total_count, average_score
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
  metric_type_list {
    is_attribute: true
    metric_name: "total_count"
    entity_type_name: "item_associations"
    user_metadata: "{\\"data_type\\":\\"int64\\"}"
  }
  metric_type_list {
    is_attribute: true
    metric_name: "average_score"
    entity_type_name: "item_associations"
    user_metadata: "{\\"data_type\\":\\"double\\"}"
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
    register_item_stats()  # Registers entity type only; use create_item_stats.py to create records
    print("\n" + "="*60)
    register_file()

