/*
 * Generated file models/nexus/v4/config/config_model.go.
 *
 * Product version: 1.0.0-SNAPSHOT
 *
 * Part of the GoLang Mock API - REST API for Mock Item Service
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module nexus.v4.config of GoLang Mock API - REST API for Mock Item Service
*/
package config
import (
  import2 "github.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/dto/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import3 "github.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/dto/models/nexus/v4/error"
  import1 "github.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/dto/models/nexus/v4/stats"
)
/*
File entity for file transfer operations
*/
type File struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /*
  MIME type of the file
  */
  ContentType *string `json:"contentType,omitempty"`
  /*
  External identifier (UUID)
  */
  ExtId *string `json:"extId,omitempty"`
  /*
  Unique identifier for the file
  */
  FileId *string `json:"fileId,omitempty"`
  /*
  Name of the file
  */
  FileName *string `json:"fileName,omitempty"`
  /*
  Size of the file in bytes
  */
  FileSize *int64 `json:"fileSize,omitempty"`
}

func (p *File) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias File

  // Step 1: Marshal the known fields
  known, err := json.Marshal(Alias(*p))
  if err != nil {
  	return nil, err
  }

    // Step 2: Convert known to map for merging
    var knownMap map[string]interface{}
    if err := json.Unmarshal(known, &knownMap); err != nil {
    	return nil, err
    }
    delete(knownMap, "$unknownFields")
  
    // Step 3: Merge unknown fields
    for k, v := range p.UnknownFields_ {
    	knownMap[k] = v
    }
  
    // Step 4: Marshal final merged map
    return json.Marshal(knownMap)
}

func (p *File) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias File
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewFile()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.ContentType != nil {
        p.ContentType = known.ContentType
    }
    if known.ExtId != nil {
        p.ExtId = known.ExtId
    }
    if known.FileId != nil {
        p.FileId = known.FileId
    }
    if known.FileName != nil {
        p.FileName = known.FileName
    }
    if known.FileSize != nil {
        p.FileSize = known.FileSize
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "contentType")
	delete(allFields, "extId")
	delete(allFields, "fileId")
	delete(allFields, "fileName")
	delete(allFields, "fileSize")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewFile() *File {
  p := new(File)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.config.File"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



/*
Item entity for mock REST API
*/
type Item struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /*
  Associated entities for this item. This field is only present when $expand=associations is specified in the query.
  */
  Associations []ItemAssociation `json:"associations,omitempty"`
  /*
  Description of the item
  */
  Description *string `json:"description,omitempty"`
  /*
  External identifier for the item (UUID)
  */
  ExtId *string `json:"extId,omitempty"`
  /*
  List of integer values (int64)
  */
  Int64List []int64 `json:"int64List,omitempty"`
  /*
  Whether the item is active
  */
  IsActive *bool `json:"isActive,omitempty"`
  /*
  Unique identifier for the item
  */
  ItemId *int `json:"itemId,omitempty"`
  /*
  Name of the item
  */
  ItemName *string `json:"itemName"`
  
  ItemStats *import1.ItemStats `json:"itemStats,omitempty"`
  /*
  Type of item
  */
  ItemType *string `json:"itemType"`
  /*
  Price/cost of the item
  */
  Price *float64 `json:"price,omitempty"`
  /*
  Priority level of the item (0-255)
  */
  Priority *int `json:"priority,omitempty"`
  /*
  Quantity/stock count of the item
  */
  Quantity *int64 `json:"quantity,omitempty"`
  /*
  Status of the item (e.g., ACTIVE, INACTIVE, PENDING)
  */
  Status *string `json:"status,omitempty"`
}

func (p *Item) MarshalJSON() ([]byte, error) {
  type ItemProxy Item

  // Step 1: Marshal known fields via proxy to enforce required fields
  baseStruct := struct {
    *ItemProxy
    ItemName *string `json:"itemName,omitempty"`
    ItemType *string `json:"itemType,omitempty"`
  }{
    ItemProxy : (*ItemProxy)(p),
    ItemName : p.ItemName,
    ItemType : p.ItemType,
  }

  known, err := json.Marshal(baseStruct)
  if err != nil {
  	return nil, err
  }

    // Step 2: Convert known to map for merging
    var knownMap map[string]interface{}
    if err := json.Unmarshal(known, &knownMap); err != nil {
    	return nil, err
    }
    delete(knownMap, "$unknownFields")
  
    // Step 3: Merge unknown fields
    for k, v := range p.UnknownFields_ {
    	knownMap[k] = v
    }
  
    // Step 4: Marshal final merged map
    return json.Marshal(knownMap)
}

func (p *Item) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias Item
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItem()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.Associations != nil {
        p.Associations = known.Associations
    }
    if known.Description != nil {
        p.Description = known.Description
    }
    if known.ExtId != nil {
        p.ExtId = known.ExtId
    }
    if known.Int64List != nil {
        p.Int64List = known.Int64List
    }
    if known.IsActive != nil {
        p.IsActive = known.IsActive
    }
    if known.ItemId != nil {
        p.ItemId = known.ItemId
    }
    if known.ItemName != nil {
        p.ItemName = known.ItemName
    }
    if known.ItemStats != nil {
        p.ItemStats = known.ItemStats
    }
    if known.ItemType != nil {
        p.ItemType = known.ItemType
    }
    if known.Price != nil {
        p.Price = known.Price
    }
    if known.Priority != nil {
        p.Priority = known.Priority
    }
    if known.Quantity != nil {
        p.Quantity = known.Quantity
    }
    if known.Status != nil {
        p.Status = known.Status
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "associations")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "int64List")
	delete(allFields, "isActive")
	delete(allFields, "itemId")
	delete(allFields, "itemName")
	delete(allFields, "itemStats")
	delete(allFields, "itemType")
	delete(allFields, "price")
	delete(allFields, "priority")
	delete(allFields, "quantity")
	delete(allFields, "status")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewItem() *Item {
  p := new(Item)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.config.Item"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




type ItemAggregate struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Label *string `json:"label,omitempty"`
  
  ResultItemDiscriminator_ *string `json:"$resultItemDiscriminator,omitempty"`
  
  Result *OneOfItemAggregateResult `json:"result,omitempty"`
}

func (p *ItemAggregate) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ItemAggregate

  // Step 1: Marshal the known fields
  known, err := json.Marshal(Alias(*p))
  if err != nil {
  	return nil, err
  }

    // Step 2: Convert known to map for merging
    var knownMap map[string]interface{}
    if err := json.Unmarshal(known, &knownMap); err != nil {
    	return nil, err
    }
    delete(knownMap, "$unknownFields")
  
    // Step 3: Merge unknown fields
    for k, v := range p.UnknownFields_ {
    	knownMap[k] = v
    }
  
    // Step 4: Marshal final merged map
    return json.Marshal(knownMap)
}

func (p *ItemAggregate) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ItemAggregate
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItemAggregate()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.Label != nil {
        p.Label = known.Label
    }
    if known.ResultItemDiscriminator_ != nil {
        p.ResultItemDiscriminator_ = known.ResultItemDiscriminator_
    }
    if known.Result != nil {
        p.Result = known.Result
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "label")
	delete(allFields, "$resultItemDiscriminator")
	delete(allFields, "result")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewItemAggregate() *ItemAggregate {
  p := new(ItemAggregate)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.config.ItemAggregate"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ItemAggregate) GetResult() interface{} {
  if nil == p.Result {
    return nil
  }
  return p.Result.GetValue()
}

func (p *ItemAggregate) SetResult(v interface{}) error {
  if nil == p.Result {
    p.Result = NewOneOfItemAggregateResult()
  }
  e := p.Result.SetValue(v)
  if nil == e {
    if nil == p.ResultItemDiscriminator_ {
      p.ResultItemDiscriminator_ = new(string)
    }
    *p.ResultItemDiscriminator_ = *p.Result.Discriminator
  }
  return e
}


/*
Association entity for items, representing related entities associated with an item
*/
type ItemAssociation struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /*
  Count of associations of this type
  */
  Count *int `json:"count,omitempty"`
  /*
  ID of associated entity
  */
  EntityId *string `json:"entityId,omitempty"`
  /*
  Type of associated entity
  */
  EntityType *string `json:"entityType,omitempty"`
  /*
  The item ID this association belongs to
  */
  ItemId *string `json:"itemId,omitempty"`
}

func (p *ItemAssociation) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ItemAssociation

  // Step 1: Marshal the known fields
  known, err := json.Marshal(Alias(*p))
  if err != nil {
  	return nil, err
  }

    // Step 2: Convert known to map for merging
    var knownMap map[string]interface{}
    if err := json.Unmarshal(known, &knownMap); err != nil {
    	return nil, err
    }
    delete(knownMap, "$unknownFields")
  
    // Step 3: Merge unknown fields
    for k, v := range p.UnknownFields_ {
    	knownMap[k] = v
    }
  
    // Step 4: Marshal final merged map
    return json.Marshal(knownMap)
}

func (p *ItemAssociation) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ItemAssociation
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItemAssociation()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.Count != nil {
        p.Count = known.Count
    }
    if known.EntityId != nil {
        p.EntityId = known.EntityId
    }
    if known.EntityType != nil {
        p.EntityType = known.EntityType
    }
    if known.ItemId != nil {
        p.ItemId = known.ItemId
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "count")
	delete(allFields, "entityId")
	delete(allFields, "entityType")
	delete(allFields, "itemId")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewItemAssociation() *ItemAssociation {
  p := new(ItemAssociation)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.config.ItemAssociation"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




type ItemAssociationProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /*
  Count of associations of this type
  */
  Count *int `json:"count,omitempty"`
  /*
  ID of associated entity
  */
  EntityId *string `json:"entityId,omitempty"`
  /*
  Type of associated entity
  */
  EntityType *string `json:"entityType,omitempty"`
  /*
  The item ID this association belongs to
  */
  ItemId *string `json:"itemId,omitempty"`
}

func (p *ItemAssociationProjection) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ItemAssociationProjection

  // Step 1: Marshal the known fields
  known, err := json.Marshal(Alias(*p))
  if err != nil {
  	return nil, err
  }

    // Step 2: Convert known to map for merging
    var knownMap map[string]interface{}
    if err := json.Unmarshal(known, &knownMap); err != nil {
    	return nil, err
    }
    delete(knownMap, "$unknownFields")
  
    // Step 3: Merge unknown fields
    for k, v := range p.UnknownFields_ {
    	knownMap[k] = v
    }
  
    // Step 4: Marshal final merged map
    return json.Marshal(knownMap)
}

func (p *ItemAssociationProjection) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ItemAssociationProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItemAssociationProjection()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.Count != nil {
        p.Count = known.Count
    }
    if known.EntityId != nil {
        p.EntityId = known.EntityId
    }
    if known.EntityType != nil {
        p.EntityType = known.EntityType
    }
    if known.ItemId != nil {
        p.ItemId = known.ItemId
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "count")
	delete(allFields, "entityId")
	delete(allFields, "entityType")
	delete(allFields, "itemId")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewItemAssociationProjection() *ItemAssociationProjection {
  p := new(ItemAssociationProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.config.ItemAssociationProjection"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




type ItemGroup struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Aggregates []ItemAggregate `json:"aggregates,omitempty"`
  
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfItemGroupData `json:"data,omitempty"`
  
  GroupItemDiscriminator_ *string `json:"$groupItemDiscriminator,omitempty"`
  
  Group *OneOfItemGroupGroup `json:"group,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ItemGroup) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ItemGroup

  // Step 1: Marshal the known fields
  known, err := json.Marshal(Alias(*p))
  if err != nil {
  	return nil, err
  }

    // Step 2: Convert known to map for merging
    var knownMap map[string]interface{}
    if err := json.Unmarshal(known, &knownMap); err != nil {
    	return nil, err
    }
    delete(knownMap, "$unknownFields")
  
    // Step 3: Merge unknown fields
    for k, v := range p.UnknownFields_ {
    	knownMap[k] = v
    }
  
    // Step 4: Marshal final merged map
    return json.Marshal(knownMap)
}

func (p *ItemGroup) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ItemGroup
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItemGroup()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.Aggregates != nil {
        p.Aggregates = known.Aggregates
    }
    if known.DataItemDiscriminator_ != nil {
        p.DataItemDiscriminator_ = known.DataItemDiscriminator_
    }
    if known.Data != nil {
        p.Data = known.Data
    }
    if known.GroupItemDiscriminator_ != nil {
        p.GroupItemDiscriminator_ = known.GroupItemDiscriminator_
    }
    if known.Group != nil {
        p.Group = known.Group
    }
    if known.Metadata != nil {
        p.Metadata = known.Metadata
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "aggregates")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "$groupItemDiscriminator")
	delete(allFields, "group")
	delete(allFields, "metadata")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewItemGroup() *ItemGroup {
  p := new(ItemGroup)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.config.ItemGroup"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ItemGroup) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ItemGroup) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfItemGroupData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



type ItemProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /*
  Associated entities for this item. This field is only present when $expand=associations is specified in the query.
  */
  Associations []ItemAssociation `json:"associations,omitempty"`
  /*
  Description of the item
  */
  Description *string `json:"description,omitempty"`
  /*
  External identifier for the item (UUID)
  */
  ExtId *string `json:"extId,omitempty"`
  /*
  List of integer values (int64)
  */
  Int64List []int64 `json:"int64List,omitempty"`
  /*
  Whether the item is active
  */
  IsActive *bool `json:"isActive,omitempty"`
  /*
  Unique identifier for the item
  */
  ItemId *int `json:"itemId,omitempty"`
  /*
  Name of the item
  */
  ItemName *string `json:"itemName"`
  
  ItemStats *import1.ItemStats `json:"itemStats,omitempty"`
  /*
  Type of item
  */
  ItemType *string `json:"itemType"`
  /*
  Price/cost of the item
  */
  Price *float64 `json:"price,omitempty"`
  /*
  Priority level of the item (0-255)
  */
  Priority *int `json:"priority,omitempty"`
  /*
  Quantity/stock count of the item
  */
  Quantity *int64 `json:"quantity,omitempty"`
  /*
  Status of the item (e.g., ACTIVE, INACTIVE, PENDING)
  */
  Status *string `json:"status,omitempty"`
}

func (p *ItemProjection) MarshalJSON() ([]byte, error) {
  type ItemProjectionProxy ItemProjection

  // Step 1: Marshal known fields via proxy to enforce required fields
  baseStruct := struct {
    *ItemProjectionProxy
    ItemName *string `json:"itemName,omitempty"`
    ItemType *string `json:"itemType,omitempty"`
  }{
    ItemProjectionProxy : (*ItemProjectionProxy)(p),
    ItemName : p.ItemName,
    ItemType : p.ItemType,
  }

  known, err := json.Marshal(baseStruct)
  if err != nil {
  	return nil, err
  }

    // Step 2: Convert known to map for merging
    var knownMap map[string]interface{}
    if err := json.Unmarshal(known, &knownMap); err != nil {
    	return nil, err
    }
    delete(knownMap, "$unknownFields")
  
    // Step 3: Merge unknown fields
    for k, v := range p.UnknownFields_ {
    	knownMap[k] = v
    }
  
    // Step 4: Marshal final merged map
    return json.Marshal(knownMap)
}

func (p *ItemProjection) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ItemProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItemProjection()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.Associations != nil {
        p.Associations = known.Associations
    }
    if known.Description != nil {
        p.Description = known.Description
    }
    if known.ExtId != nil {
        p.ExtId = known.ExtId
    }
    if known.Int64List != nil {
        p.Int64List = known.Int64List
    }
    if known.IsActive != nil {
        p.IsActive = known.IsActive
    }
    if known.ItemId != nil {
        p.ItemId = known.ItemId
    }
    if known.ItemName != nil {
        p.ItemName = known.ItemName
    }
    if known.ItemStats != nil {
        p.ItemStats = known.ItemStats
    }
    if known.ItemType != nil {
        p.ItemType = known.ItemType
    }
    if known.Price != nil {
        p.Price = known.Price
    }
    if known.Priority != nil {
        p.Priority = known.Priority
    }
    if known.Quantity != nil {
        p.Quantity = known.Quantity
    }
    if known.Status != nil {
        p.Status = known.Status
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "associations")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "int64List")
	delete(allFields, "isActive")
	delete(allFields, "itemId")
	delete(allFields, "itemName")
	delete(allFields, "itemStats")
	delete(allFields, "itemType")
	delete(allFields, "price")
	delete(allFields, "priority")
	delete(allFields, "quantity")
	delete(allFields, "status")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewItemProjection() *ItemProjection {
  p := new(ItemProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.config.ItemProjection"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




type ItemTimeValuePair struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  TimeStamp *int64 `json:"timeStamp,omitempty"`
  
  Value *int64 `json:"value,omitempty"`
}

func (p *ItemTimeValuePair) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ItemTimeValuePair

  // Step 1: Marshal the known fields
  known, err := json.Marshal(Alias(*p))
  if err != nil {
  	return nil, err
  }

    // Step 2: Convert known to map for merging
    var knownMap map[string]interface{}
    if err := json.Unmarshal(known, &knownMap); err != nil {
    	return nil, err
    }
    delete(knownMap, "$unknownFields")
  
    // Step 3: Merge unknown fields
    for k, v := range p.UnknownFields_ {
    	knownMap[k] = v
    }
  
    // Step 4: Marshal final merged map
    return json.Marshal(knownMap)
}

func (p *ItemTimeValuePair) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ItemTimeValuePair
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItemTimeValuePair()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.TimeStamp != nil {
        p.TimeStamp = known.TimeStamp
    }
    if known.Value != nil {
        p.Value = known.Value
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "timeStamp")
	delete(allFields, "value")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewItemTimeValuePair() *ItemTimeValuePair {
  p := new(ItemTimeValuePair)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.config.ItemTimeValuePair"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



/*
REST response for all response codes in API path /nexus/v4.1/config/items Get operation
*/
type ListItemsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /*
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfListItemsApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListItemsApiResponse) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ListItemsApiResponse

  // Step 1: Marshal the known fields
  known, err := json.Marshal(Alias(*p))
  if err != nil {
  	return nil, err
  }

    // Step 2: Convert known to map for merging
    var knownMap map[string]interface{}
    if err := json.Unmarshal(known, &knownMap); err != nil {
    	return nil, err
    }
    delete(knownMap, "$unknownFields")
  
    // Step 3: Merge unknown fields
    for k, v := range p.UnknownFields_ {
    	knownMap[k] = v
    }
  
    // Step 4: Marshal final merged map
    return json.Marshal(knownMap)
}

func (p *ListItemsApiResponse) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListItemsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewListItemsApiResponse()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.DataItemDiscriminator_ != nil {
        p.DataItemDiscriminator_ = known.DataItemDiscriminator_
    }
    if known.Data != nil {
        p.Data = known.Data
    }
    if known.Metadata != nil {
        p.Metadata = known.Metadata
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewListItemsApiResponse() *ListItemsApiResponse {
  p := new(ListItemsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.config.ListItemsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ListItemsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ListItemsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfListItemsApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}


type OneOfItemAggregateResult struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType2003 *int `json:"-"`
  oneOfType2005 *float64 `json:"-"`
  oneOfType2006 []ItemTimeValuePair `json:"-"`
  oneOfType2004 *int64 `json:"-"`
}

func NewOneOfItemAggregateResult() *OneOfItemAggregateResult {
  p := new(OneOfItemAggregateResult)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfItemAggregateResult) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfItemAggregateResult is nil"))
  }
  switch v.(type) {
    case int:
      if nil == p.oneOfType2003 {p.oneOfType2003 = new(int)}
      *p.oneOfType2003 = v.(int)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Integer"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Integer"
    case float64:
      if nil == p.oneOfType2005 {p.oneOfType2005 = new(float64)}
      *p.oneOfType2005 = v.(float64)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Double"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Double"
    case []ItemTimeValuePair:
      p.oneOfType2006 = v.([]ItemTimeValuePair)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.config.ItemTimeValuePair>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.config.ItemTimeValuePair>"
    case int64:
      if nil == p.oneOfType2004 {p.oneOfType2004 = new(int64)}
      *p.oneOfType2004 = v.(int64)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Long"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Long"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfItemAggregateResult) GetValue() interface{} {
  if "Integer" == *p.Discriminator {
    return *p.oneOfType2003
  }
  if "Double" == *p.Discriminator {
    return *p.oneOfType2005
  }
  if "List<nexus.v4.config.ItemTimeValuePair>" == *p.Discriminator {
    return p.oneOfType2006
  }
  if "Long" == *p.Discriminator {
    return *p.oneOfType2004
  }
  return nil
}

func (p *OneOfItemAggregateResult) UnmarshalJSON(b []byte) error {

  // Try to handle nested structure like {"": {"value": {...}}}
  // This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
  var rawMap map[string]interface{}
  if err := json.Unmarshal(b, &rawMap); err == nil {
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["Integer"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2003 := new(int)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2003)
          if unmarshalErr == nil {
              if nil == p.oneOfType2003 {p.oneOfType2003 = new(int)}
              *p.oneOfType2003 = *vOneOfType2003
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "Integer"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "Integer"
              return nil
          }
        }
      }
    }
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["Double"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2005 := new(float64)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2005)
          if unmarshalErr == nil {
              if nil == p.oneOfType2005 {p.oneOfType2005 = new(float64)}
              *p.oneOfType2005 = *vOneOfType2005
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "Double"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "Double"
              return nil
          }
        }
      }
    }
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["List<nexus.v4.config.ItemTimeValuePair>"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2006 := new([]ItemTimeValuePair)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2006)
          if unmarshalErr == nil {
            // For arrays, verify the array item ObjectType matches
            if vOneOfType2006 == nil || len(*vOneOfType2006) == 0 || ((*vOneOfType2006)[0].ObjectType_ != nil && "nexus.v4.config.ItemTimeValuePair" == *((*vOneOfType2006)[0].ObjectType_)) {
              p.oneOfType2006 = *vOneOfType2006
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "List<nexus.v4.config.ItemTimeValuePair>"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "List<nexus.v4.config.ItemTimeValuePair>"
              return nil
            }
          }
        }
      }
    }
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["Long"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2004 := new(int64)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2004)
          if unmarshalErr == nil {
              if nil == p.oneOfType2004 {p.oneOfType2004 = new(int64)}
              *p.oneOfType2004 = *vOneOfType2004
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "Long"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "Long"
              return nil
          }
        }
      }
    }
  }

  // Fallback: try direct unmarshalling (for non-nested structures)
  vOneOfType2003 := new(int)
  if err := json.Unmarshal(b, vOneOfType2003); err == nil {
      if nil == p.oneOfType2003 {p.oneOfType2003 = new(int)}
      *p.oneOfType2003 = *vOneOfType2003
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Integer"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Integer"
      return nil
  }
  vOneOfType2005 := new(float64)
  if err := json.Unmarshal(b, vOneOfType2005); err == nil {
      if nil == p.oneOfType2005 {p.oneOfType2005 = new(float64)}
      *p.oneOfType2005 = *vOneOfType2005
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Double"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Double"
      return nil
  }
  vOneOfType2006 := new([]ItemTimeValuePair)
  if err := json.Unmarshal(b, vOneOfType2006); err == nil {
    if len(*vOneOfType2006) == 0 || (vOneOfType2006 != nil && (*vOneOfType2006)[0].ObjectType_ != nil && "nexus.v4.config.ItemTimeValuePair" == *((*vOneOfType2006)[0].ObjectType_)) {
      p.oneOfType2006 = *vOneOfType2006
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.config.ItemTimeValuePair>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.config.ItemTimeValuePair>"
      return nil
    }
  }
  vOneOfType2004 := new(int64)
  if err := json.Unmarshal(b, vOneOfType2004); err == nil {
      if nil == p.oneOfType2004 {p.oneOfType2004 = new(int64)}
      *p.oneOfType2004 = *vOneOfType2004
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Long"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Long"
      return nil
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfItemAggregateResult"))
}

func (p *OneOfItemAggregateResult) MarshalJSON() ([]byte, error) {
  if "Integer" == *p.Discriminator {
    return json.Marshal(p.oneOfType2003)
  }
  if "Double" == *p.Discriminator {
    return json.Marshal(p.oneOfType2005)
  }
  if "List<nexus.v4.config.ItemTimeValuePair>" == *p.Discriminator {
    return json.Marshal(p.oneOfType2006)
  }
  if "Long" == *p.Discriminator {
    return json.Marshal(p.oneOfType2004)
  }
  return nil, errors.New("No value to marshal for OneOfItemAggregateResult")
}

type OneOfListItemsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType401 []ItemProjection `json:"-"`
  oneOfType2001 []Item `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType402 []ItemGroup `json:"-"`
}

func NewOneOfListItemsApiResponseData() *OneOfListItemsApiResponseData {
  p := new(OneOfListItemsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfListItemsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfListItemsApiResponseData is nil"))
  }
  switch v.(type) {
    case []ItemProjection:
      p.oneOfType401 = v.([]ItemProjection)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.config.ItemProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.config.ItemProjection>"
    case []Item:
      p.oneOfType2001 = v.([]Item)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.config.Item>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.config.Item>"
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []ItemGroup:
      p.oneOfType402 = v.([]ItemGroup)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.config.ItemGroup>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.config.ItemGroup>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfListItemsApiResponseData) GetValue() interface{} {
  if "List<nexus.v4.config.ItemProjection>" == *p.Discriminator {
    return p.oneOfType401
  }
  if "List<nexus.v4.config.Item>" == *p.Discriminator {
    return p.oneOfType2001
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<nexus.v4.config.ItemGroup>" == *p.Discriminator {
    return p.oneOfType402
  }
  return nil
}

func (p *OneOfListItemsApiResponseData) UnmarshalJSON(b []byte) error {

  // Try to handle nested structure like {"": {"value": {...}}}
  // This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
  var rawMap map[string]interface{}
  if err := json.Unmarshal(b, &rawMap); err == nil {
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["List<nexus.v4.config.ItemProjection>"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType401 := new([]ItemProjection)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
          if unmarshalErr == nil {
            // For arrays, verify the array item ObjectType matches
            if vOneOfType401 == nil || len(*vOneOfType401) == 0 || ((*vOneOfType401)[0].ObjectType_ != nil && "nexus.v4.config.ItemProjection" == *((*vOneOfType401)[0].ObjectType_)) {
              p.oneOfType401 = *vOneOfType401
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "List<nexus.v4.config.ItemProjection>"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "List<nexus.v4.config.ItemProjection>"
              return nil
            }
          }
        }
      }
    }
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["List<nexus.v4.config.Item>"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2001 := new([]Item)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
          if unmarshalErr == nil {
            // For arrays, verify the array item ObjectType matches
            if vOneOfType2001 == nil || len(*vOneOfType2001) == 0 || ((*vOneOfType2001)[0].ObjectType_ != nil && "nexus.v4.config.Item" == *((*vOneOfType2001)[0].ObjectType_)) {
              p.oneOfType2001 = *vOneOfType2001
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "List<nexus.v4.config.Item>"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "List<nexus.v4.config.Item>"
              return nil
            }
          }
        }
      }
    }
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType400 := new(import3.ErrorResponse)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
          if unmarshalErr == nil {
            // For struct items, verify the ObjectType matches
            if vOneOfType400.ObjectType_ != nil && "nexus.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
              if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
              *p.oneOfType400 = *vOneOfType400
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = *p.oneOfType400.ObjectType_
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = *p.oneOfType400.ObjectType_
              return nil
            }
          }
        }
      }
    }
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["List<nexus.v4.config.ItemGroup>"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType402 := new([]ItemGroup)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType402)
          if unmarshalErr == nil {
            // For arrays, verify the array item ObjectType matches
            if vOneOfType402 == nil || len(*vOneOfType402) == 0 || ((*vOneOfType402)[0].ObjectType_ != nil && "nexus.v4.config.ItemGroup" == *((*vOneOfType402)[0].ObjectType_)) {
              p.oneOfType402 = *vOneOfType402
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "List<nexus.v4.config.ItemGroup>"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "List<nexus.v4.config.ItemGroup>"
              return nil
            }
          }
        }
      }
    }
  }

  // Fallback: try direct unmarshalling (for non-nested structures)
  vOneOfType401 := new([]ItemProjection)
  if err := json.Unmarshal(b, vOneOfType401); err == nil {
    if len(*vOneOfType401) == 0 || (vOneOfType401 != nil && (*vOneOfType401)[0].ObjectType_ != nil && "nexus.v4.config.ItemProjection" == *((*vOneOfType401)[0].ObjectType_)) {
      p.oneOfType401 = *vOneOfType401
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.config.ItemProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.config.ItemProjection>"
      return nil
    }
  }
  vOneOfType2001 := new([]Item)
  if err := json.Unmarshal(b, vOneOfType2001); err == nil {
    if len(*vOneOfType2001) == 0 || (vOneOfType2001 != nil && (*vOneOfType2001)[0].ObjectType_ != nil && "nexus.v4.config.Item" == *((*vOneOfType2001)[0].ObjectType_)) {
      p.oneOfType2001 = *vOneOfType2001
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.config.Item>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.config.Item>"
      return nil
    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if vOneOfType400.ObjectType_ != nil && "nexus.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
  }
  vOneOfType402 := new([]ItemGroup)
  if err := json.Unmarshal(b, vOneOfType402); err == nil {
    if len(*vOneOfType402) == 0 || (vOneOfType402 != nil && (*vOneOfType402)[0].ObjectType_ != nil && "nexus.v4.config.ItemGroup" == *((*vOneOfType402)[0].ObjectType_)) {
      p.oneOfType402 = *vOneOfType402
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.config.ItemGroup>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.config.ItemGroup>"
      return nil
    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListItemsApiResponseData"))
}

func (p *OneOfListItemsApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<nexus.v4.config.ItemProjection>" == *p.Discriminator {
    return json.Marshal(p.oneOfType401)
  }
  if "List<nexus.v4.config.Item>" == *p.Discriminator {
    return json.Marshal(p.oneOfType2001)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<nexus.v4.config.ItemGroup>" == *p.Discriminator {
    return json.Marshal(p.oneOfType402)
  }
  return nil, errors.New("No value to marshal for OneOfListItemsApiResponseData")
}

type OneOfItemGroupData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType2008 []Item `json:"-"`
}

func NewOneOfItemGroupData() *OneOfItemGroupData {
  p := new(OneOfItemGroupData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfItemGroupData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfItemGroupData is nil"))
  }
  switch v.(type) {
    case []Item:
      p.oneOfType2008 = v.([]Item)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.config.Item>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.config.Item>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfItemGroupData) GetValue() interface{} {
  if "List<nexus.v4.config.Item>" == *p.Discriminator {
    return p.oneOfType2008
  }
  return nil
}

func (p *OneOfItemGroupData) UnmarshalJSON(b []byte) error {

  // Try to handle nested structure like {"": {"value": {...}}}
  // This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
  var rawMap map[string]interface{}
  if err := json.Unmarshal(b, &rawMap); err == nil {
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["List<nexus.v4.config.Item>"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2008 := new([]Item)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2008)
          if unmarshalErr == nil {
            // For arrays, verify the array item ObjectType matches
            if vOneOfType2008 == nil || len(*vOneOfType2008) == 0 || ((*vOneOfType2008)[0].ObjectType_ != nil && "nexus.v4.config.Item" == *((*vOneOfType2008)[0].ObjectType_)) {
              p.oneOfType2008 = *vOneOfType2008
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "List<nexus.v4.config.Item>"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "List<nexus.v4.config.Item>"
              return nil
            }
          }
        }
      }
    }
  }

  // Fallback: try direct unmarshalling (for non-nested structures)
  vOneOfType2008 := new([]Item)
  if err := json.Unmarshal(b, vOneOfType2008); err == nil {
    if len(*vOneOfType2008) == 0 || (vOneOfType2008 != nil && (*vOneOfType2008)[0].ObjectType_ != nil && "nexus.v4.config.Item" == *((*vOneOfType2008)[0].ObjectType_)) {
      p.oneOfType2008 = *vOneOfType2008
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.config.Item>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.config.Item>"
      return nil
    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfItemGroupData"))
}

func (p *OneOfItemGroupData) MarshalJSON() ([]byte, error) {
  if "List<nexus.v4.config.Item>" == *p.Discriminator {
    return json.Marshal(p.oneOfType2008)
  }
  return nil, errors.New("No value to marshal for OneOfItemGroupData")
}

type OneOfItemGroupGroup struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType2005 *float64 `json:"-"`
  oneOfType2004 *int64 `json:"-"`
  oneOfType2002 *string `json:"-"`
  oneOfType2003 *int `json:"-"`
  oneOfType2006 *bool `json:"-"`
}

func NewOneOfItemGroupGroup() *OneOfItemGroupGroup {
  p := new(OneOfItemGroupGroup)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfItemGroupGroup) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfItemGroupGroup is nil"))
  }
  switch v.(type) {
    case float64:
      if nil == p.oneOfType2005 {p.oneOfType2005 = new(float64)}
      *p.oneOfType2005 = v.(float64)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Double"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Double"
    case int64:
      if nil == p.oneOfType2004 {p.oneOfType2004 = new(int64)}
      *p.oneOfType2004 = v.(int64)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Long"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Long"
    case string:
      if nil == p.oneOfType2002 {p.oneOfType2002 = new(string)}
      *p.oneOfType2002 = v.(string)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "String"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "String"
    case int:
      if nil == p.oneOfType2003 {p.oneOfType2003 = new(int)}
      *p.oneOfType2003 = v.(int)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Integer"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Integer"
    case bool:
      if nil == p.oneOfType2006 {p.oneOfType2006 = new(bool)}
      *p.oneOfType2006 = v.(bool)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Boolean"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Boolean"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfItemGroupGroup) GetValue() interface{} {
  if "Double" == *p.Discriminator {
    return *p.oneOfType2005
  }
  if "Long" == *p.Discriminator {
    return *p.oneOfType2004
  }
  if "String" == *p.Discriminator {
    return *p.oneOfType2002
  }
  if "Integer" == *p.Discriminator {
    return *p.oneOfType2003
  }
  if "Boolean" == *p.Discriminator {
    return *p.oneOfType2006
  }
  return nil
}

func (p *OneOfItemGroupGroup) UnmarshalJSON(b []byte) error {

  // Try to handle nested structure like {"": {"value": {...}}}
  // This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
  var rawMap map[string]interface{}
  if err := json.Unmarshal(b, &rawMap); err == nil {
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["Double"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2005 := new(float64)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2005)
          if unmarshalErr == nil {
              if nil == p.oneOfType2005 {p.oneOfType2005 = new(float64)}
              *p.oneOfType2005 = *vOneOfType2005
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "Double"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "Double"
              return nil
          }
        }
      }
    }
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["Long"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2004 := new(int64)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2004)
          if unmarshalErr == nil {
              if nil == p.oneOfType2004 {p.oneOfType2004 = new(int64)}
              *p.oneOfType2004 = *vOneOfType2004
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "Long"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "Long"
              return nil
          }
        }
      }
    }
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["String"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2002 := new(string)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2002)
          if unmarshalErr == nil {
              if nil == p.oneOfType2002 {p.oneOfType2002 = new(string)}
              *p.oneOfType2002 = *vOneOfType2002
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "String"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "String"
              return nil
          }
        }
      }
    }
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["Integer"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2003 := new(int)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2003)
          if unmarshalErr == nil {
              if nil == p.oneOfType2003 {p.oneOfType2003 = new(int)}
              *p.oneOfType2003 = *vOneOfType2003
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "Integer"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "Integer"
              return nil
          }
        }
      }
    }
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["Boolean"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2006 := new(bool)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2006)
          if unmarshalErr == nil {
              if nil == p.oneOfType2006 {p.oneOfType2006 = new(bool)}
              *p.oneOfType2006 = *vOneOfType2006
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "Boolean"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "Boolean"
              return nil
          }
        }
      }
    }
  }

  // Fallback: try direct unmarshalling (for non-nested structures)
  vOneOfType2005 := new(float64)
  if err := json.Unmarshal(b, vOneOfType2005); err == nil {
      if nil == p.oneOfType2005 {p.oneOfType2005 = new(float64)}
      *p.oneOfType2005 = *vOneOfType2005
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Double"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Double"
      return nil
  }
  vOneOfType2004 := new(int64)
  if err := json.Unmarshal(b, vOneOfType2004); err == nil {
      if nil == p.oneOfType2004 {p.oneOfType2004 = new(int64)}
      *p.oneOfType2004 = *vOneOfType2004
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Long"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Long"
      return nil
  }
  vOneOfType2002 := new(string)
  if err := json.Unmarshal(b, vOneOfType2002); err == nil {
      if nil == p.oneOfType2002 {p.oneOfType2002 = new(string)}
      *p.oneOfType2002 = *vOneOfType2002
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "String"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "String"
      return nil
  }
  vOneOfType2003 := new(int)
  if err := json.Unmarshal(b, vOneOfType2003); err == nil {
      if nil == p.oneOfType2003 {p.oneOfType2003 = new(int)}
      *p.oneOfType2003 = *vOneOfType2003
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Integer"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Integer"
      return nil
  }
  vOneOfType2006 := new(bool)
  if err := json.Unmarshal(b, vOneOfType2006); err == nil {
      if nil == p.oneOfType2006 {p.oneOfType2006 = new(bool)}
      *p.oneOfType2006 = *vOneOfType2006
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Boolean"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Boolean"
      return nil
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfItemGroupGroup"))
}

func (p *OneOfItemGroupGroup) MarshalJSON() ([]byte, error) {
  if "Double" == *p.Discriminator {
    return json.Marshal(p.oneOfType2005)
  }
  if "Long" == *p.Discriminator {
    return json.Marshal(p.oneOfType2004)
  }
  if "String" == *p.Discriminator {
    return json.Marshal(p.oneOfType2002)
  }
  if "Integer" == *p.Discriminator {
    return json.Marshal(p.oneOfType2003)
  }
  if "Boolean" == *p.Discriminator {
    return json.Marshal(p.oneOfType2006)
  }
  return nil, errors.New("No value to marshal for OneOfItemGroupGroup")
}


type FileDetail struct {
	Path *string `json:"-"`
	ObjectType_ *string `json:"-"`
}

func NewFileDetail() *FileDetail {
	p := new(FileDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "FileDetail"

	return p
}

