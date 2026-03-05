#!/usr/bin/env python

"""
Single script to register and populate 'item' and 'item_stats' tables in IDF.
Creates 120 items and 120 item_stats with a strict one-to-one mapping.

Usage:
  python3 register_create_item_idf.py
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
import time
import env
import gflags
gflags.FLAGS([])
from google.protobuf.text_format import Merge
from insights_interface.insights_interface_pb2 import *
from insights_interface.insights_interface import *

# Global variable to store item UUIDs for stats registration
item_uuid_list = []


def register_item():
  """Register 'item' entity type, its metric types, and populate 120 items."""
  global item_uuid_list
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
    print("Registered 'item' entity type")
  except InsightsInterfaceError as ex:
    print("Error registering item entity type: " + ex.message + "\n")
    print(ex.ret)

  # Step 2: Register Metric Types (Attributes)
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
    user_metadata: "{\\"data_type\\":\\"int64\\"}"
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
    print("Registered 'item' metric types (attributes)")
  except InsightsInterfaceError as ex:
    print("Error registering item metric types: " + ex.message + "\n")
    print(ex.ret)

  # Step 3: Populate 120 items (100 TYPE1 + 20 TYPE2)
  # Use deterministic UUIDs so each run overwrites the same entities (no stale data)
  print("\nCreating 120 items in IDF (100 TYPE1 + 20 TYPE2)...")
  uuid_list = [str(uuid.uuid5(uuid.NAMESPACE_DNS, f"nexus-item-{i}")) for i in range(120)]
  print(f"  Deterministic item UUIDs (first 3): {uuid_list[0]}, {uuid_list[1]}, {uuid_list[2]}")
  status_values = ["pending", "active", "inactive"]

  # Create 100 TYPE1 items
  for i in range(100):
    arg = UpdateEntityArg()
    item_name = "test item " + str(i)
    description = "test item description " + str(i)
    item_uuid = uuid_list[i]
    quantity = random.randint(10, 200)
    price = round(random.uniform(10.0, 50.0), 1)
    is_active = random.choice([True, False])
    priority = random.randint(1, 5)
    status = random.choice(status_values)
    int64_list = [random.randint(100, 200), random.randint(1, 50), random.randint(1, 100)]

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
          int64_value: 1
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
      if (i + 1) % 20 == 0:
        print(f"  Created {i + 1}/100 TYPE1 items...")
    except InsightsInterfaceError as ex:
      print(f"Error creating TYPE1 item {i + 1}: " + ex.message)

  # Create 20 TYPE2 items
  print("  Creating TYPE2 items...")
  for i in range(20):
    arg = UpdateEntityArg()
    item_id = 200 + i + 1
    item_name = "TYPE2 Item " + str(i + 1)
    description = "Description for TYPE2 item " + str(i + 1)
    item_uuid = uuid_list[100 + i]
    quantity = 50 + (i * 5)
    price = round(20.0 + (i * 0.5), 1)
    is_active = (i % 2 == 0)
    priority = 2
    status = status_values[i % 3]
    int64_list = [random.randint(100, 200), random.randint(1, 50), random.randint(1, 100)]

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
          int64_value: 2
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
        print(f"  Created {i + 1}/20 TYPE2 items...")
    except InsightsInterfaceError as ex:
      print(f"Error creating TYPE2 item {i + 1}: " + ex.message)

  item_uuid_list = uuid_list
  print(f"Successfully created 120 items (100 TYPE1 + 20 TYPE2)")
  return uuid_list


def register_item_stats():
  """Register 'item_stats' entity type and create time-series data (one-to-one mapping with items)."""
  global item_uuid_list
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
    print("Registered 'item_stats' entity type")
  except InsightsInterfaceError as ex:
    print("Error registering item_stats entity type: " + ex.message + "\n")
    print(ex.ret)

  # Step 2: Register Metric Types
  # Attributes (is_attribute: true): item_ext_id (foreign key for JOIN)
  # Time-series metrics (is_attribute: false): age, heart_rate, food_intake
  arg = RegisterMetricTypesArg()
  query = '''
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
    print("Registered 'item_stats' metric types")
  except InsightsInterfaceError as ex:
    print("Error registering item_stats metric types: " + ex.message + "\n")
    print(ex.ret)

  # Step 3: Create ONE item_stats entity per item (one-to-one mapping)
  current_time_ms = int(time.time() * 1000)
  thirty_days_ms = 30 * 24 * 60 * 60 * 1000
  start_time_ms = current_time_ms - thirty_days_ms
  interval_ms = 30 * 60 * 1000  # 30 minutes

  num_items_with_stats = len(item_uuid_list)  # All 120 items
  num_data_points = 100  # 100 time-series data points per metric

  print(f"\nCreating ONE item_stats entity per item (one-to-one mapping)")
  print(f"Total item_stats entities to create: {num_items_with_stats}")
  print(f"Each entity will have {num_data_points} time-series data points per metric (TSDB storage)")
  sample_stats_uuid = str(uuid.uuid5(uuid.NAMESPACE_DNS, "nexus-item-stats-0"))
  print(f"  Deterministic item_stats UUID[0]: {sample_stats_uuid}")
  print(f"  Time range: {start_time_ms}ms to {current_time_ms}ms")

  entity_info_list = []

  for item_idx in range(num_items_with_stats):
    item_ext_id = item_uuid_list[item_idx]
    # Deterministic UUID — same entity_id every run so old data is overwritten
    stats_uuid = str(uuid.uuid5(uuid.NAMESPACE_DNS, f"nexus-item-stats-{item_idx}"))
    age_value = random.randint(1, 10)

    entity_info_list.append({
      'entity_id': stats_uuid,
      'item_ext_id': item_ext_id,
      'age': age_value,
    })

    # Create entity with attribute field (item_ext_id only — JOIN key)
    # full_update: true to remove old timestamp/speed attributes from previous runs
    arg = UpdateEntityArg()
    query = '''
    entity_guid {
      entity_type_name: "item_stats"
      entity_id: "''' + stats_uuid + '''"
    }
    full_update: true
    attribute_data_arg_list {
      attribute_data {
        name: "item_ext_id"
        value {
          str_value: "''' + item_ext_id + '''"
        }
      }
    }
    '''
    Merge(query, arg)
    try:
      ret = insights_interface.UpdateEntity(arg)
      if (item_idx + 1) % 20 == 0:
        print(f"  Created {item_idx + 1}/{num_items_with_stats} item_stats entities...")
    except InsightsInterfaceError as ex:
      print(f"Error creating item_stats entity: {ex.message}")
      continue

  # Step 4: Push time-series metrics using PutMetricData (in batches)
  print(f"\nPushing time-series metric data to TSDB...")

  batch_size = 10
  for batch_start in range(0, len(entity_info_list), batch_size):
    batch_end = min(batch_start + batch_size, len(entity_info_list))
    batch_entities = entity_info_list[batch_start:batch_end]

    metric_data_arg = PutMetricDataArg()

    for entity_info in batch_entities:
      ewm = metric_data_arg.entity_with_metric_list.add()
      ewm.entity_guid.entity_type_name = "item_stats"
      ewm.entity_guid.entity_id = entity_info['entity_id']

      age_metric = ewm.metric_data_list.add()
      age_metric.name = "age"

      heart_rate_metric = ewm.metric_data_list.add()
      heart_rate_metric.name = "heart_rate"

      food_intake_metric = ewm.metric_data_list.add()
      food_intake_metric.name = "food_intake"

      for data_point_idx in range(num_data_points):
        timestamp_ms = start_time_ms + (data_point_idx * interval_ms * (1440 // num_data_points))
        timestamp_usecs = timestamp_ms * 1000

        age_val = age_metric.value_list.add()
        age_val.timestamp_usecs = timestamp_usecs
        age_val.value.int64_value = entity_info['age']

        heart_rate_val = heart_rate_metric.value_list.add()
        heart_rate_val.timestamp_usecs = timestamp_usecs
        heart_rate_val.value.int64_value = 60 + random.randint(0, 60)

        food_intake_val = food_intake_metric.value_list.add()
        food_intake_val.timestamp_usecs = timestamp_usecs
        food_intake_val.value.double_value = round(100 + random.random() * 400, 2)

    try:
      ret = insights_interface.PutMetricData(metric_data_arg)
      print(f"  Pushed metrics for entities {batch_start + 1}-{batch_end}/{len(entity_info_list)}")
    except InsightsInterfaceError as ex:
      print(f"Error pushing metric data: {ex.message}")

  print(f"Successfully created {num_items_with_stats} item_stats entities with TSDB time-series data")


if __name__ == "__main__":
  with warnings.catch_warnings():
    warnings.simplefilter("ignore")
    print("=" * 60)
    print("Starting IDF registration for item + item_stats...")
    print("=" * 60)

    print("\n--- Registering and populating 'item' table ---")
    register_item()
    print("Item registration complete")

    print("\n--- Registering and populating 'item_stats' table ---")
    register_item_stats()
    print("Item stats registration complete")

    print("\n" + "=" * 60)
    print("All registrations complete!")
    print(f"  - 120 items created (100 TYPE1 + 20 TYPE2)")
    print(f"  - 120 item_stats created (1-to-1 mapping with items)")
    print(f"  - item_stats linked via 'item_ext_id' -> item 'ext_id'")
    print("=" * 60)
