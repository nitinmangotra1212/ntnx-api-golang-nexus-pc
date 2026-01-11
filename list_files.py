#!/usr/bin/env python
"""
Script to list all uploaded files with their real filenames from IDF
Run this on PC: python3 list_files.py
"""

import os
import sys

# Setup path for PC environment (same as setup_nexus_idf.py)
VIRTUALENV_PATH = "/home/nutanix/cluster/.venv/bin/bin/python3.9"
if os.path.exists(VIRTUALENV_PATH):
    if os.environ.get("PYTHON_TARGET_VERSION") is None:
        os.environ["PYTHON_TARGET_VERSION"] = "3.9"
    if os.environ.get("PYTHON_TARGET_PATH") is None:
        os.environ["PYTHON_TARGET_PATH"] = VIRTUALENV_PATH

sys.path.append("/home/nutanix/cluster/bin")

import warnings
import env
from google.protobuf.text_format import Merge
from insights_interface.insights_interface_pb2 import *
from insights_interface.insights_interface import *

def list_files():
    """List all files from IDF with their metadata"""
    insights_interface = InsightsInterface("127.0.0.1", "2027")
    
    # Since constructing complex IDF queries in Python protobuf text format is difficult,
    # we'll use a simpler approach: list files from filesystem and query IDF for each
    # OR try to use a basic query structure
    
    # Let's try a simpler approach: get files from filesystem and query IDF metadata for each
    import glob
    import os
    
    downloads_path = "/home/nutanix/Downloads"
    files = []
    
    # Get all files from Downloads directory (they're named by extId)
    if os.path.exists(downloads_path):
        for filepath in glob.glob(os.path.join(downloads_path, "*")):
            if os.path.isfile(filepath):
                ext_id = os.path.basename(filepath)
                # Skip if it doesn't look like a UUID
                if len(ext_id) == 36 and ext_id.count('-') == 4:
                    files.append(ext_id)
    
    if not files:
        print("No files found in /home/nutanix/Downloads")
        return
    
    print(f"\n{'ExtId':<40} {'Filename':<50} {'Size':<20} {'Content Type':<20}")
    print("=" * 130)
    
    # Query IDF for metadata of each file
    for ext_id in sorted(files, reverse=True):
        get_arg = GetEntitiesArg()
        query_text = f'''
        entity_guid_list {{
          entity_type_name: "file"
          entity_id: "{ext_id}"
        }}
        '''
        
        try:
            Merge(query_text, get_arg)
            # Try to get entity from IDF
            # The method name might vary - try common variations
            ret = None
            error_msg = None
            
            # Use GetEntities method (confirmed available from debug output)
            ret = insights_interface.GetEntities(get_arg)
            
            # Extract metadata
            filename = ext_id  # Default to extId
            file_size = "N/A"
            content_type = "N/A"
            
            if ret:
                # Check different possible response structures
                entities = None
                if hasattr(ret, 'entity'):
                    entities = ret.entity
                elif hasattr(ret, 'entities'):
                    entities = ret.entities
                elif hasattr(ret, 'GetEntity'):
                    entities = ret.GetEntity()
                
                if entities:
                    for entity in entities:
                        if hasattr(entity, 'attribute_data_map'):
                            for attr in entity.attribute_data_map:
                                attr_name = attr.name
                                if not attr_name:
                                    continue
                                
                                if attr_name == "file_name" and attr.value and attr.value.str_value:
                                    filename = attr.value.str_value
                                elif attr_name == "file_size" and attr.value:
                                    file_size = str(attr.value.int64_value)
                                elif attr_name == "content_type" and attr.value and attr.value.str_value:
                                    content_type = attr.value.str_value
            
            # Get actual file size from filesystem if not in IDF
            file_path = os.path.join(downloads_path, ext_id)
            if os.path.exists(file_path):
                actual_size = os.path.getsize(file_path)
                if file_size == "N/A":
                    file_size = str(actual_size)
            
            # Format file size
            if file_size != "N/A":
                try:
                    size_int = int(file_size)
                    if size_int > 1024 * 1024:
                        file_size = f"{size_int / (1024*1024):.2f} MB"
                    elif size_int > 1024:
                        file_size = f"{size_int / 1024:.2f} KB"
                    else:
                        file_size = f"{size_int} bytes"
                except:
                    pass
            
            print(f"{ext_id:<40} {filename:<50} {file_size:<20} {content_type:<20}")
            
        except Exception as e:
            # If IDF query fails, still show the file with filesystem info
            file_path = os.path.join(downloads_path, ext_id)
            if os.path.exists(file_path):
                actual_size = os.path.getsize(file_path)
                size_str = f"{actual_size / (1024*1024):.2f} MB" if actual_size > 1024*1024 else f"{actual_size / 1024:.2f} KB"
                # Show error in content type column for debugging (first time only)
                error_display = "N/A" if error_msg else "N/A (IDF query failed)"
                print(f"{ext_id:<40} {ext_id:<50} {size_str:<20} {error_display:<20}")
    
    print(f"\nTotal files: {len(files)}")
    return
    
    try:
        ret = insights_interface.GetEntitiesWithMetricsRet(query_arg)
        
        if not ret or not ret.group_results_list:
            print("No files found in IDF")
            return
        
        group_results = ret.group_results_list[0]
        entities = group_results.raw_results
        
        if not entities:
            print("No files found in IDF")
            return
        
        print(f"\n{'ExtId':<40} {'Filename':<50} {'Size':<20} {'Content Type':<20}")
        print("=" * 130)
        
        for entity_with_metric in entities:
            entity = entity_with_metric.entity
            ext_id = entity.entity_id if entity.entity_id else "N/A"
            
            # Extract attributes
            filename = "N/A"
            file_size = "N/A"
            content_type = "N/A"
            
            for attr in entity.attribute_data_map:
                attr_name = attr.name
                if not attr_name:
                    continue
                
                if attr_name == "file_name" and attr.value and attr.value.str_value:
                    filename = attr.value.str_value
                elif attr_name == "file_size" and attr.value:
                    file_size = str(attr.value.int64_value)
                elif attr_name == "content_type" and attr.value and attr.value.str_value:
                    content_type = attr.value.str_value
            
            # Format file size
            if file_size != "N/A":
                try:
                    size_int = int(file_size)
                    if size_int > 1024 * 1024:
                        file_size = f"{size_int / (1024*1024):.2f} MB"
                    elif size_int > 1024:
                        file_size = f"{size_int / 1024:.2f} KB"
                    else:
                        file_size = f"{size_int} bytes"
                except:
                    pass
            
            print(f"{ext_id:<40} {filename:<50} {file_size:<20} {content_type:<20}")
        
        print(f"\nTotal files: {len(entities)}")
        
    except InsightsInterfaceError as ex:
        print(f"Error querying IDF: {ex.message}")
        if hasattr(ex, 'ret'):
            print(ex.ret)
    except Exception as e:
        print(f"Unexpected error: {e}")
        import traceback
        traceback.print_exc()

if __name__ == "__main__":
    # Suppress gflags warnings
    import sys
    import logging
    logging.getLogger().setLevel(logging.ERROR)
    
    with warnings.catch_warnings():
        warnings.simplefilter("ignore")
        # Parse flags to avoid gflags errors
        try:
            import gflags
            FLAGS = gflags.FLAGS
            FLAGS(sys.argv)
        except:
            pass
        list_files()

