# How OData Parsing Was Implemented - Technical Deep Dive

## Overview

This document explains **how** OData parsing was implemented in the Nexus service, step-by-step, so you can explain it to your team.

---

## Architecture Flow

```
HTTP Request with OData Query
    ↓
Adonis (prism-service) - Extracts query params
    ↓
gRPC Call to Go Service - Forwards original path in metadata
    ↓
Go gRPC Service - Extracts OData params from context
    ↓
OData Parser - Converts OData expressions to IDF format
    ↓
IDF Query Evaluator - Generates IDF query clauses
    ↓
IDF Service - Executes query with filters/sorting
    ↓
Repository - Maps IDF entities to protobuf Items
    ↓
gRPC Response - Returns filtered/sorted results
```

---

## Step-by-Step Implementation

### Step 1: Extract OData Query Parameters

**File:** `golang-nexus-service/utils/query/query_utils.go`

**What it does:**
- Extracts OData query parameters from HTTP request
- Gets the original path from gRPC context metadata (Adonis forwards it)
- Parses URL to extract `$filter`, `$orderby`, `$select`, `$expand`, `$page`, `$limit`

**Code:**
```go
func ExtractQueryParamsFromContext(ctx context.Context) *models.QueryParams {
    // Get original HTTP path from gRPC context metadata
    path := responseUtils.GetPathFromGrpcContext(ctx)
    
    // Parse URL to extract query parameters
    parsedURL, err := url.Parse(path)
    values := parsedURL.Query()
    
    // Extract OData parameters
    queryParams.Filter = values.Get("$filter")    // e.g., "itemName eq 'test'"
    queryParams.Orderby = values.Get("$orderby")   // e.g., "itemId asc"
    queryParams.Select = values.Get("$select")     // e.g., "itemId,itemName"
    queryParams.Expand = values.Get("$expand")
    queryParams.Page = parseInt(values.Get("$page"))
    queryParams.Limit = parseInt(values.Get("$limit"))
    
    return queryParams
}
```

**Example:**
- User sends: `GET /api/nexus/v4.1/config/items?$filter=itemName eq 'test'`
- Extracted: `queryParams.Filter = "itemName eq 'test'"`

---

### Step 2: Create EDM (Entity Data Model) Bindings

**File:** `golang-nexus-service/idf/odata_parser.go` → `GetNexusEntityBindings()`

**What it does:**
- Creates EDM bindings that map OData field names (camelCase) to IDF column names (snake_case)
- Defines which properties are filterable and sortable
- Sets property types (EdmString, EdmInt64, etc.)

**Why needed:**
- OData uses camelCase: `itemName`, `itemId`
- IDF uses snake_case: `item_name`, `item_id`
- EDM bindings provide the mapping between them

**Code:**
```go
func GetNexusEntityBindings() []*edm.EdmEntityBinding {
    binding := new(edm.EdmEntityBinding)
    
    // Property mappings: OData field → IDF column
    binding.PropertyMappings = make(map[string]string)
    binding.PropertyMappings["itemName"] = "item_name"  // camelCase → snake_case
    binding.PropertyMappings["itemId"] = "item_id"
    binding.PropertyMappings["itemType"] = "item_type"
    binding.PropertyMappings["extId"] = "ext_id"
    
    // Create properties with types
    itemNameProp := new(edm.EdmProperty)
    itemNameProp.Name = "itemName"           // OData field name
    itemNameProp.Type = string(edm.EdmString)
    itemNameProp.MappedName = "item_name"     // IDF column name
    itemNameProp.IsFilterable = true          // Can use in $filter
    itemNameProp.IsSortable = true            // Can use in $orderby
    
    // ... more properties ...
    
    return []*edm.EdmEntityBinding{binding}
}
```

**Result:**
- Parser knows: `itemName` (OData) = `item_name` (IDF)
- Parser knows: `itemName` is filterable and sortable

---

### Step 3: Create EDM Provider and OData Parser

**File:** `golang-nexus-service/idf/odata_parser.go` → `GenerateListQuery()`

**What it does:**
- Creates EDM provider with entity bindings
- Creates OData parser
- Sets up query parameters from user input

**Code:**
```go
func GenerateListQuery(queryParams *models.QueryParams, resourcePath string,
    entityName string, defaultSortColumn string) (*GetEntitiesWithMetricsArg, error) {
    
    // Step 1: Get EDM bindings (field name mappings)
    entityBindingList := GetNexusEntityBindings()
    
    // Step 2: Create EDM provider
    edmProvider := edm.NewCustomEdmProvider(entityBindingList)
    
    // Step 3: Create OData parser
    odataParser := parser.NewParser(edmProvider)
    
    // Step 4: Set query parameters
    queryParam := parser.NewQueryParam()
    if queryParams.Filter != "" {
        queryParam.SetFilter(queryParams.Filter)  // "itemName eq 'test'"
    }
    if queryParams.Orderby != "" {
        queryParam.SetOrderBy(queryParams.Orderby)  // "itemId asc"
    }
    if queryParams.Select != "" {
        queryParam.SetSelect(queryParams.Select)
    }
    
    // ... continue to parsing ...
}
```

**What happens:**
- EDM provider knows how to map field names
- Parser is ready to parse OData expressions
- Query parameters are set from user input

---

### Step 4: Parse OData Expression

**File:** `golang-nexus-service/idf/odata_parser.go` → `GenerateListQuery()`

**What it does:**
- Parses the OData expression string into a structured format
- Maps OData field names to IDF column names using EDM bindings
- Validates syntax and operators

**Code:**
```go
    // Parse OData query parameters
    uriInfo, parseErr := odataParser.ParserWithQueryParam(queryParam, resourcePath)
    if parseErr != nil {
        return nil, fmt.Errorf("invalid OData query: %w", parseErr)
    }
```

**Example:**
- Input: `$filter=itemName eq 'test'`
- Parser reads: "itemName eq 'test'"
- EDM maps: `itemName` → `item_name`
- Result: Parsed expression with `item_name` (IDF column name)

**What `uriInfo` contains:**
- Parsed filter expression (with IDF column names)
- Parsed orderby expression (with IDF column names)
- Select fields (with IDF column names)
- All validated and ready for IDF

---

### Step 5: Convert to IDF Query Format

**File:** `golang-nexus-service/idf/odata_parser.go` → `GenerateListQuery()`

**What it does:**
- Converts parsed OData (uriInfo) to IDF query format
- Uses IDF Query Evaluator to generate IDF query clauses

**Code:**
```go
    // Use IDF query evaluator to convert parsed OData to IDF query
    idfQueryEval := idf.IDFQueryEvaluator{}
    idfQuery, evalErr := idfQueryEval.GetQuery(uriInfo, resourcePath)
    if evalErr != nil {
        return nil, fmt.Errorf("failed to evaluate OData query: %w", evalErr)
    }
```

**What `idfQuery` contains:**
- `WhereClause` - Filter conditions (from `$filter`)
- `GroupBy.RawSortOrder` - Sort order (from `$orderby`)
- `GroupBy.RawColumns` - Selected columns (from `$select`)

**Example:**
- Input: `$filter=itemName eq 'test'`
- `idfQuery.WhereClause` = IDF filter expression with `item_name = 'test'`

---

### Step 6: Construct Final IDF Query

**File:** `golang-nexus-service/idf/odata_parser.go` → `constructIDFQuery()`

**What it does:**
- Combines parsed OData query with pagination
- Sets up columns to fetch
- Creates final IDF query ready for execution

**Code:**
```go
func constructIDFQuery(queryParams *models.QueryParams, idfQuery *insights_interface.Query,
    entityType string, defaultSortColumn string) (*GetEntitiesWithMetricsArg, error) {
    
    // Build base query
    query, err := idfQr.QUERY(entityType + "ListQuery").
        FROM(entityType).Proto()
    
    // Use columns from OData $select (if specified)
    if idfQuery.GetGroupBy() != nil && len(idfQuery.GetGroupBy().RawColumns) > 0 {
        query.GroupBy.RawColumns = idfQuery.GetGroupBy().RawColumns
    } else {
        // Default: fetch all item columns
        query.GroupBy.RawColumns = [...] // item_id, item_name, etc.
    }
    
    // Add sorting from OData $orderby
    if idfQuery.GetGroupBy() != nil && idfQuery.GetGroupBy().GetGroupSortOrder() != nil {
        query.GroupBy.RawSortOrder = idfQuery.GetGroupBy().GetGroupSortOrder()
    }
    
    // Add pagination
    query.GroupBy.RawLimit = &QueryLimit{
        Limit:  proto.Int64(int64(limit)),
        Offset: proto.Int64(int64(offset)),
    }
    
    // Add filter from OData $filter
    query.WhereClause = idfQuery.GetWhereClause()
    
    return &GetEntitiesWithMetricsArg{Query: query}, nil
}
```

**Result:**
- Complete IDF query with:
  - Filter conditions (from `$filter`)
  - Sort order (from `$orderby`)
  - Selected columns (from `$select`)
  - Pagination (from `$page` and `$limit`)

---

### Step 7: Execute IDF Query

**File:** `golang-nexus-service/idf/idf_item_repository.go` → `ListItems()`

**What it does:**
- Calls `GenerateListQuery()` to get IDF query
- Executes query against IDF service
- Gets filtered/sorted results

**Code:**
```go
func (r *ItemRepositoryImpl) ListItems(queryParams *models.QueryParams) ([]*pb.Item, int64, error) {
    // Use OData parser to generate IDF query
    queryArg, err := GenerateListQuery(queryParams, itemListPath, itemEntityTypeName, itemIdAttr)
    if err != nil {
        return nil, 0, err
    }
    
    // Query IDF
    idfClient := external.Interfaces().IdfClient()
    queryResponse, err := idfClient.GetEntitiesWithMetricsRet(queryArg)
    
    // Convert IDF entities to protobuf Items
    entities := ConvertEntitiesWithMetricToEntities(queryResponse.GetGroupResultsList()[0].GetRawResults())
    for _, entity := range entities {
        item := r.mapIdfAttributeToItem(entity)
        items = append(items, item)
    }
    
    return items, totalCount, nil
}
```

**Result:**
- IDF returns filtered and sorted entities
- Repository maps them to protobuf Items
- Returns to gRPC service

---

### Step 8: Error Handling

**File:** `golang-nexus-service/grpc/item_grpc_service.go` → `handleODataError()`

**What it does:**
- Catches OData parsing errors
- Converts them to user-friendly gRPC status errors
- Returns appropriate HTTP status codes

**Code:**
```go
func handleODataError(err error, queryParams *models.QueryParams) error {
    errStr := err.Error()
    
    // Check for OData parsing errors
    if strings.Contains(errStr, "invalid OData query") {
        return status.Errorf(codes.InvalidArgument,
            "Invalid OData query syntax: %v. Please check your expression.", err)
    }
    
    if strings.Contains(errStr, "property") && strings.Contains(errStr, "not found") {
        return status.Errorf(codes.InvalidArgument,
            "Unknown property in OData query: %v. Please check field names.", err)
    }
    
    // ... more error types ...
}
```

**Result:**
- User gets clear error messages
- Proper HTTP status codes (400 for invalid queries)

---

## Complete Flow Example

### User Request:
```
GET /api/nexus/v4.1/config/items?$filter=itemName eq 'test'&$orderby=itemId asc&$page=0&$limit=10
```

### Step-by-Step Processing:

1. **Extract Query Params:**
   ```go
   queryParams.Filter = "itemName eq 'test'"
   queryParams.Orderby = "itemId asc"
   queryParams.Page = 0
   queryParams.Limit = 10
   ```

2. **Create EDM Bindings:**
   ```go
   PropertyMappings["itemName"] = "item_name"
   PropertyMappings["itemId"] = "item_id"
   ```

3. **Parse OData:**
   ```go
   // Parser reads: "itemName eq 'test'"
   // EDM maps: itemName → item_name
   // Result: Parsed expression with item_name
   ```

4. **Convert to IDF Query:**
   ```go
   idfQuery.WhereClause = {
     ComparisonExpr: {
       Operator: EQ,
       Lhs: { Column: "item_name" },  // ✅ Mapped!
       Rhs: { Value: "test" }
     }
   }
   idfQuery.GroupBy.RawSortOrder = [{
     Column: "item_id",  // ✅ Mapped!
     Order: ASC
   }]
   ```

5. **Execute IDF Query:**
   ```go
   // IDF executes: WHERE item_name = 'test' ORDER BY item_id ASC LIMIT 10
   // Returns: Filtered and sorted items
   ```

6. **Map to Protobuf:**
   ```go
   // IDF returns: item_name = "test", item_id = 1
   // Maps to: ItemName = "test", ItemId = 1
   ```

7. **Return Response:**
   ```json
   {
     "data": [
       {
         "itemId": 1,
         "itemName": "test",
         ...
       }
     ],
     "metadata": {...}
   }
   ```

---

## Key Components

### 1. EDM Bindings (`GetNexusEntityBindings()`)
- **Purpose:** Maps OData field names → IDF column names
- **Location:** `idf/odata_parser.go`
- **Why:** OData uses camelCase, IDF uses snake_case

### 2. OData Parser (`GenerateListQuery()`)
- **Purpose:** Parses OData expressions and converts to IDF format
- **Location:** `idf/odata_parser.go`
- **Uses:** `ntnx-api-odata-go` library (same as az-manager)

### 3. IDF Query Evaluator
- **Purpose:** Converts parsed OData to IDF query clauses
- **Location:** `idf/odata_parser.go` (uses `idf.IDFQueryEvaluator`)
- **Output:** `WhereClause`, `GroupSortOrder`, `RawColumns`

### 4. Error Handler (`handleODataError()`)
- **Purpose:** Converts parsing errors to user-friendly messages
- **Location:** `grpc/item_grpc_service.go`
- **Returns:** gRPC status errors with proper codes

---

## Dependencies

### Required Package:
```go
github.com/nutanix-core/ntnx-api-odata-go v1.0.27
```

**Sub-packages used:**
- `odata/edm` - EDM provider and bindings
- `odata/uri/parser` - OData parser
- `db/idf` - IDF query evaluator

**Why this package:**
- Same package used by az-manager and guru
- Well-tested and maintained
- Standard for Nutanix APIs

---

## Alignment with az-manager

### Same Pattern:
```go
// az-manager pattern:
edmProvider := edm.NewCustomEdmProvider(entityBindings)
odataParser := parser.NewParser(edmProvider)
uriInfo, err := odataParser.ParserWithQueryParam(queryParam, resourcePath)
idfQueryEval := idf.IDFQueryEvaluator{}
idfQuery, err := idfQueryEval.GetQuery(uriInfo, resourcePath)

// Our implementation (same pattern):
edmProvider := edm.NewCustomEdmProvider(entityBindings)
odataParser := parser.NewParser(edmProvider)
uriInfo, err := odataParser.ParserWithQueryParam(queryParam, resourcePath)
idfQueryEval := idf.IDFQueryEvaluator{}
idfQuery, err := idfQueryEval.GetQuery(uriInfo, resourcePath)
```

### Same Structure:
- `GenerateListQuery()` function (same name as az-manager)
- `constructIDFQuery()` helper function
- EDM bindings with property mappings
- Error handling with context

---

## Code Files Overview

### New Files Created:
1. **`idf/odata_parser.go`**
   - `GenerateListQuery()` - Main OData parser function
   - `constructIDFQuery()` - Builds final IDF query
   - `GetNexusEntityBindings()` - Creates EDM bindings
   - `createItemEntityBinding()` - Item entity binding

2. **`utils/odata/odata_error_handler.go`**
   - Error handling utilities (for future use)
   - Currently, error handling is in gRPC service

### Modified Files:
1. **`idf/idf_item_repository.go`**
   - `ListItems()` now calls `GenerateListQuery()`
   - Uses OData parser instead of manual query building

2. **`grpc/item_grpc_service.go`**
   - Added `handleODataError()` function
   - Integrates error handling for OData parsing

3. **`go.mod`**
   - Added `ntnx-api-odata-go v1.0.27` dependency

---

## How It Works: Detailed Example

### Example 1: Simple Filter

**User Request:**
```
GET /api/nexus/v4.1/config/items?$filter=itemName eq 'test item 0'
```

**Processing:**

1. **Extract:**
   ```go
   queryParams.Filter = "itemName eq 'test item 0'"
   ```

2. **EDM Mapping:**
   ```go
   // Parser sees: "itemName"
   // EDM maps: "itemName" → "item_name"
   ```

3. **Parse:**
   ```go
   // Parser converts: "itemName eq 'test item 0'"
   // To: ComparisonExpression with:
   //   - Lhs: "item_name" (IDF column)
   //   - Operator: EQ
   //   - Rhs: "test item 0"
   ```

4. **IDF Query:**
   ```go
   query.WhereClause = {
     ComparisonExpr: {
       Operator: EQ,
       Lhs: { Column: "item_name" },
       Rhs: { Value: "test item 0" }
     }
   }
   ```

5. **IDF Execution:**
   ```
   SELECT * FROM item WHERE item_name = 'test item 0'
   ```

6. **Result:**
   - Returns only items where `item_name = 'test item 0'`

---

### Example 2: Filter + Sort

**User Request:**
```
GET /api/nexus/v4.1/config/items?$filter=itemType eq 'TYPE1'&$orderby=itemName asc
```

**Processing:**

1. **Extract:**
   ```go
   queryParams.Filter = "itemType eq 'TYPE1'"
   queryParams.Orderby = "itemName asc"
   ```

2. **EDM Mapping:**
   ```go
   "itemType" → "item_type"
   "itemName" → "item_name"
   ```

3. **Parse:**
   ```go
   // Filter: item_type = 'TYPE1'
   // Orderby: item_name ASC
   ```

4. **IDF Query:**
   ```go
   query.WhereClause = {
     ComparisonExpr: {
       Operator: EQ,
       Lhs: { Column: "item_type" },
       Rhs: { Value: "TYPE1" }
     }
   }
   query.GroupBy.RawSortOrder = [{
     Column: "item_name",
     Order: ASC
   }]
   ```

5. **IDF Execution:**
   ```
   SELECT * FROM item 
   WHERE item_type = 'TYPE1' 
   ORDER BY item_name ASC
   ```

6. **Result:**
   - Returns TYPE1 items, sorted by name

---

## Error Handling Flow

### Example: Invalid Field Name

**User Request:**
```
GET /api/nexus/v4.1/config/items?$filter=invalidField eq 'test'
```

**Processing:**

1. **Parse:**
   ```go
   // Parser tries to find "invalidField" in EDM bindings
   // Not found! Returns error
   ```

2. **Error:**
   ```go
   err = "property 'invalidField' not found"
   ```

3. **Error Handler:**
   ```go
   // handleODataError() detects: "property" + "not found"
   return status.Errorf(codes.InvalidArgument,
       "Unknown property in OData query: property 'invalidField' not found")
   ```

4. **Response:**
   ```json
   {
     "error": {
       "message": "Unknown property in OData query: property 'invalidField' not found. Please check field names (itemId, itemName, itemType, extId)."
     }
   }
   ```
   - HTTP Status: 400 Bad Request

---

## Why This Implementation?

### 1. **Standard Library**
- Uses `ntnx-api-odata-go` (same as az-manager)
- Well-tested and maintained
- Handles complex expressions

### 2. **EDM Bindings**
- Maps OData field names to IDF columns
- Defines filterable/sortable properties
- Type-safe (validates types)

### 3. **Error Handling**
- User-friendly error messages
- Proper HTTP status codes
- Clear guidance on what went wrong

### 4. **Aligned with az-manager**
- Same pattern and structure
- Same packages and functions
- Consistent across Nutanix APIs

---

## Summary for Team Presentation

### What We Built:
1. **OData Parser** - Converts user queries to IDF format
2. **EDM Bindings** - Maps OData fields to IDF columns
3. **Error Handler** - User-friendly error messages
4. **Integration** - Works with existing IDF repository

### How It Works:
1. User sends OData query (`$filter=itemName eq 'test'`)
2. Parser extracts and validates query
3. EDM maps field names (itemName → item_name)
4. IDF query evaluator converts to IDF format
5. IDF executes filtered/sorted query
6. Results mapped back to protobuf Items

### Why It Matters:
- **Before:** Users got all data, no filtering
- **After:** Users can filter, sort, and select fields
- **Standard:** Same as other Nutanix APIs (az-manager, guru)

### Key Files:
- `idf/odata_parser.go` - Main parser implementation
- `idf/idf_item_repository.go` - Uses parser
- `grpc/item_grpc_service.go` - Error handling

### Dependencies:
- `ntnx-api-odata-go v1.0.27` - OData parsing library

---

**This implementation follows the exact same pattern as az-manager, ensuring consistency across Nutanix APIs.**

