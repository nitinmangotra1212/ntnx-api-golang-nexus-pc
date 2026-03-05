/*
 * Generated file models/edm/nexus/v4/stats/stats_model.go.
 *
 * Product version: 1.0.0-SNAPSHOT
 *
 * Part of the GoLang Mock API - REST API for Mock Item Service
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

package stats

import (
  "github.com/nutanix-core/ntnx-api-odata-go/odata/edm"
)


func NewItemStats() *edm.EdmEntityBinding {

  p := new(edm.EdmEntityBinding)
  // set Edm Property Mapping
  p.PropertyMappings = make(map[string]string)
  p.PropertyMappings["heartRate"] = "heart_rate"
  p.PropertyMappings["foodIntake"] = "food_intake"
  p.PropertyMappings["age"] = "age"

  filterProperties := make(map[string]bool)
  // set filterable properties in a map
  filterProperties["age"] = true
  filterProperties["heartRate"] = true
  filterProperties["foodIntake"] = true

  sortableProperties := make(map[string]bool)
  // set sortable properties in a map
  sortableProperties["age"] = true
  sortableProperties["heartRate"] = true
  sortableProperties["foodIntake"] = true

  groupableProperties := make(map[string]bool)
  // set groupable properties in a map
  groupableProperties["age"] = true
  groupableProperties["heartRate"] = true
  groupableProperties["foodIntake"] = true

  // set Edm Properties
  var properties []*edm.EdmProperty
  itemExtIdProperty := new(edm.EdmProperty)
  itemExtIdProperty.Name = "itemExtId"
  itemExtIdProperty.IsCollection = false
  itemExtIdProperty.Type = string(edm.EdmString)
  itemExtIdProperty.MappedName = p.PropertyMappings["itemExtId"]
  itemExtIdProperty.IsFilterable = filterProperties["itemExtId"]
  itemExtIdProperty.IsSortable = sortableProperties["itemExtId"]
  itemExtIdProperty.IsGroupable = groupableProperties["itemExtId"]
  properties = append(properties, itemExtIdProperty)

  ageProperty := new(edm.EdmProperty)
  ageProperty.Name = "age"
  ageProperty.IsCollection = true
  ageProperty.Type = string(edm.EdmInt64)
  ageProperty.MappedName = p.PropertyMappings["age"]
  ageProperty.IsFilterable = filterProperties["age"]
  ageProperty.IsSortable = sortableProperties["age"]
  ageProperty.IsGroupable = groupableProperties["age"]
  properties = append(properties, ageProperty)

  heartRateProperty := new(edm.EdmProperty)
  heartRateProperty.Name = "heartRate"
  heartRateProperty.IsCollection = true
  heartRateProperty.Type = string(edm.EdmInt64)
  heartRateProperty.MappedName = p.PropertyMappings["heartRate"]
  heartRateProperty.IsFilterable = filterProperties["heartRate"]
  heartRateProperty.IsSortable = sortableProperties["heartRate"]
  heartRateProperty.IsGroupable = groupableProperties["heartRate"]
  properties = append(properties, heartRateProperty)

  foodIntakeProperty := new(edm.EdmProperty)
  foodIntakeProperty.Name = "foodIntake"
  foodIntakeProperty.IsCollection = true
  foodIntakeProperty.Type = string(edm.EdmInt64)
  foodIntakeProperty.MappedName = p.PropertyMappings["foodIntake"]
  foodIntakeProperty.IsFilterable = filterProperties["foodIntake"]
  foodIntakeProperty.IsSortable = sortableProperties["foodIntake"]
  foodIntakeProperty.IsGroupable = groupableProperties["foodIntake"]
  properties = append(properties, foodIntakeProperty)



  // set Edm Entity Type
  entityType := new(edm.EdmEntityType)
  entityType.Name = "nexus"+"stats"+"itemstat"
  entityType.Properties = properties
  p.EntityType = entityType

  // set Edm Entity Set
  entitySet := new(edm.EdmEntitySet)
  entitySet.Name = "nexus"+"stats"+"itemstats"
  entitySet.EntityType = edm.GetFullQualifiedName(edm.NamespaceEntities, "nexus"+"stats"+"itemstat")
  entitySet.IncludeInServiceDocument = true
  entitySet.TableName = "item_stats"
  p.EntitySet = entitySet


  p.RbacEntityName = "ItemStats"


  p.Namespace = "nexus"
  p.Module = "stats"

  return p
}


// Get all the entity bindings of module stats
func GetAllEntityBindings() []*edm.EdmEntityBinding {
  var entityBindingList []*edm.EdmEntityBinding
  entityBindingList = append(entityBindingList, NewItemStats())
  return entityBindingList
}
