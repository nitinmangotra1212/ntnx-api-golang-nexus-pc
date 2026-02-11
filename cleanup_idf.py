#!/usr/bin/env python

"""
Delete Script for IDF Entities
This script deletes all entity instances (data) of type: item, item_associations, and item_stats
Use this to clean up IDF entity data before re-registering entities

Note: This deletes entity DATA, not entity type definitions.
Entity type definitions cannot be deleted through the IDF API.
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
gflags.FLAGS([])

from google.protobuf.text_format import Merge
from insights_interface.insights_interface_pb2 import (
    BatchDeleteEntitiesArg, DeleteEntityArg, GetEntitiesWithMetricsArg, EntityGuid
)
from insights_interface.insights_interface import InsightsInterface

# Query template for fetching entities
QUERY_TEMPLATE = """
query {{
  entity_list {{
    entity_type_name: "{entity_type}"
  }}
  where_clause {{
    lhs {{
      comparison_expr {{
        lhs {{
          leaf {{
            column: "_master_cluster_uuid_"
          }}
        }}
        operator: kExists
      }}
    }}
    operator: kNot
  }}
  group_by {{
    raw_limit {{
      limit: 1000
    }}
    raw_columns {{
      column: "_cas_value_"
    }}
  }}
}}
"""

class DeleteEntities(object):
    """
    Class to delete all entity instances of a given type
    """

    def __init__(self, entity_type):
        """
        Initialize the deletion class
        
        Args:
            entity_type (str): Entity type to delete instances of (e.g., "item", "item_associations")
        """
        self.insights = InsightsInterface("127.0.0.1", "2027")
        self.entity_type = entity_type
        self.entity_data_list = None

    @staticmethod
    def delete_entity_arg(entity_type, entity_id, cas_value):
        """
        Create a DeleteEntityArg for a given entity
        
        Args:
            entity_type (str): Entity type name
            entity_id (str): Entity id
            cas_value (int/bool): Current cas value of the entity
            
        Returns:
            DeleteEntityArg: Arg for delete RPC
        """
        guid = EntityGuid()
        guid.entity_type_name = entity_type
        guid.entity_id = entity_id
        dea = DeleteEntityArg()
        dea.entity_guid.CopyFrom(guid)
        if cas_value is not False:
            dea.cas_value = cas_value + 1
        return dea

    def get_entity_data(self):
        """
        Get list of entities of the given type
        
        Returns:
            None (sets self.entity_data_list)
        """
        self.entity_data_list = []
        arg = GetEntitiesWithMetricsArg()
        
        # Build query string for this entity type
        query_str = QUERY_TEMPLATE.format(entity_type=self.entity_type)
        
        # Merge the query into the arg
        Merge(query_str, arg)
        
        try:
            ret = self.insights.GetEntitiesWithMetrics(arg)
            
            if ret.total_group_count == 0:
                return
            
            # Extract entity IDs and cas values
            for each in ret.group_results_list[0].raw_results:
                e_id = each.entity_guid.entity_id
                
                # Get cas value if available
                if len(each.metric_data_list) > 0 and len(each.metric_data_list[0].value_list) == 1:
                    cas_value = each.metric_data_list[0].value_list[0].value.uint64_value
                else:
                    cas_value = False
                    
                self.entity_data_list.append((e_id, cas_value))
                
        except Exception as ex:
            print(f"  Note: No {self.entity_type} entities found or entity type doesn't exist yet")
            self.entity_data_list = []

    def batch_delete(self):
        """
        Batch delete list of entity instances
        """
        arg = BatchDeleteEntitiesArg()
        delete_args = []
        
        for e_id, cas_value in self.entity_data_list:
            delete_args.append(self.delete_entity_arg(self.entity_type, e_id, cas_value))
        
        arg.entity_list.extend(delete_args)
        print(f"  Deleting {len(delete_args)} {self.entity_type} entities...")
        
        try:
            self.insights.BatchDeleteEntities(arg)
            print(f"  ✓ Successfully deleted {len(delete_args)} entities")
        except Exception as error:
            print(f"  Warning: {error}")

    def run(self):
        """
        Method to delete all entity instances of given type
        
        Returns:
            int: Total number of entities deleted
        """
        total_deleted = 0
        
        while True:
            self.get_entity_data()
            
            if not self.entity_data_list:
                break
            
            batch_size = len(self.entity_data_list)
            self.batch_delete()
            total_deleted += batch_size
        
        return total_deleted


def main():
    """Main function to delete all entity instances"""
    print("\n" + "="*60)
    print("IDF Entity Data Deletion Script")
    print("="*60)
    print("This script will delete ALL entity instances (data) of:")
    print("  - item_associations")
    print("  - item_stats")
    print("  - item")
    print()
    print("NOTE: This deletes entity DATA only.")
    print("      Entity type definitions cannot be deleted via IDF API.")
    print("="*60)
    
    # Ask for confirmation
    confirmation = input("\nAre you sure you want to proceed? (yes/no): ")
    if confirmation.lower() not in ['yes', 'y']:
        print("Deletion cancelled.")
        return
    
    total_deleted = 0
    
    # Delete entity instances in order (children first, then parent)
    # Order matters because of relationships
    entity_types = [
        "item_associations",  # Delete associations first (child of item)
        "item_stats",         # Delete stats second (child of item)
        "item"                # Delete items last (parent)
    ]
    
    for entity_type in entity_types:
        print(f"\n{'='*60}")
        print(f"Processing: {entity_type}")
        print(f"{'='*60}")
        
        deleter = DeleteEntities(entity_type)
        count = deleter.run()
        total_deleted += count
        
        if count > 0:
            print(f"  Total {entity_type} entities deleted: {count}")
        else:
            print(f"  No {entity_type} entities to delete")
    
    print(f"\n{'='*60}")
    print(f"SUMMARY")
    print(f"{'='*60}")
    print(f"Total entity instances deleted: {total_deleted}")
    print(f"{'='*60}\n")
    
    if total_deleted > 0:
        print("✓ Cleanup complete! Entity data has been removed.")
        print("  (Entity type definitions remain in the schema)")
        print("\n💡 Next steps:")
        print("   1. Run setup_nexus_idf.py to register entity types and create fresh data")
        print("   2. Run create_associations.py to create associations (optional)")
        print("   3. Run create_item_stats.py to create item stats with time-series data (optional)")
    else:
        print("✓ No entity instances found (clean state)")


if __name__ == "__main__":
    with warnings.catch_warnings():
        warnings.simplefilter("ignore")
        main()
