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
  // set Edm Property Mapping
  p.PropertyMappings = make(map[string]string)
  p.PropertyMappings["itemId"] = "item_id"
  p.PropertyMappings["itemName"] = "item_name"
  p.PropertyMappings["itemType"] = "item_type"
  p.PropertyMappings["quantity"] = "quantity"
  p.PropertyMappings["price"] = "price"
  p.PropertyMappings["int64List"] = "int64_list"
  p.PropertyMappings["description"] = "description"
  p.PropertyMappings["extId"] = "ext_id"
  p.PropertyMappings["isActive"] = "is_active"
  p.PropertyMappings["priority"] = "priority"
  p.PropertyMappings["status"] = "status"

  filterProperties := make(map[string]bool)
  // set filterable properties in a map
  filterProperties["itemId"] = true
  filterProperties["itemName"] = true
  filterProperties["itemType"] = true
  filterProperties["extId"] = true
  filterProperties["quantity"] = true
  filterProperties["price"] = true
  filterProperties["isActive"] = true
  filterProperties["status"] = true
  filterProperties["int64List"] = true

  sortableProperties := make(map[string]bool)
  // set sortable properties in a map
  sortableProperties["itemId"] = true
  sortableProperties["itemName"] = true
  sortableProperties["itemType"] = true
  sortableProperties["quantity"] = true
  sortableProperties["price"] = true
  sortableProperties["priority"] = true

  groupableProperties := make(map[string]bool)
  // set groupable properties in a map
  groupableProperties["itemId"] = true
  groupableProperties["itemName"] = true
  groupableProperties["itemType"] = true
  groupableProperties["description"] = true
  groupableProperties["extId"] = true
  groupableProperties["quantity"] = true
  groupableProperties["price"] = true
  groupableProperties["isActive"] = true
  groupableProperties["priority"] = true
  groupableProperties["status"] = true
  groupableProperties["int64List"] = true

  // set Edm Properties
  var properties []*edm.EdmProperty
  itemIdProperty := new(edm.EdmProperty)
  itemIdProperty.Name = "itemId"
  itemIdProperty.IsCollection = false
  itemIdProperty.Type = string(edm.EdmInt32)
  itemIdProperty.MappedName = p.PropertyMappings["itemId"]
  itemIdProperty.IsFilterable = filterProperties["itemId"]
  itemIdProperty.IsSortable = sortableProperties["itemId"]
  itemIdProperty.IsGroupable = groupableProperties["itemId"]
  properties = append(properties, itemIdProperty)

  itemNameProperty := new(edm.EdmProperty)
  itemNameProperty.Name = "itemName"
  itemNameProperty.IsCollection = false
  itemNameProperty.Type = string(edm.EdmString)
  itemNameProperty.MappedName = p.PropertyMappings["itemName"]
  itemNameProperty.IsFilterable = filterProperties["itemName"]
  itemNameProperty.IsSortable = sortableProperties["itemName"]
  itemNameProperty.IsGroupable = groupableProperties["itemName"]
  properties = append(properties, itemNameProperty)

  itemTypeProperty := new(edm.EdmProperty)
  itemTypeProperty.Name = "itemType"
  itemTypeProperty.IsCollection = false
  itemTypeProperty.Type = string(edm.EdmString)
  itemTypeProperty.MappedName = p.PropertyMappings["itemType"]
  itemTypeProperty.IsFilterable = filterProperties["itemType"]
  itemTypeProperty.IsSortable = sortableProperties["itemType"]
  itemTypeProperty.IsGroupable = groupableProperties["itemType"]
  properties = append(properties, itemTypeProperty)

  descriptionProperty := new(edm.EdmProperty)
  descriptionProperty.Name = "description"
  descriptionProperty.IsCollection = false
  descriptionProperty.Type = string(edm.EdmString)
  descriptionProperty.MappedName = p.PropertyMappings["description"]
  descriptionProperty.IsFilterable = filterProperties["description"]
  descriptionProperty.IsSortable = sortableProperties["description"]
  descriptionProperty.IsGroupable = groupableProperties["description"]
  properties = append(properties, descriptionProperty)

  extIdProperty := new(edm.EdmProperty)
  extIdProperty.Name = "extId"
  extIdProperty.IsCollection = false
  extIdProperty.Type = string(edm.EdmString)
  extIdProperty.MappedName = p.PropertyMappings["extId"]
  extIdProperty.IsFilterable = filterProperties["extId"]
  extIdProperty.IsSortable = sortableProperties["extId"]
  extIdProperty.IsGroupable = groupableProperties["extId"]
  properties = append(properties, extIdProperty)

  quantityProperty := new(edm.EdmProperty)
  quantityProperty.Name = "quantity"
  quantityProperty.IsCollection = false
  quantityProperty.Type = string(edm.EdmInt64)
  quantityProperty.MappedName = p.PropertyMappings["quantity"]
  quantityProperty.IsFilterable = filterProperties["quantity"]
  quantityProperty.IsSortable = sortableProperties["quantity"]
  quantityProperty.IsGroupable = groupableProperties["quantity"]
  properties = append(properties, quantityProperty)

  priceProperty := new(edm.EdmProperty)
  priceProperty.Name = "price"
  priceProperty.IsCollection = false
  priceProperty.Type = string(edm.EdmDouble)
  priceProperty.MappedName = p.PropertyMappings["price"]
  priceProperty.IsFilterable = filterProperties["price"]
  priceProperty.IsSortable = sortableProperties["price"]
  priceProperty.IsGroupable = groupableProperties["price"]
  properties = append(properties, priceProperty)

  isActiveProperty := new(edm.EdmProperty)
  isActiveProperty.Name = "isActive"
  isActiveProperty.IsCollection = false
  isActiveProperty.Type = string(edm.EdmBoolean)
  isActiveProperty.MappedName = p.PropertyMappings["isActive"]
  isActiveProperty.IsFilterable = filterProperties["isActive"]
  isActiveProperty.IsSortable = sortableProperties["isActive"]
  isActiveProperty.IsGroupable = groupableProperties["isActive"]
  properties = append(properties, isActiveProperty)

  priorityProperty := new(edm.EdmProperty)
  priorityProperty.Name = "priority"
  priorityProperty.IsCollection = false
  priorityProperty.Type = string(edm.EdmInt32)
  priorityProperty.MappedName = p.PropertyMappings["priority"]
  priorityProperty.IsFilterable = filterProperties["priority"]
  priorityProperty.IsSortable = sortableProperties["priority"]
  priorityProperty.IsGroupable = groupableProperties["priority"]
  properties = append(properties, priorityProperty)

  statusProperty := new(edm.EdmProperty)
  statusProperty.Name = "status"
  statusProperty.IsCollection = false
  statusProperty.Type = string(edm.EdmString)
  statusProperty.MappedName = p.PropertyMappings["status"]
  statusProperty.IsFilterable = filterProperties["status"]
  statusProperty.IsSortable = sortableProperties["status"]
  statusProperty.IsGroupable = groupableProperties["status"]
  properties = append(properties, statusProperty)

  int64ListProperty := new(edm.EdmProperty)
  int64ListProperty.Name = "int64List"
  int64ListProperty.IsCollection = true
  int64ListProperty.Type = string(edm.EdmInt64)
  int64ListProperty.MappedName = p.PropertyMappings["int64List"]
  int64ListProperty.IsFilterable = filterProperties["int64List"]
  int64ListProperty.IsSortable = sortableProperties["int64List"]
  int64ListProperty.IsGroupable = groupableProperties["int64List"]
  properties = append(properties, int64ListProperty)



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
