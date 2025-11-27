# OData Parser Implementation - M2 Milestone

## ✅ Implementation Complete

All M2 milestone tasks have been implemented and aligned with az-manager:

1. ✅ **Create an OData query parser** (2 points)
2. ✅ **Create an exception handler for the query parser** (1 point)
3. ✅ **Fetch data from IDF for LIST API** (2 points) - Enhanced with OData
4. ✅ **Map IDF call response to the model schema attributes response/DTO** (2 points)

---

## Implementation Details

### 1. OData Query Parser (`idf/odata_parser.go`)

**Function:** `GenerateListQuery()`

**What it does:**
- Parses OData query parameters (`$filter`, `$orderby`, `$select`, `$expand`)
- Converts OData expressions to IDF query clauses
- Maps OData field names (camelCase) to IDF attribute names (snake_case)
- Generates IDF `WhereClause` from `$filter`
- Generates IDF `GroupSortOrder` from `$orderby`
- Handles pagination (`$page`, `$limit`)

**Aligned with az-manager:**
- Uses same packages: `ntnx-api-odata-go`
- Same pattern: `edm.NewCustomEdmProvider()` → `parser.NewParser()` → `IDFQueryEvaluator`
- Same structure: `GenerateListQuery()` → `constructIDFQuery()`

**Example usage:**
```go
// OData query: $filter=itemName eq 'test'&$orderby=itemId asc&$page=0&$limit=10
queryArg, err := GenerateListQuery(queryParams, "/items", "item", "item_id")
// Returns IDF query with WhereClause and GroupSortOrder populated
```

---

### 2. Exception Handler (`utils/odata/odata_error_handler.go`)

**Functions:**
- `HandleODataParseError()` - Handles parsing errors
- `HandleODataEvaluationError()` - Handles evaluation errors
- `ValidateODataQueryParams()` - Validates query parameters
- `WrapODataError()` - Wraps errors with context

**What it does:**
- Converts OData parsing errors to user-friendly gRPC status errors
- Returns appropriate HTTP status codes (400 Bad Request for invalid queries)
- Provides clear error messages for:
  - Invalid syntax
  - Unknown properties
  - Unsupported operators
  - Type mismatches

**Integration:**
- Called from `item_grpc_service.go` → `handleODataError()`
- Returns gRPC status errors with proper codes

---

### 3. EDM Bindings (`idf/odata_parser.go` → `GetNexusEntityBindings()`)

**What it does:**
- Creates EDM (Entity Data Model) bindings for Item entity
- Maps OData field names → IDF column names:
  - `itemId` → `item_id`
  - `itemName` → `item_name`
  - `itemType` → `item_type`
  - `description` → `description`
  - `extId` → `ext_id`
- Defines filterable and sortable properties
- Sets property types (EdmString, EdmInt64)

**Note:** In a full implementation, these would be auto-generated from YAML definitions (like az-manager does). For now, we create them manually.

---

### 4. Integration with IDF ListItems

**Before (basic):**
```go
// Manual query building
query, err := idfQr.QUERY("itemListQuery").FROM("item").Proto()
query.GroupBy.RawColumns = [...] // Manual column selection
query.GroupBy.RawLimit = &QueryLimit{...} // Manual pagination
```

**After (with OData):**
```go
// OData parser handles everything
queryArg, err := GenerateListQuery(queryParams, itemListPath, itemEntityTypeName, itemIdAttr)
// Automatically:
// - Parses $filter → WhereClause
// - Parses $orderby → GroupSortOrder
// - Parses $select → RawColumns
// - Handles pagination
```

---

## Supported OData Features

### ✅ Fully Supported:
- **$filter** - Filter expressions (e.g., `itemName eq 'test'`, `itemId gt 10`)
- **$orderby** - Sorting (e.g., `itemName asc, itemId desc`)
- **$select** - Field selection (e.g., `itemId,itemName`)
- **$page** - Pagination page number
- **$limit** - Page size

### ⚠️ Partially Supported:
- **$expand** - Parsed but not fully evaluated (GraphQL evaluator not implemented)

---

## Error Handling

**Invalid OData Query Examples:**

1. **Invalid syntax:**
   ```
   $filter=itemName eq
   → Error: Invalid OData query syntax in '$filter=itemName eq': syntax error
   ```

2. **Unknown property:**
   ```
   $filter=unknownField eq 'test'
   → Error: Unknown property in OData query: property 'unknownField' not found
   ```

3. **Unsupported operator:**
   ```
   $filter=itemName contains 'test'
   → Error: Unsupported operator in OData query: operator 'contains' not supported
   ```

---

## Code Structure

```
golang-nexus-service/
├── idf/
│   ├── odata_parser.go          ← OData parser (GenerateListQuery)
│   ├── idf_item_repository.go   ← Uses OData parser
│   └── idf_utils.go             ← Utility functions
├── utils/
│   └── odata/
│       └── odata_error_handler.go ← Error handling
└── grpc/
    └── item_grpc_service.go    ← Error handling integration
```

---

## Dependencies Added

```go
github.com/nutanix-core/ntnx-api-odata-go v1.0.27
```

**Packages used:**
- `odata/edm` - EDM provider and bindings
- `odata/uri/parser` - OData parser
- `db/idf` - IDF query evaluator

---

## Testing Examples

### Example 1: Filter by itemName
```
GET /api/nexus/v4.1/config/items?$filter=itemName eq 'test item 0'
```
- OData parser converts `itemName` → `item_name`
- Generates IDF `WhereClause` with equality condition
- Returns filtered results

### Example 2: Sort by itemId
```
GET /api/nexus/v4.1/config/items?$orderby=itemId asc
```
- OData parser converts `itemId` → `item_id`
- Generates IDF `GroupSortOrder`
- Returns sorted results

### Example 3: Combined filter and sort
```
GET /api/nexus/v4.1/config/items?$filter=itemType eq 'TYPE1'&$orderby=itemName asc&$page=0&$limit=10
```
- Parses both `$filter` and `$orderby`
- Applies pagination
- Returns filtered, sorted, paginated results

---

## Alignment with az-manager

✅ **Same packages:** `ntnx-api-odata-go`  
✅ **Same pattern:** EDM Provider → Parser → IDF Evaluator  
✅ **Same structure:** `GenerateListQuery()` → `constructIDFQuery()`  
✅ **Same error handling:** Wrapped errors with context  
✅ **Same query building:** Uses IDF query builder and evaluator  

**Differences (by design):**
- EDM bindings are manually created (az-manager has auto-generated)
- No GraphQL evaluator (not needed for basic queries)
- Simplified for Item entity only

---

## Next Steps (Future Enhancements)

1. **Auto-generate EDM bindings** from YAML definitions (like az-manager)
2. **Add GraphQL evaluator** for `$expand` support
3. **Add more OData functions** (startswith, endswith, etc.)
4. **Add validation** for query parameter limits

---

## Summary

✅ **M2 Milestone Complete:**
- OData parser implemented and aligned with az-manager
- Exception handler provides user-friendly errors
- IDF fetching enhanced with OData support
- Response mapping already working (from previous milestone)

**Total Points:** 7/7 ✅

