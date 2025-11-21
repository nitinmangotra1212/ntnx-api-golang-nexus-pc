#!/bin/bash
#
# CHANGE THIS FILE TO PUSH YOUR GO CODE TO A GO REPOSITORY
API_SERVER_SOURCE_PATH=../../generated-code
DTO_PATH=$API_SERVER_SOURCE_PATH/dto
 
mkdir -p $DTO_PATH
cp -r target/generated-sources/swagger/src/* $DTO_PATH/

export old_path="models/"
export new_path="github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto/models/"

export folder_path=$DTO_PATH/models
echo "Start go dto"
for file in $(find $folder_path -type f)
do
    if [[ -f "$file" && "$file" =~ \.go$ ]]; then
        echo "Processing file: $file"
        # Fix import paths: models/... -> github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto/models/...
        # Handle imports with quotes (standard Go import format)
        sed -i '' "s#\"models/#\"github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto/models/#g" "$file"
        # Also handle imports without quotes (backup pattern for edge cases)
        sed -i '' "s#models/#github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto/models/#g" "$file"
    fi
done
echo "Done"

