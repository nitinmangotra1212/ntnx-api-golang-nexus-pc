/*
 * Generated file models/nexus/v4/stats/stats_model.go.
 *
 * Product version: 1.0.0-SNAPSHOT
 *
 * Part of the GoLang Mock API - REST API for Mock Item Service
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module nexus.v4.stats of GoLang Mock API - REST API for Mock Item Service
*/
package stats
import (
  import1 "github.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/dto/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import2 "github.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/dto/models/nexus/v4/error"
  "time"
)
/*
Time-value pair for double/float metrics
*/
type DoubleTimeValuePair struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /*
  Timestamp when the value was recorded
  */
  Timestamp *time.Time `json:"timestamp,omitempty"`
  /*
  Double value
  */
  Value *float64 `json:"value,omitempty"`
}

func (p *DoubleTimeValuePair) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias DoubleTimeValuePair

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

func (p *DoubleTimeValuePair) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias DoubleTimeValuePair
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewDoubleTimeValuePair()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.Timestamp != nil {
        p.Timestamp = known.Timestamp
    }
    if known.Value != nil {
        p.Value = known.Value
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "timestamp")
	delete(allFields, "value")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewDoubleTimeValuePair() *DoubleTimeValuePair {
  p := new(DoubleTimeValuePair)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.stats.DoubleTimeValuePair"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



/*
Time-value pair for integer metrics
*/
type IntegerTimeValuePair struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /*
  Timestamp when the value was recorded
  */
  Timestamp *time.Time `json:"timestamp,omitempty"`
  /*
  Integer value
  */
  Value *int `json:"value,omitempty"`
}

func (p *IntegerTimeValuePair) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias IntegerTimeValuePair

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

func (p *IntegerTimeValuePair) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias IntegerTimeValuePair
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewIntegerTimeValuePair()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.Timestamp != nil {
        p.Timestamp = known.Timestamp
    }
    if known.Value != nil {
        p.Value = known.Value
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "timestamp")
	delete(allFields, "value")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewIntegerTimeValuePair() *IntegerTimeValuePair {
  p := new(IntegerTimeValuePair)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.stats.IntegerTimeValuePair"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



/*
Statistics entity for items, representing time-series or analytical data for an item
*/
type ItemStats struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /*
  Age of the item (time-series data as array of time-value pairs)
  */
  Age []IntegerTimeValuePair `json:"age,omitempty"`
  /*
  Food intake measurement (time-series data as array of time-value pairs)
  */
  FoodIntake []DoubleTimeValuePair `json:"foodIntake,omitempty"`
  /*
  Heart rate measurement (time-series data as array of time-value pairs)
  */
  HeartRate []IntegerTimeValuePair `json:"heartRate,omitempty"`
  /*
  External identifier of the item this stats record belongs to (foreign key to Item.extId)
  */
  ItemExtId *string `json:"itemExtId,omitempty"`
  /*
  Unique identifier for the stats record (primary key)
  */
  StatsExtId *string `json:"statsExtId,omitempty"`
}

func (p *ItemStats) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ItemStats

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

func (p *ItemStats) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ItemStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItemStats()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.Age != nil {
        p.Age = known.Age
    }
    if known.FoodIntake != nil {
        p.FoodIntake = known.FoodIntake
    }
    if known.HeartRate != nil {
        p.HeartRate = known.HeartRate
    }
    if known.ItemExtId != nil {
        p.ItemExtId = known.ItemExtId
    }
    if known.StatsExtId != nil {
        p.StatsExtId = known.StatsExtId
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "age")
	delete(allFields, "foodIntake")
	delete(allFields, "heartRate")
	delete(allFields, "itemExtId")
	delete(allFields, "statsExtId")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewItemStats() *ItemStats {
  p := new(ItemStats)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.stats.ItemStats"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




type ItemStatsAggregate struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Label *string `json:"label,omitempty"`
  
  ResultItemDiscriminator_ *string `json:"$resultItemDiscriminator,omitempty"`
  
  Result *OneOfItemStatsAggregateResult `json:"result,omitempty"`
}

func (p *ItemStatsAggregate) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ItemStatsAggregate

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

func (p *ItemStatsAggregate) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ItemStatsAggregate
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItemStatsAggregate()

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

func NewItemStatsAggregate() *ItemStatsAggregate {
  p := new(ItemStatsAggregate)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.stats.ItemStatsAggregate"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ItemStatsAggregate) GetResult() interface{} {
  if nil == p.Result {
    return nil
  }
  return p.Result.GetValue()
}

func (p *ItemStatsAggregate) SetResult(v interface{}) error {
  if nil == p.Result {
    p.Result = NewOneOfItemStatsAggregateResult()
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



type ItemStatsGroup struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Aggregates []ItemStatsAggregate `json:"aggregates,omitempty"`
  
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfItemStatsGroupData `json:"data,omitempty"`
  
  GroupItemDiscriminator_ *string `json:"$groupItemDiscriminator,omitempty"`
  
  Group *OneOfItemStatsGroupGroup `json:"group,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ItemStatsGroup) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ItemStatsGroup

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

func (p *ItemStatsGroup) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ItemStatsGroup
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItemStatsGroup()

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

func NewItemStatsGroup() *ItemStatsGroup {
  p := new(ItemStatsGroup)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.stats.ItemStatsGroup"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ItemStatsGroup) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ItemStatsGroup) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfItemStatsGroupData()
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



type ItemStatsProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /*
  Age of the item (time-series data as array of time-value pairs)
  */
  Age []IntegerTimeValuePair `json:"age,omitempty"`
  /*
  Food intake measurement (time-series data as array of time-value pairs)
  */
  FoodIntake []DoubleTimeValuePair `json:"foodIntake,omitempty"`
  /*
  Heart rate measurement (time-series data as array of time-value pairs)
  */
  HeartRate []IntegerTimeValuePair `json:"heartRate,omitempty"`
  /*
  External identifier of the item this stats record belongs to (foreign key to Item.extId)
  */
  ItemExtId *string `json:"itemExtId,omitempty"`
  /*
  Unique identifier for the stats record (primary key)
  */
  StatsExtId *string `json:"statsExtId,omitempty"`
}

func (p *ItemStatsProjection) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ItemStatsProjection

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

func (p *ItemStatsProjection) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ItemStatsProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItemStatsProjection()

    if known.ObjectType_ != nil {
        p.ObjectType_ = known.ObjectType_
    }
    if known.Reserved_ != nil {
        p.Reserved_ = known.Reserved_
    }
    if known.UnknownFields_ != nil {
        p.UnknownFields_ = known.UnknownFields_
    }
    if known.Age != nil {
        p.Age = known.Age
    }
    if known.FoodIntake != nil {
        p.FoodIntake = known.FoodIntake
    }
    if known.HeartRate != nil {
        p.HeartRate = known.HeartRate
    }
    if known.ItemExtId != nil {
        p.ItemExtId = known.ItemExtId
    }
    if known.StatsExtId != nil {
        p.StatsExtId = known.StatsExtId
    }

    // Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "age")
	delete(allFields, "foodIntake")
	delete(allFields, "heartRate")
	delete(allFields, "itemExtId")
	delete(allFields, "statsExtId")

    // Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
      p.UnknownFields_[key] = value
    }

	return nil
}

func NewItemStatsProjection() *ItemStatsProjection {
  p := new(ItemStatsProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.stats.ItemStatsProjection"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




type ItemStatsTimeValuePair struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  TimeStamp *int64 `json:"timeStamp,omitempty"`
  
  Value *int64 `json:"value,omitempty"`
}

func (p *ItemStatsTimeValuePair) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ItemStatsTimeValuePair

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

func (p *ItemStatsTimeValuePair) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ItemStatsTimeValuePair
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewItemStatsTimeValuePair()

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

func NewItemStatsTimeValuePair() *ItemStatsTimeValuePair {
  p := new(ItemStatsTimeValuePair)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.stats.ItemStatsTimeValuePair"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



/*
REST response for all response codes in API path /nexus/v4.1/stats/item-stats Get operation
*/
type ListItemStatsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /*
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfListItemStatsApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListItemStatsApiResponse) MarshalJSON() ([]byte, error) {
  // Create Alias to avoid infinite recursion
  type Alias ListItemStatsApiResponse

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

func (p *ListItemStatsApiResponse) UnmarshalJSON(b []byte) error {
    // Step 1: Unmarshal into a generic map to capture all fields
    var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

    // Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListItemStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

    // Step 3: Assign known fields
	*p = *NewListItemStatsApiResponse()

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

func NewListItemStatsApiResponse() *ListItemStatsApiResponse {
  p := new(ListItemStatsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "nexus.v4.stats.ListItemStatsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ListItemStatsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ListItemStatsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfListItemStatsApiResponseData()
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


type OneOfItemStatsGroupGroup struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType2005 *float64 `json:"-"`
  oneOfType2004 *int64 `json:"-"`
  oneOfType2002 *string `json:"-"`
  oneOfType2003 *int `json:"-"`
  oneOfType2006 *bool `json:"-"`
}

func NewOneOfItemStatsGroupGroup() *OneOfItemStatsGroupGroup {
  p := new(OneOfItemStatsGroupGroup)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfItemStatsGroupGroup) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfItemStatsGroupGroup is nil"))
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

func (p *OneOfItemStatsGroupGroup) GetValue() interface{} {
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

func (p *OneOfItemStatsGroupGroup) UnmarshalJSON(b []byte) error {

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
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfItemStatsGroupGroup"))
}

func (p *OneOfItemStatsGroupGroup) MarshalJSON() ([]byte, error) {
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
  return nil, errors.New("No value to marshal for OneOfItemStatsGroupGroup")
}

type OneOfItemStatsAggregateResult struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType2003 *int `json:"-"`
  oneOfType2005 *float64 `json:"-"`
  oneOfType2006 []ItemStatsTimeValuePair `json:"-"`
  oneOfType2004 *int64 `json:"-"`
}

func NewOneOfItemStatsAggregateResult() *OneOfItemStatsAggregateResult {
  p := new(OneOfItemStatsAggregateResult)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfItemStatsAggregateResult) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfItemStatsAggregateResult is nil"))
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
    case []ItemStatsTimeValuePair:
      p.oneOfType2006 = v.([]ItemStatsTimeValuePair)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.stats.ItemStatsTimeValuePair>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.stats.ItemStatsTimeValuePair>"
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

func (p *OneOfItemStatsAggregateResult) GetValue() interface{} {
  if "Integer" == *p.Discriminator {
    return *p.oneOfType2003
  }
  if "Double" == *p.Discriminator {
    return *p.oneOfType2005
  }
  if "List<nexus.v4.stats.ItemStatsTimeValuePair>" == *p.Discriminator {
    return p.oneOfType2006
  }
  if "Long" == *p.Discriminator {
    return *p.oneOfType2004
  }
  return nil
}

func (p *OneOfItemStatsAggregateResult) UnmarshalJSON(b []byte) error {

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
    if nestedMap, ok := rawMap["List<nexus.v4.stats.ItemStatsTimeValuePair>"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2006 := new([]ItemStatsTimeValuePair)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2006)
          if unmarshalErr == nil {
            // For arrays, verify the array item ObjectType matches
            if vOneOfType2006 == nil || len(*vOneOfType2006) == 0 || ((*vOneOfType2006)[0].ObjectType_ != nil && "nexus.v4.stats.ItemStatsTimeValuePair" == *((*vOneOfType2006)[0].ObjectType_)) {
              p.oneOfType2006 = *vOneOfType2006
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "List<nexus.v4.stats.ItemStatsTimeValuePair>"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "List<nexus.v4.stats.ItemStatsTimeValuePair>"
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
  vOneOfType2006 := new([]ItemStatsTimeValuePair)
  if err := json.Unmarshal(b, vOneOfType2006); err == nil {
    if len(*vOneOfType2006) == 0 || (vOneOfType2006 != nil && (*vOneOfType2006)[0].ObjectType_ != nil && "nexus.v4.stats.ItemStatsTimeValuePair" == *((*vOneOfType2006)[0].ObjectType_)) {
      p.oneOfType2006 = *vOneOfType2006
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.stats.ItemStatsTimeValuePair>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.stats.ItemStatsTimeValuePair>"
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
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfItemStatsAggregateResult"))
}

func (p *OneOfItemStatsAggregateResult) MarshalJSON() ([]byte, error) {
  if "Integer" == *p.Discriminator {
    return json.Marshal(p.oneOfType2003)
  }
  if "Double" == *p.Discriminator {
    return json.Marshal(p.oneOfType2005)
  }
  if "List<nexus.v4.stats.ItemStatsTimeValuePair>" == *p.Discriminator {
    return json.Marshal(p.oneOfType2006)
  }
  if "Long" == *p.Discriminator {
    return json.Marshal(p.oneOfType2004)
  }
  return nil, errors.New("No value to marshal for OneOfItemStatsAggregateResult")
}

type OneOfItemStatsGroupData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType2008 []ItemStats `json:"-"`
}

func NewOneOfItemStatsGroupData() *OneOfItemStatsGroupData {
  p := new(OneOfItemStatsGroupData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfItemStatsGroupData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfItemStatsGroupData is nil"))
  }
  switch v.(type) {
    case []ItemStats:
      p.oneOfType2008 = v.([]ItemStats)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.stats.ItemStats>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.stats.ItemStats>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfItemStatsGroupData) GetValue() interface{} {
  if "List<nexus.v4.stats.ItemStats>" == *p.Discriminator {
    return p.oneOfType2008
  }
  return nil
}

func (p *OneOfItemStatsGroupData) UnmarshalJSON(b []byte) error {

  // Try to handle nested structure like {"": {"value": {...}}}
  // This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
  var rawMap map[string]interface{}
  if err := json.Unmarshal(b, &rawMap); err == nil {
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["List<nexus.v4.stats.ItemStats>"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2008 := new([]ItemStats)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2008)
          if unmarshalErr == nil {
            // For arrays, verify the array item ObjectType matches
            if vOneOfType2008 == nil || len(*vOneOfType2008) == 0 || ((*vOneOfType2008)[0].ObjectType_ != nil && "nexus.v4.stats.ItemStats" == *((*vOneOfType2008)[0].ObjectType_)) {
              p.oneOfType2008 = *vOneOfType2008
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "List<nexus.v4.stats.ItemStats>"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "List<nexus.v4.stats.ItemStats>"
              return nil
            }
          }
        }
      }
    }
  }

  // Fallback: try direct unmarshalling (for non-nested structures)
  vOneOfType2008 := new([]ItemStats)
  if err := json.Unmarshal(b, vOneOfType2008); err == nil {
    if len(*vOneOfType2008) == 0 || (vOneOfType2008 != nil && (*vOneOfType2008)[0].ObjectType_ != nil && "nexus.v4.stats.ItemStats" == *((*vOneOfType2008)[0].ObjectType_)) {
      p.oneOfType2008 = *vOneOfType2008
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.stats.ItemStats>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.stats.ItemStats>"
      return nil
    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfItemStatsGroupData"))
}

func (p *OneOfItemStatsGroupData) MarshalJSON() ([]byte, error) {
  if "List<nexus.v4.stats.ItemStats>" == *p.Discriminator {
    return json.Marshal(p.oneOfType2008)
  }
  return nil, errors.New("No value to marshal for OneOfItemStatsGroupData")
}

type OneOfListItemStatsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType402 []ItemStatsGroup `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType2001 []ItemStats `json:"-"`
  oneOfType401 []ItemStatsProjection `json:"-"`
}

func NewOneOfListItemStatsApiResponseData() *OneOfListItemStatsApiResponseData {
  p := new(OneOfListItemStatsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfListItemStatsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfListItemStatsApiResponseData is nil"))
  }
  switch v.(type) {
    case []ItemStatsGroup:
      p.oneOfType402 = v.([]ItemStatsGroup)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.stats.ItemStatsGroup>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.stats.ItemStatsGroup>"
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []ItemStats:
      p.oneOfType2001 = v.([]ItemStats)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.stats.ItemStats>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.stats.ItemStats>"
    case []ItemStatsProjection:
      p.oneOfType401 = v.([]ItemStatsProjection)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.stats.ItemStatsProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.stats.ItemStatsProjection>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfListItemStatsApiResponseData) GetValue() interface{} {
  if "List<nexus.v4.stats.ItemStatsGroup>" == *p.Discriminator {
    return p.oneOfType402
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<nexus.v4.stats.ItemStats>" == *p.Discriminator {
    return p.oneOfType2001
  }
  if "List<nexus.v4.stats.ItemStatsProjection>" == *p.Discriminator {
    return p.oneOfType401
  }
  return nil
}

func (p *OneOfListItemStatsApiResponseData) UnmarshalJSON(b []byte) error {

  // Try to handle nested structure like {"": {"value": {...}}}
  // This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
  var rawMap map[string]interface{}
  if err := json.Unmarshal(b, &rawMap); err == nil {
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["List<nexus.v4.stats.ItemStatsGroup>"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType402 := new([]ItemStatsGroup)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType402)
          if unmarshalErr == nil {
            // For arrays, verify the array item ObjectType matches
            if vOneOfType402 == nil || len(*vOneOfType402) == 0 || ((*vOneOfType402)[0].ObjectType_ != nil && "nexus.v4.stats.ItemStatsGroup" == *((*vOneOfType402)[0].ObjectType_)) {
              p.oneOfType402 = *vOneOfType402
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "List<nexus.v4.stats.ItemStatsGroup>"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "List<nexus.v4.stats.ItemStatsGroup>"
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
          vOneOfType400 := new(import2.ErrorResponse)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
          if unmarshalErr == nil {
            // For struct items, verify the ObjectType matches
            if vOneOfType400.ObjectType_ != nil && "nexus.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
              if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
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
    if nestedMap, ok := rawMap["List<nexus.v4.stats.ItemStats>"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType2001 := new([]ItemStats)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
          if unmarshalErr == nil {
            // For arrays, verify the array item ObjectType matches
            if vOneOfType2001 == nil || len(*vOneOfType2001) == 0 || ((*vOneOfType2001)[0].ObjectType_ != nil && "nexus.v4.stats.ItemStats" == *((*vOneOfType2001)[0].ObjectType_)) {
              p.oneOfType2001 = *vOneOfType2001
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "List<nexus.v4.stats.ItemStats>"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "List<nexus.v4.stats.ItemStats>"
              return nil
            }
          }
        }
      }
    }
    // Check if this field name exists in the map (handles nested structure)
    if nestedMap, ok := rawMap["List<nexus.v4.stats.ItemStatsProjection>"].(map[string]interface{}); ok {
      // Check for "value" wrapper
      if valueData, ok := nestedMap["value"]; ok {
        valueJSON, marshalErr := json.Marshal(valueData)
        if marshalErr == nil {
          vOneOfType401 := new([]ItemStatsProjection)
          var unmarshalErr error
          // Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
          unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
          if unmarshalErr == nil {
            // For arrays, verify the array item ObjectType matches
            if vOneOfType401 == nil || len(*vOneOfType401) == 0 || ((*vOneOfType401)[0].ObjectType_ != nil && "nexus.v4.stats.ItemStatsProjection" == *((*vOneOfType401)[0].ObjectType_)) {
              p.oneOfType401 = *vOneOfType401
              if nil == p.Discriminator {p.Discriminator = new(string)}
              *p.Discriminator = "List<nexus.v4.stats.ItemStatsProjection>"
              if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
              *p.ObjectType_ = "List<nexus.v4.stats.ItemStatsProjection>"
              return nil
            }
          }
        }
      }
    }
  }

  // Fallback: try direct unmarshalling (for non-nested structures)
  vOneOfType402 := new([]ItemStatsGroup)
  if err := json.Unmarshal(b, vOneOfType402); err == nil {
    if len(*vOneOfType402) == 0 || (vOneOfType402 != nil && (*vOneOfType402)[0].ObjectType_ != nil && "nexus.v4.stats.ItemStatsGroup" == *((*vOneOfType402)[0].ObjectType_)) {
      p.oneOfType402 = *vOneOfType402
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.stats.ItemStatsGroup>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.stats.ItemStatsGroup>"
      return nil
    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if vOneOfType400.ObjectType_ != nil && "nexus.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
  }
  vOneOfType2001 := new([]ItemStats)
  if err := json.Unmarshal(b, vOneOfType2001); err == nil {
    if len(*vOneOfType2001) == 0 || (vOneOfType2001 != nil && (*vOneOfType2001)[0].ObjectType_ != nil && "nexus.v4.stats.ItemStats" == *((*vOneOfType2001)[0].ObjectType_)) {
      p.oneOfType2001 = *vOneOfType2001
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.stats.ItemStats>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.stats.ItemStats>"
      return nil
    }
  }
  vOneOfType401 := new([]ItemStatsProjection)
  if err := json.Unmarshal(b, vOneOfType401); err == nil {
    if len(*vOneOfType401) == 0 || (vOneOfType401 != nil && (*vOneOfType401)[0].ObjectType_ != nil && "nexus.v4.stats.ItemStatsProjection" == *((*vOneOfType401)[0].ObjectType_)) {
      p.oneOfType401 = *vOneOfType401
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<nexus.v4.stats.ItemStatsProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<nexus.v4.stats.ItemStatsProjection>"
      return nil
    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListItemStatsApiResponseData"))
}

func (p *OneOfListItemStatsApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<nexus.v4.stats.ItemStatsGroup>" == *p.Discriminator {
    return json.Marshal(p.oneOfType402)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<nexus.v4.stats.ItemStats>" == *p.Discriminator {
    return json.Marshal(p.oneOfType2001)
  }
  if "List<nexus.v4.stats.ItemStatsProjection>" == *p.Discriminator {
    return json.Marshal(p.oneOfType401)
  }
  return nil, errors.New("No value to marshal for OneOfListItemStatsApiResponseData")
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

