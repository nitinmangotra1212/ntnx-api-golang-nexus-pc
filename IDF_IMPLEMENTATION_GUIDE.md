# IDF Implementation Guide for golang-mock

## What is IDF?

**IDF (Insights Data Format)** is Nutanix's internal database/storage layer used for:
- **Persistent storage** of entities (domains, BOMs, etc.)
- **Querying with OData filters** (`$filter`, `$orderby`, `$select`, `$expand`)
- **Pagination** support
- **Complex queries** with grouping, sorting, and metrics

## How IDF Works in az-manager

### Architecture Pattern

```
Service Layer (domain_Services.go)
    ↓
Repository Interface (db.DomainRepository)
    ↓
IDF Implementation (idf.DomainRepositoryImpl)
    ↓
IDF Client (external/idf/idf_client.go)
    ↓
Insights Service (go-cache/insights/insights_interface)
```

### Key Components

1. **Repository Interface** (`db/domain_repository.go`)
   ```go
   type DomainRepository interface {
       CreateDomain(*models.DomainEntity) error
       GetDomainById(extId string) (*models.DomainEntity, error)
       ListDomains(queryParams *models.QueryParams) ([]*lifecycleConfig.Domain, int64, error)
   }
   ```

2. **IDF Repository Implementation** (`idf/idf_domain_repository.go`)
   - Implements the repository interface
   - Converts protobuf models ↔ IDF attributes
   - Handles Create, Read, List operations

3. **IDF Client** (`external/idf/idf_client.go`)
   - Wraps `InsightsService` from `go-cache/insights/insights_interface`
   - Provides: `GetEntityRet`, `UpdateEntityRet`, `GetEntitiesWithMetricsRet`

4. **IDF Utils** (`idf/idf_utils.go`)
   - `AddAttribute()` - Converts Go types to IDF `AttributeDataArg`
   - `CreateDataArg()` - Type conversion (string, int, bool, arrays, etc.)
   - `constructIDFQuery()` - Builds IDF queries from OData params

## Do You Need IDF for golang-mock?

### ❌ **You DON'T need IDF if:**
- ✅ You're fine with **in-memory storage** (current implementation)
- ✅ You don't need **persistent storage** across restarts
- ✅ You don't need **complex OData filtering** (simple filtering is fine)
- ✅ This is a **pure mock service** for testing

### ✅ **You DO need IDF if:**
- You want **persistent storage** (data survives restarts)
- You need **OData query support** (`$filter`, `$orderby`, etc.)
- You want to **integrate with Nutanix IDF infrastructure**
- You need **complex queries** with grouping/sorting

## What You Need to Do (If Implementing IDF)

### Step 1: Add Dependencies

**File**: `go.mod`

```go
require (
    // ... existing dependencies ...
    github.com/nutanix-core/go-cache v0.0.0-20240613003120-de1e4c3ed003
    github.com/nutanix-core/ntnx-api-odata-go v1.0.27
)
```

### Step 2: Create Repository Interface

**File**: `golang-mock-service/db/cat_repository.go` (NEW)

```go
package db

import (
    pb "github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/config"
    "github.com/nutanix/ntnx-api-golang-mock/golang-mock-service/models"
)

type CatRepository interface {
    CreateCat(catEntity *models.CatEntity) error
    GetCatById(extId string) (*models.CatEntity, error)
    ListCats(queryParams *models.QueryParams) ([]*pb.Cat, int64, error)
    UpdateCat(extId string, catEntity *models.CatEntity) error
    DeleteCat(extId string) error
}
```

### Step 3: Create IDF Client

**File**: `golang-mock-service/external/idf/interface.go` (NEW)

```go
package idf

import (
    "github.com/nutanix-core/go-cache/insights/insights_interface"
)

type IdfClientIfc interface {
    GetEntityRet(getArg *insights_interface.GetEntitiesArg) (*insights_interface.GetEntitiesRet, error)
    UpdateEntityRet(updateArg *insights_interface.UpdateEntityArg) (*insights_interface.UpdateEntityRet, error)
    GetEntitiesWithMetricsRet(getEntitiesWithMetricsArg *insights_interface.GetEntitiesWithMetricsArg) (*insights_interface.GetEntitiesWithMetricsRet, error)
    GetInsightsService() insights_interface.InsightsServiceInterface
}

type IdfClientImpl struct {
    IdfSvc insights_interface.InsightsServiceInterface
}

func NewIdfClient(host string, port uint16) IdfClientImpl {
    IdfService := insights_interface.NewInsightsService(host, port)
    return IdfClientImpl{
        IdfSvc: IdfService,
    }
}
```

**File**: `golang-mock-service/external/idf/idf_client.go` (NEW)

```go
package idf

import (
    "github.com/nutanix-core/go-cache/insights/insights_interface"
    log "github.com/sirupsen/logrus"
)

func (idf IdfClientImpl) GetEntityRet(getArg *insights_interface.GetEntitiesArg) (*insights_interface.GetEntitiesRet, error) {
    if getArg == nil {
        log.Error("Nil get argument while trying to read from IDF")
        return nil, fmt.Errorf("invalid argument")
    }
    getResponse := &insights_interface.GetEntitiesRet{}
    err := idf.IdfSvc.SendMsgWithTimeout("GetOperationIdf", getArg, getResponse, nil, 30)
    return getResponse, err
}

func (idf IdfClientImpl) UpdateEntityRet(updateArg *insights_interface.UpdateEntityArg) (*insights_interface.UpdateEntityRet, error) {
    if updateArg == nil {
        log.Error("Invalid update argument")
        return nil, fmt.Errorf("invalid argument")
    }
    updateResponse := &insights_interface.UpdateEntityRet{}
    err := idf.IdfSvc.SendMsgWithTimeout("UpdateOperationIdf", updateArg, updateResponse, nil, 30)
    return updateResponse, err
}

func (idf IdfClientImpl) GetEntitiesWithMetricsRet(getEntitiesWithMetricsArg *insights_interface.GetEntitiesWithMetricsArg) (*insights_interface.GetEntitiesWithMetricsRet, error) {
    if getEntitiesWithMetricsArg == nil {
        log.Error("Invalid getEntitiesWithMetrics argument")
        return nil, fmt.Errorf("invalid argument")
    }
    getResponse := &insights_interface.GetEntitiesWithMetricsRet{}
    err := idf.IdfSvc.SendMsgWithTimeout("GetEntitiesWithMetricsOperationIdf", getEntitiesWithMetricsArg, getResponse, nil, 30)
    return getResponse, err
}

func (idf IdfClientImpl) GetInsightsService() insights_interface.InsightsServiceInterface {
    return idf.IdfSvc
}
```

### Step 4: Create IDF Utils

**File**: `golang-mock-service/idf/idf_utils.go` (NEW)

```go
package idf

import (
    idfIfc "github.com/nutanix-core/go-cache/insights/insights_interface"
    log "github.com/sirupsen/logrus"
    "google.golang.org/protobuf/proto"
)

// AddAttribute adds an attribute to the attribute data arg list
func AddAttribute(attributeDataArgList *[]*idfIfc.AttributeDataArg, name string, value interface{}) {
    dataArg := CreateDataArg(name, value)
    if dataArg == nil {
        log.Errorf("failed to create data arg for attribute %s", name)
        return
    }
    *attributeDataArgList = append(*attributeDataArgList, dataArg)
}

// CreateDataArg creates a data arg for the given name and value
func CreateDataArg(name string, value interface{}) *idfIfc.AttributeDataArg {
    dataValue := &idfIfc.DataValue{}

    switch val := value.(type) {
    case string:
        dataValue.ValueType = &idfIfc.DataValue_StrValue{StrValue: val}
    case int32:
        dataValue.ValueType = &idfIfc.DataValue_Int64Value{Int64Value: int64(val)}
    case int64:
        dataValue.ValueType = &idfIfc.DataValue_Int64Value{Int64Value: val}
    case bool:
        dataValue.ValueType = &idfIfc.DataValue_BoolValue{BoolValue: val}
    case []string:
        dataValue.ValueType = &idfIfc.DataValue_StrList_{
            StrList: &idfIfc.DataValue_StrList{ValueList: val},
        }
    default:
        log.Errorf("Unsupported type for attribute %s: %T", name, value)
        return nil
    }

    return &idfIfc.AttributeDataArg{
        AttributeData: &idfIfc.AttributeData{
            Name:  proto.String(name),
            Value: dataValue,
        },
    }
}
```

### Step 5: Create IDF Repository Implementation

**File**: `golang-mock-service/idf/idf_cat_repository.go` (NEW)

```go
package idf

import (
    "github.com/google/uuid"
    "github.com/nutanix-core/go-cache/insights/insights_interface"
    idfQr "github.com/nutanix-core/go-cache/insights/insights_interface/query"
    pb "github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/config"
    "github.com/nutanix/ntnx-api-golang-mock/golang-mock-service/db"
    "github.com/nutanix/ntnx-api-golang-mock/golang-mock-service/external/idf"
    "github.com/nutanix/ntnx-api-golang-mock/golang-mock-service/models"
    log "github.com/sirupsen/logrus"
    "google.golang.org/protobuf/proto"
)

type CatRepositoryImpl struct {
    idfClient idf.IdfClientIfc
}

const (
    catEntityTypeName = "mock_cat"
    catIdAttr         = "cat_id"
    catNameAttr       = "cat_name"
    catTypeAttr       = "cat_type"
    descriptionAttr   = "description"
)

func NewCatRepository(idfClient idf.IdfClientIfc) db.CatRepository {
    return &CatRepositoryImpl{
        idfClient: idfClient,
    }
}

func (r *CatRepositoryImpl) CreateCat(catEntity *models.CatEntity) error {
    catUuid := uuid.New().String()
    attributeDataArgList := []*insights_interface.AttributeDataArg{}

    // Add cat attributes
    if catEntity.Cat.CatId != nil {
        AddAttribute(&attributeDataArgList, catIdAttr, *catEntity.Cat.CatId)
    }
    if catEntity.Cat.CatName != nil {
        AddAttribute(&attributeDataArgList, catNameAttr, *catEntity.Cat.CatName)
    }
    if catEntity.Cat.CatType != nil {
        AddAttribute(&attributeDataArgList, catTypeAttr, *catEntity.Cat.CatType)
    }
    if catEntity.Cat.Description != nil {
        AddAttribute(&attributeDataArgList, descriptionAttr, *catEntity.Cat.Description)
    }

    updateArg := &insights_interface.UpdateEntityArg{
        EntityGuid: &insights_interface.EntityGuid{
            EntityTypeName: proto.String(catEntityTypeName),
            EntityId:       &catUuid,
        },
        AttributeDataArgList: attributeDataArgList,
    }

    _, err := r.idfClient.UpdateEntityRet(updateArg)
    if err != nil {
        log.Errorf("Failed to create cat: %v", err)
        return err
    }

    // Set extId
    if catEntity.Cat.Base == nil {
        catEntity.Cat.Base = &response.ExternalizableAbstractModel{}
    }
    catEntity.Cat.Base.ExtId = &catUuid

    return nil
}

func (r *CatRepositoryImpl) ListCats(queryParams *models.QueryParams) ([]*pb.Cat, int64, error) {
    // Build IDF query
    query, err := idfQr.QUERY(catEntityTypeName + "ListQuery").
        FROM(catEntityTypeName).Proto()
    if err != nil {
        return nil, 0, err
    }

    // Add pagination
    page := queryParams.Page
    limit := queryParams.Limit
    if limit <= 0 {
        limit = 50
    }
    offset := page * limit

    if query.GroupBy == nil {
        query.GroupBy = &insights_interface.QueryGroupBy{}
    }
    query.GroupBy.RawLimit = &insights_interface.QueryLimit{
        Limit:  proto.Int64(int64(limit)),
        Offset: proto.Int64(int64(offset)),
    }

    // Add OData filter if provided
    if queryParams.Filter != "" {
        // Parse OData filter and convert to IDF WhereClause
        // (This requires ntnx-api-odata-go)
    }

    queryArg := &insights_interface.GetEntitiesWithMetricsArg{
        Query: query,
    }

    // Query IDF
    queryResponse, err := r.idfClient.GetEntitiesWithMetricsRet(queryArg)
    if err != nil {
        return nil, 0, err
    }

    // Convert IDF entities to Cat protobufs
    var cats []*pb.Cat
    groupResults := queryResponse.GetGroupResultsList()
    if len(groupResults) == 0 {
        return []*pb.Cat{}, 0, nil
    }

    entities := ConvertEntitiesWithMetricToEntities(groupResults[0].GetRawResults())
    for _, entity := range entities {
        cat := r.mapIdfAttributeToCat(entity)
        cats = append(cats, cat)
    }

    totalCount := groupResults[0].GetTotalEntityCount()
    return cats, totalCount, nil
}

func (r *CatRepositoryImpl) mapIdfAttributeToCat(entity *insights_interface.Entity) *pb.Cat {
    cat := &pb.Cat{}
    for _, attr := range entity.GetAttributeDataMap() {
        switch attr.GetName() {
        case catIdAttr:
            val := int32(attr.GetValue().GetInt64Value())
            cat.CatId = &val
        case catNameAttr:
            val := attr.GetValue().GetStrValue()
            cat.CatName = &val
        case catTypeAttr:
            val := attr.GetValue().GetStrValue()
            cat.CatType = &val
        case descriptionAttr:
            val := attr.GetValue().GetStrValue()
            cat.Description = &val
        }
    }
    return cat
}
```

### Step 6: Update Service to Use Repository

**File**: `golang-mock-service/grpc/cat_grpc_service.go`

```go
// Replace in-memory map with repository
type CatGrpcService struct {
    pb.UnimplementedCatServiceServer
    catRepository db.CatRepository  // ← Add this
}

func NewCatGrpcService(catRepository db.CatRepository) *CatGrpcService {
    return &CatGrpcService{
        catRepository: catRepository,
    }
}

func (s *CatGrpcService) ListCats(ctx context.Context, req *pb.ListCatsArg) (*pb.ListCatsRet, error) {
    // Convert request to QueryParams
    queryParams := &models.QueryParams{
        Page:   req.GetXPage(),
        Limit:  req.GetXLimit(),
        Filter: req.GetXFilter(),
        // ... other params
    }

    // Use repository instead of in-memory map
    cats, totalCount, err := s.catRepository.ListCats(queryParams)
    if err != nil {
        return nil, err
    }

    // Create response with metadata
    isPaginated := req.GetXPage() > 0 || req.GetXLimit() > 0
    apiUrl := responseUtils.GetApiUrl(ctx, queryParams.Filter, "", "", &queryParams.Limit, &queryParams.Page)
    paginationLinks := responseUtils.GetPaginationLinks(totalCount, apiUrl)
    
    return responseUtils.CreateListCatsResponse(cats, paginationLinks, isPaginated, int32(totalCount)), nil
}
```

### Step 7: Initialize IDF Client in Main

**File**: `golang-mock-service/server/main.go`

```go
import (
    "github.com/nutanix/ntnx-api-golang-mock/golang-mock-service/external/idf"
    "github.com/nutanix/ntnx-api-golang-mock/golang-mock-service/idf"
    "github.com/nutanix/ntnx-api-golang-mock/golang-mock-service/db"
)

func main() {
    // ... existing code ...

    // Initialize IDF client (if using IDF)
    idfClient := idf.NewIdfClient("localhost", 9876)  // Adjust host/port as needed
    catRepository := idf.NewCatRepository(idfClient)

    // Pass repository to service
    catService := grpc.NewCatGrpcService(catRepository)
    
    // ... rest of code ...
}
```

## Summary: What You Need to Do

### If You DON'T Need IDF (Recommended for Mock):
- ✅ **Nothing!** Your current in-memory implementation is fine
- ✅ Keep using the `map[int32]*pb.Cat` approach
- ✅ Simple and works for mocking/testing

### If You DO Need IDF:
1. ✅ Add dependencies (`go-cache`, `ntnx-api-odata-go`)
2. ✅ Create repository interface (`db/cat_repository.go`)
3. ✅ Create IDF client (`external/idf/`)
4. ✅ Create IDF utils (`idf/idf_utils.go`)
5. ✅ Create IDF repository implementation (`idf/idf_cat_repository.go`)
6. ✅ Update service to use repository instead of map
7. ✅ Initialize IDF client in `main.go`

## Recommendation

**For a mock service, you probably DON'T need IDF** unless:
- You need persistent storage
- You need OData query support
- You're integrating with Nutanix infrastructure

**Your current in-memory approach is simpler and sufficient for mocking!**

---

**Last Updated**: 2025-01-26

