#!/bin/bash
#
# CHANGE THIS FILE TO PUSH YOUR GO CODE TO A GO REPOSITORY
API_SERVER_SOURCE_PATH=../../generated-code
DTO_PATH=$API_SERVER_SOURCE_PATH/dto
 
mkdir -p $DTO_PATH
cp -r target/generated-sources/swagger/src/* $DTO_PATH/

export old_path="models/"
export new_path="github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto/models/"

export folder_path=$DTO_PATH/src/models
echo "Start go dto"
if [ ! -d "$folder_path" ]; then
    echo "Error: $folder_path does not exist"
    exit 1
fi
for file in $(find $folder_path -type f -name "*.go")
do
    if [[ -f "$file" ]]; then
        echo "Processing file: $file"
        # Fix import paths: models/... -> github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto/models/...
        # Handle imports with quotes (standard Go import format), including import aliases
        # Pattern matches: import1 "models/... or "models/... or import "models/...
        # Only replace "models/ if it hasn't already been replaced (doesn't start with github.com)
        if grep -q '"models/' "$file"; then
            sed -i '' 's|"models/|"github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto/models/|g' "$file"
        fi
    fi
done
echo "Done"

