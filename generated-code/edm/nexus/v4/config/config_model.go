/*
 * Generated file models/edm/nexus/v4/config/config_model.go.
 *
 * Product version: 1.0.0-SNAPSHOT
 *
 * Part of the GoLang Mock API - REST API for Mock Item Service
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

package config

import (
  "github.com/nutanix-core/ntnx-api-odata-go/odata/edm"
)


func NewItem() *edm.EdmEntityBinding {

  p := new(edm.EdmEntityBinding)

  filterProperties := make(map[string]bool)
  // set filterable properties in a map
  filterProperties["itemId"] = true
  filterProperties["itemName"] = true
  filterProperties["itemType"] = true
  filterProperties["extId"] = true

  sortableProperties := make(map[string]bool)
  // set sortable properties in a map
  sortableProperties["itemId"] = true
  sortableProperties["itemName"] = true
  sortableProperties["itemType"] = true


  // set Edm Properties
  var properties []*edm.EdmProperty
  itemIdProperty := new(edm.EdmProperty)
  itemIdProperty.Name = "itemId"
  itemIdProperty.IsCollection = false
  itemIdProperty.Type = string(edm.EdmInt32)
  itemIdProperty.IsFilterable = filterProperties["itemId"]
  itemIdProperty.IsSortable = sortableProperties["itemId"]
  properties = append(properties, itemIdProperty)

  itemNameProperty := new(edm.EdmProperty)
  itemNameProperty.Name = "itemName"
  itemNameProperty.IsCollection = false
  itemNameProperty.Type = string(edm.EdmString)
  itemNameProperty.IsFilterable = filterProperties["itemName"]
  itemNameProperty.IsSortable = sortableProperties["itemName"]
  properties = append(properties, itemNameProperty)

  itemTypeProperty := new(edm.EdmProperty)
  itemTypeProperty.Name = "itemType"
  itemTypeProperty.IsCollection = false
  itemTypeProperty.Type = string(edm.EdmString)
  itemTypeProperty.IsFilterable = filterProperties["itemType"]
  itemTypeProperty.IsSortable = sortableProperties["itemType"]
  properties = append(properties, itemTypeProperty)

  descriptionProperty := new(edm.EdmProperty)
  descriptionProperty.Name = "description"
  descriptionProperty.IsCollection = false
  descriptionProperty.Type = string(edm.EdmString)
  descriptionProperty.IsFilterable = filterProperties["description"]
  descriptionProperty.IsSortable = sortableProperties["description"]
  properties = append(properties, descriptionProperty)

  extIdProperty := new(edm.EdmProperty)
  extIdProperty.Name = "extId"
  extIdProperty.IsCollection = false
  extIdProperty.Type = string(edm.EdmString)
  extIdProperty.IsFilterable = filterProperties["extId"]
  extIdProperty.IsSortable = sortableProperties["extId"]
  properties = append(properties, extIdProperty)



  // set Edm Entity Type
  entityType := new(edm.EdmEntityType)
  entityType.Name = "item"
  entityType.Properties = properties
  p.EntityType = entityType

  // set Edm Entity Set
  entitySet := new(edm.EdmEntitySet)
  entitySet.Name = "items"
  entitySet.EntityType = edm.GetFullQualifiedName(edm.NamespaceEntities, "item")
  entitySet.IncludeInServiceDocument = true
  entitySet.TableName = "item"
  p.EntitySet = entitySet


  p.RbacEntityName = "Item"



  return p
}


func NewItemAssociation() *edm.EdmEntityBinding {

  p := new(edm.EdmEntityBinding)
  // set Edm Property Mapping
  p.PropertyMappings = make(map[string]string)
  p.PropertyMappings["itemId"] = "item_id"
  p.PropertyMappings["entityType"] = "entity_type"
  p.PropertyMappings["count"] = "count"
  p.PropertyMappings["entityId"] = "entity_id"

  filterProperties := make(map[string]bool)
  // set filterable properties in a map
  filterProperties["entityType"] = true
  filterProperties["count"] = true



  // set Edm Properties
  var properties []*edm.EdmProperty
  itemIdProperty := new(edm.EdmProperty)
  itemIdProperty.Name = "itemId"
  itemIdProperty.IsCollection = false
  itemIdProperty.Type = string(edm.EdmString)
  itemIdProperty.MappedName = p.PropertyMappings["itemId"]
  itemIdProperty.IsFilterable = filterProperties["itemId"]
  properties = append(properties, itemIdProperty)

  entityTypeProperty := new(edm.EdmProperty)
  entityTypeProperty.Name = "entityType"
  entityTypeProperty.IsCollection = false
  entityTypeProperty.Type = string(edm.EdmString)
  entityTypeProperty.MappedName = p.PropertyMappings["entityType"]
  entityTypeProperty.IsFilterable = filterProperties["entityType"]
  properties = append(properties, entityTypeProperty)

  entityIdProperty := new(edm.EdmProperty)
  entityIdProperty.Name = "entityId"
  entityIdProperty.IsCollection = false
  entityIdProperty.Type = string(edm.EdmString)
  entityIdProperty.MappedName = p.PropertyMappings["entityId"]
  entityIdProperty.IsFilterable = filterProperties["entityId"]
  properties = append(properties, entityIdProperty)

  countProperty := new(edm.EdmProperty)
  countProperty.Name = "count"
  countProperty.IsCollection = false
  countProperty.Type = string(edm.EdmInt32)
  countProperty.MappedName = p.PropertyMappings["count"]
  countProperty.IsFilterable = filterProperties["count"]
  properties = append(properties, countProperty)



  // set Edm Entity Type
  entityType := new(edm.EdmEntityType)
  entityType.Name = "itemassociation"
  entityType.Properties = properties
  p.EntityType = entityType

  // set Edm Entity Set
  entitySet := new(edm.EdmEntitySet)
  entitySet.Name = "itemassociationSet"
  entitySet.EntityType = edm.GetFullQualifiedName(edm.NamespaceEntities, "itemassociation")
  entitySet.IncludeInServiceDocument = true
  entitySet.TableName = "item_associations"
  p.EntitySet = entitySet


  p.RbacEntityName = "ItemAssociation"



  return p
}


// Get all the entity bindings of module config
func GetAllEntityBindings() []*edm.EdmEntityBinding {
  var entityBindingList []*edm.EdmEntityBinding
  entityBindingList = append(entityBindingList, NewItem())
  entityBindingList = append(entityBindingList, NewItemAssociation())
  return entityBindingList
}
