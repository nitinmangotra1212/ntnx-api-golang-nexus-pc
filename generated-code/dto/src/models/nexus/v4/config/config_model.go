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
  import1 "github.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/dto/models/nexus/v4/error"
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
  Unique identifier for the item
  */
  ItemId *int `json:"itemId,omitempty"`
  /*
  Name of the item
  */
  ItemName *string `json:"itemName"`
  /*
  Type of item
  */
  ItemType *string `json:"itemType"`
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
    if known.ItemId != nil {
        p.ItemId = known.ItemId
    }
    if known.ItemName != nil {
        p.ItemName = known.ItemName
    }
    if known.ItemType != nil {
        p.ItemType = known.ItemType
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "associations")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "itemId")
	delete(allFields, "itemName")
	delete(allFields, "itemType")

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
  Unique identifier for the item
  */
  ItemId *int `json:"itemId,omitempty"`
  /*
  Name of the item
  */
  ItemName *string `json:"itemName"`
  /*
  Type of item
  */
  ItemType *string `json:"itemType"`
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
    if known.ItemId != nil {
        p.ItemId = known.ItemId
    }
    if known.ItemName != nil {
        p.ItemName = known.ItemName
    }
    if known.ItemType != nil {
        p.ItemType = known.ItemType
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "associations")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "itemId")
	delete(allFields, "itemName")
	delete(allFields, "itemType")

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


type OneOfListItemsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType401 []ItemProjection `json:"-"`
  oneOfType2001 []Item `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
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
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
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
          vOneOfType400 := new(import1.ErrorResponse)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
          if unmarshalErr == nil {
            // For struct items, verify the ObjectType matches
            if vOneOfType400.ObjectType_ != nil && "nexus.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
              if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
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
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if vOneOfType400.ObjectType_ != nil && "nexus.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
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
  return nil, errors.New("No value to marshal for OneOfListItemsApiResponseData")
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

