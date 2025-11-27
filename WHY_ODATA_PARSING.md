# What is OData Parsing and Why Do We Use It?

## What is OData?

**OData (Open Data Protocol)** is a standard protocol for building and consuming RESTful APIs. It provides a uniform way to query and manipulate data.

Think of it as a **query language for REST APIs** - similar to SQL for databases, but for HTTP APIs.

---

## Real-World Example

### Without OData (Basic API):
```
GET /api/nexus/v4.1/config/items?page=0&limit=10
```
**Problem:** You can only get all items, no filtering or sorting!

### With OData (Powerful API):
```
GET /api/nexus/v4.1/config/items?$filter=itemName eq 'test'&$orderby=itemId asc&$page=0&$limit=10
```
**Benefit:** You can filter, sort, select specific fields, and paginate - all in one request!

---

## OData Query Parameters

### 1. **$filter** - Filtering (WHERE clause in SQL)
```
$filter=itemName eq 'test item 0'
```
**Meaning:** "Give me only items where itemName equals 'test item 0'"

**More examples:**
- `$filter=itemId gt 5` → "itemId greater than 5"
- `$filter=itemType eq 'TYPE1' and itemId lt 10` → "itemType is TYPE1 AND itemId less than 10"
- `$filter=itemName startswith 'test'` → "itemName starts with 'test'"

### 2. **$orderby** - Sorting (ORDER BY in SQL)
```
$orderby=itemId asc
```
**Meaning:** "Sort results by itemId in ascending order"

**More examples:**
- `$orderby=itemName desc` → "Sort by itemName descending"
- `$orderby=itemId asc,itemName desc` → "Sort by itemId ascending, then itemName descending"

### 3. **$select** - Field Selection (SELECT in SQL)
```
$select=itemId,itemName
```
**Meaning:** "Only return itemId and itemName fields (not itemType, description, etc.)"

### 4. **$expand** - Related Data (JOIN in SQL)
```
$expand=relatedItems
```
**Meaning:** "Include related items in the response"

### 5. **$page and $limit** - Pagination
```
$page=0&$limit=10
```
**Meaning:** "Get page 0 (first page) with 10 items per page"

---

## Why Do We Need OData Parsing?

### Problem: OData is Human-Readable, IDF Needs Machine-Readable

**User sends (OData):**
```
$filter=itemName eq 'test'
```

**IDF needs (internal format):**
```go
WhereClause: {
  ComparisonExpr: {
    Operator: EQ,
    Lhs: { Column: "item_name" },  // Note: snake_case!
    Rhs: { Value: "test" }
  }
}
```

**The Parser Does:**
1. **Parses** the OData expression (`itemName eq 'test'`)
2. **Maps** field names (`itemName` → `item_name`)
3. **Converts** to IDF query format (WhereClause)
4. **Validates** syntax and operators

---

## Without OData Parser (What We Had Before)

### Before:
```go
// User sends: $filter=itemName eq 'test'
// We just stored it as a string:
queryParams.Filter = "itemName eq 'test'"

// IDF query had NO filtering:
query.WhereClause = nil  // ❌ No filter applied!
```

**Result:** Filter parameter was ignored! Users couldn't filter data.

---

## With OData Parser (What We Have Now)

### After:
```go
// User sends: $filter=itemName eq 'test'
// Parser converts it:
query.WhereClause = {
  ComparisonExpr: {
    Operator: EQ,
    Lhs: { Column: "item_name" },  // ✅ Mapped from itemName
    Rhs: { Value: "test" }
  }
}
```

**Result:** Filter works! Users can filter, sort, and select fields.

---

## The Parsing Process (Step by Step)

### Step 1: User Request
```
GET /api/nexus/v4.1/config/items?$filter=itemName eq 'test'&$orderby=itemId asc
```

### Step 2: Extract Query Parameters
```go
queryParams := ExtractQueryParamsFromContext(ctx)
// queryParams.Filter = "itemName eq 'test'"
// queryParams.Orderby = "itemId asc"
```

### Step 3: OData Parser
```go
// Create EDM provider (maps itemName → item_name)
edmProvider := edm.NewCustomEdmProvider(entityBindings)

// Create parser
odataParser := parser.NewParser(edmProvider)

// Parse the filter expression
uriInfo, err := odataParser.ParserWithQueryParam(queryParam, "/items")
// uriInfo now contains parsed expression with field names mapped
```

### Step 4: IDF Query Evaluator
```go
// Convert parsed OData to IDF query format
idfQueryEval := idf.IDFQueryEvaluator{}
idfQuery, err := idfQueryEval.GetQuery(uriInfo, "/items")
// idfQuery.WhereClause now contains the filter in IDF format
// idfQuery.GroupBy.RawSortOrder now contains the sort order
```

### Step 5: Execute IDF Query
```go
// IDF executes the query with filter and sort
queryArg := &GetEntitiesWithMetricsArg{
  Query: idfQuery,  // Contains WhereClause and GroupSortOrder
}
response, err := idfClient.GetEntitiesWithMetricsRet(queryArg)
```

### Step 6: Return Results
```json
{
  "data": [
    {
      "itemId": 1,
      "itemName": "test",  // ✅ Filtered!
      ...
    }
  ]
}
```

---

## Why Not Just Parse Manually?

### Manual Parsing (Complex and Error-Prone):
```go
// ❌ Manual parsing - lots of code, error-prone
if strings.Contains(filter, "eq") {
  parts := strings.Split(filter, " eq ")
  fieldName := parts[0]  // "itemName"
  value := parts[1]      // "'test'"
  
  // Map field name
  idfColumn := mapFieldName(fieldName)  // "item_name"
  
  // Handle different operators (eq, ne, gt, lt, ge, le, and, or, not...)
  // Handle nested expressions
  // Handle type conversions
  // Handle edge cases
  // ... hundreds of lines of code ...
}
```

### OData Parser (Standard and Reliable):
```go
// ✅ OData parser - handles everything
odataParser := parser.NewParser(edmProvider)
uriInfo, err := odataParser.ParserWithQueryParam(queryParam, resourcePath)
// Done! Handles all operators, nested expressions, type conversions, etc.
```

**Benefits:**
- ✅ Standard protocol (works the same way across all Nutanix APIs)
- ✅ Handles complex expressions (nested AND/OR, functions, etc.)
- ✅ Type-safe (validates types, converts values)
- ✅ Well-tested (used by az-manager, guru, and other services)
- ✅ Less code (we don't write parsing logic ourselves)

---

## Field Name Mapping (Why EDM Bindings?)

### Problem: Different Naming Conventions

**OData (API layer):** Uses camelCase
- `itemName`, `itemId`, `itemType`

**IDF (Database layer):** Uses snake_case
- `item_name`, `item_id`, `item_type`

### Solution: EDM (Entity Data Model) Bindings

**EDM Bindings** map between the two:
```go
PropertyMappings = {
  "itemName": "item_name",  // OData → IDF
  "itemId": "item_id",
  "itemType": "item_type",
  ...
}
```

**How it works:**
1. User sends: `$filter=itemName eq 'test'`
2. Parser sees: `itemName` (OData field)
3. EDM maps: `itemName` → `item_name` (IDF column)
4. IDF query uses: `item_name` ✅

---

## Real-World Use Cases

### Use Case 1: Search by Name
```
GET /api/nexus/v4.1/config/items?$filter=itemName eq 'test item 0'
```
**Without parser:** Returns all items (filter ignored)  
**With parser:** Returns only items with name 'test item 0' ✅

### Use Case 2: Sort by ID
```
GET /api/nexus/v4.1/config/items?$orderby=itemId asc
```
**Without parser:** Returns items in random order  
**With parser:** Returns items sorted by itemId ascending ✅

### Use Case 3: Filter and Sort
```
GET /api/nexus/v4.1/config/items?$filter=itemType eq 'TYPE1'&$orderby=itemName asc
```
**Without parser:** Returns all items, no filtering or sorting  
**With parser:** Returns TYPE1 items, sorted by name ✅

### Use Case 4: Select Specific Fields
```
GET /api/nexus/v4.1/config/items?$select=itemId,itemName
```
**Without parser:** Returns all fields  
**With parser:** Returns only itemId and itemName (smaller response) ✅

---

## Comparison: Before vs After

### Before (Without OData Parser):
```go
// User: $filter=itemName eq 'test'
queryParams.Filter = "itemName eq 'test'"  // Just a string

// IDF query:
query.WhereClause = nil  // ❌ Filter not applied!

// Result: Returns ALL items (filter ignored)
```

### After (With OData Parser):
```go
// User: $filter=itemName eq 'test'
queryParams.Filter = "itemName eq 'test'"

// OData parser converts:
query.WhereClause = {
  ComparisonExpr: {
    Operator: EQ,
    Lhs: { Column: "item_name" },  // ✅ Mapped!
    Rhs: { Value: "test" }
  }
}

// Result: Returns only items where item_name = 'test' ✅
```

---

## Why az-manager Uses It

**az-manager** (and other Nutanix services) use OData parsing because:

1. **Standard Protocol:** All Nutanix APIs support OData queries
2. **User Expectations:** Users expect to filter/sort via `$filter` and `$orderby`
3. **Consistency:** Same query syntax across all APIs
4. **Powerful:** Supports complex queries (nested filters, functions, etc.)

**Example from az-manager:**
```
GET /api/az-manager/v4.1/config/domains?$filter=name eq 'test'&$orderby=createdTime desc
```
- Uses OData parser to convert to IDF query
- Returns filtered and sorted results

**We follow the same pattern** to maintain consistency.

---

## Summary

### What is OData Parsing?
- Converting human-readable OData query strings (`$filter=itemName eq 'test'`) into machine-readable IDF query format
- Mapping OData field names (camelCase) to IDF column names (snake_case)
- Validating syntax and operators

### Why Do We Use It?
1. **User Experience:** Users can filter, sort, and select fields
2. **Standard Protocol:** Same as other Nutanix APIs (az-manager, guru, etc.)
3. **Less Code:** Don't write parsing logic ourselves
4. **Well-Tested:** Uses proven library (`ntnx-api-odata-go`)
5. **Powerful:** Supports complex queries (nested expressions, functions)

### Without It:
- ❌ Users can't filter data
- ❌ Users can't sort data
- ❌ Users can't select specific fields
- ❌ API is less useful

### With It:
- ✅ Users can filter: `$filter=itemName eq 'test'`
- ✅ Users can sort: `$orderby=itemId asc`
- ✅ Users can select: `$select=itemId,itemName`
- ✅ API is powerful and user-friendly

---

**In Simple Terms:**
OData parsing is like a **translator** that converts user-friendly query strings into the format that IDF (our database) understands. Without it, users can't filter or sort - they just get all data. With it, users have full control over what data they get!

