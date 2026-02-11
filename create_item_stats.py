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
import time
import env
from google.protobuf.text_format import Merge
from insights_interface.insights_interface_pb2 import *
from insights_interface.insights_interface import *

def create_item_stats_for_items(item_ext_ids):
  """Create sample item_stats records for given item extIds
  
  Following the pattern from cat_stats:
  - Attributes (is_attribute: true): stats_ext_id, item_ext_id - set via UpdateEntity
  - Time-series metrics (is_attribute: false): age, heart_rate, food_intake - set via PutMetricData
  """
  insights_interface = InsightsInterface("127.0.0.1", "2027")
  
  print(f"\n📝 Creating item_stats records for {len(item_ext_ids)} items...")
  print("   Using time-series metrics (TSDB) for age, heart_rate, food_intake")
  stats_count = 0
  
  # Time-series data parameters
  # Create data over 30 days to allow testing with old timestamps
  # BUT ensure the latest data point is very recent (within last hour) so IDF returns it
  current_time_ms = int(time.time() * 1000)
  current_time_usecs = current_time_ms * 1000
  # Use last 30 days (wide range for testing)
  thirty_days_ms = 30 * 24 * 60 * 60 * 1000  # 30 days
  start_time_ms = current_time_ms - thirty_days_ms
  # Use 1 hour intervals to spread data over 30 days
  interval_ms = 60 * 60 * 1000  # 1 hour
  num_data_points = 30  # 30 time-series data points per metric (spread over 30 days)
  
  # Ensure the last data point is very recent (within last hour) so IDF returns it
  # Calculate what the last timestamp would be
  last_timestamp_ms = start_time_ms + ((num_data_points - 1) * interval_ms)
  if last_timestamp_ms < (current_time_ms - (60 * 60 * 1000)):  # If last point is > 1 hour ago
    # Adjust: make last point 30 minutes ago, and recalculate start_time
    end_time_ms = current_time_ms - (30 * 60 * 1000)  # 30 minutes ago
    start_time_ms = end_time_ms - ((num_data_points - 1) * interval_ms)
    print(f"   ⚠️  Adjusted time range to ensure latest data point is recent (30 minutes ago)")
  
  print(f"📊 Time-series parameters:")
  print(f"   Current time: {current_time_ms} ms ({time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(current_time_ms/1000))})")
  print(f"   Start time: {start_time_ms} ms ({time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(start_time_ms/1000))})")
  print(f"   Interval: {interval_ms} ms ({interval_ms/(60*60*1000)} hours)")
  print(f"   Data points per metric: {num_data_points}")
  print(f"   ⚠️  Creating data over last 30 days to allow testing with old timestamps")
  
  # Store entity info for batch metric data push
  entity_info_list = []
  
  for item_ext_id in item_ext_ids:
    # Each item can have 1-3 stats records (to test one-to-many relationship)
    num_stats_records = random.randint(1, 3)
    
    for j in range(num_stats_records):
      stats_uuid = str(uuid.uuid4())
      
      # Generate base values for time-series metrics
      age_base = random.randint(1, 10)
      heart_rate_base = random.randint(60, 100)
      food_intake_base = round(random.uniform(100.0, 500.0), 1)
      
      # Step 1: Create entity with ONLY attribute fields (stats_ext_id, item_ext_id)
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
      '''
      Merge(query, arg)
      
      try:
        ret = insights_interface.UpdateEntity(arg)
        stats_count += 1
        
        # Store entity info for metric data push
        entity_info_list.append({
          'entity_id': stats_uuid,
          'item_ext_id': item_ext_id,
          'age_base': age_base,
          'heart_rate_base': heart_rate_base,
          'food_intake_base': food_intake_base
        })
      except InsightsInterfaceError as ex:
        print(f"❌ Error creating stats entity for item {item_ext_id}: " + ex.message)
        continue
  
  # Step 2: Push time-series metrics using PutMetricData
  print(f"\n📊 Pushing time-series metric data to TSDB for {len(entity_info_list)} entities...")
  
  # Process entities in batches to avoid too large requests
  batch_size = 10  # entities per batch
  
  for batch_start in range(0, len(entity_info_list), batch_size):
    batch_end = min(batch_start + batch_size, len(entity_info_list))
    batch_entities = entity_info_list[batch_start:batch_end]
    
    metric_data_arg = PutMetricDataArg()
    
    for entity_info in batch_entities:
      # Add entity with its metrics
      ewm = metric_data_arg.entity_with_metric_list.add()
      ewm.entity_guid.entity_type_name = "item_stats"
      ewm.entity_guid.entity_id = entity_info['entity_id']
      
      # Prepare time-series data for all metrics
      age_metric = ewm.metric_data_list.add()
      age_metric.name = "age"
      
      heart_rate_metric = ewm.metric_data_list.add()
      heart_rate_metric.name = "heart_rate"
      
      food_intake_metric = ewm.metric_data_list.add()
      food_intake_metric.name = "food_intake"
      
      # Add time-series data points
      # Spread data over 30 days (from start_time_ms to current_time_ms)
      for data_point_idx in range(num_data_points):
        # Calculate timestamp: spread evenly over the time range
        # Start from oldest (start_time_ms) and move forward
        time_offset_ms = data_point_idx * interval_ms
        timestamp_ms = start_time_ms + time_offset_ms
        
        # Ensure timestamp doesn't exceed current time (safety check)
        if timestamp_ms > current_time_ms:
          # If we've exceeded current time, use current time minus remaining intervals
          remaining_points = num_data_points - data_point_idx
          timestamp_ms = current_time_ms - (remaining_points * interval_ms)
          if timestamp_ms < start_time_ms:
            timestamp_ms = start_time_ms
        
        timestamp_usecs = timestamp_ms * 1000
        
        # Debug: Log first entity's first and last data points
        if entity_info == batch_entities[0] and data_point_idx == 0:
          first_timestamp_str = time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(timestamp_ms/1000))
          print(f"    📅 First data point: {first_timestamp_str} ({timestamp_ms} ms)")
        if entity_info == batch_entities[0] and data_point_idx == num_data_points - 1:
          last_timestamp_str = time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(timestamp_ms/1000))
          print(f"    📅 Last data point: {last_timestamp_str} ({timestamp_ms} ms)")
          print(f"    ⚠️  Data spans 30 days to allow testing with old timestamps")
        
        # Age value (slight variation over time)
        age_val = age_metric.value_list.add()
        age_val.timestamp_usecs = timestamp_usecs
        age_val.value.int64_value = entity_info['age_base'] + random.randint(-1, 1)
        
        # Heart rate value (random variation for each time point)
        heart_rate_val = heart_rate_metric.value_list.add()
        heart_rate_val.timestamp_usecs = timestamp_usecs
        heart_rate_val.value.int64_value = entity_info['heart_rate_base'] + random.randint(-10, 10)
        
        # Food intake value (random variation for each time point)
        food_intake_val = food_intake_metric.value_list.add()
        food_intake_val.timestamp_usecs = timestamp_usecs
        food_intake_val.value.double_value = round(entity_info['food_intake_base'] + random.uniform(-50.0, 50.0), 1)
    
    # Push batch of entity metrics
    # Debug: Log what we're about to push
    total_values = 0
    for ewm in metric_data_arg.entity_with_metric_list:
      for metric in ewm.metric_data_list:
        total_values += len(metric.value_list)
    print(f"  📤 Pushing batch {batch_start+1}-{batch_end}: {len(metric_data_arg.entity_with_metric_list)} entities, {total_values} total metric values")
    
    try:
      ret = insights_interface.PutMetricData(metric_data_arg)
      print(f"  ✅ PutMetricData call succeeded for entities {batch_start+1}-{batch_end}")
      if ret:
        print(f"     Response: rpc_execution_time_usecs: {ret.rpc_execution_time_usecs if hasattr(ret, 'rpc_execution_time_usecs') else 'N/A'}")
      
      # Wait a moment for TSDB to process each batch
      # IDF needs time to index time-series metrics before they're queryable
      wait_time = 2 if batch_start == 0 else 1  # Longer wait for first batch
      print(f"  ⏳ Waiting {wait_time} second(s) for TSDB to index metrics...")
      time.sleep(wait_time)
    except InsightsInterfaceError as ex:
      print(f"  ❌ Error pushing metric data: {ex.message}")
      if hasattr(ex, 'ret'):
        print(f"     Error response: {ex.ret}")
      import traceback
      traceback.print_exc()
      # Continue with next batch even if this one fails
      continue
    except Exception as ex:
      print(f"  ❌ Unexpected error pushing metric data: {ex}")
      import traceback
      traceback.print_exc()
      # Continue with next batch even if this one fails
      continue
  
  print(f"\n  ✅ Created {stats_count} item_stats entities with TSDB time-series data!")
  print(f"     Each entity has {num_data_points} time-series data points per metric (age, heart_rate, food_intake)")
  print(f"\n  ⏳ Waiting 10 seconds for IDF to fully index the time-series metrics...")
  print(f"     IDF needs time to index metrics before they're queryable via GetEntitiesWithMetricsRet")
  print(f"     This is especially important for time-series metrics (TSDB)")
  time.sleep(10)
  print(f"  ✅ Done! Data should now be queryable.")
  print(f"\n  💡 Next steps:")
  print(f"     1. Run: python3 verify_item_stats_metrics.py")
  print(f"     2. Check if ValueCount > 0 for age, heart_rate, food_intake")
  print(f"     3. If ValueCount=0, wait another 10-30 seconds and try again")
  print(f"     4. Test API: curl -k 'https://10.114.163.60:9440/api/nexus/v4.1/config/items?$expand=itemStats&$limit=5'")
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
      "3fd49ec8-eff3-45cf-96bc-37228e5c3d4d",
      "e75e47c5-5015-4cca-8557-18a9bb5fe4ef",
      "14e608e7-8868-4253-a9ec-fe8e84eeadeb",
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

