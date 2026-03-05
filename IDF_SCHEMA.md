# IDF Schema Reference

This document describes the entity types and their schemas registered in IDF (Insights Data Framework) by the `register_create_item_idf.py` script.

**Script:** `ntnx-api-golang-nexus-pc/register_create_item_idf.py`

**IDF Connection:** `127.0.0.1:2027` (localhost on PC)

---

## Entity Relationship Diagram

```
┌──────────────────────┐         ┌──────────────────────────┐
│        item           │         │       item_stats          │
│  (config entity)      │ 1 ── 1 │  (stats entity / TSDB)    │
├──────────────────────┤         ├──────────────────────────┤
│  entity_id (UUID)     │◄────────│  item_ext_id (FK)         │
│  ext_id (= entity_id) │         │  entity_id (UUID)         │
│  item_id              │         │  age          [timeseries] │
│  item_name            │         │  heart_rate   [timeseries] │
│  item_type            │         │  food_intake  [timeseries] │
│  description          │         └──────────────────────────┘
│  quantity             │
│  price                │
│  is_active            │
│  priority             │
│  status               │
│  int64_list           │
└──────────────────────┘

JOIN key:  item.ext_id = item_stats.item_ext_id
```

---

## Entity Type: `item`

Config entity representing an item. 120 entities are created by the script (100 TYPE1 + 20 TYPE2).

### Entity Type Configuration

| Property | Value |
|---|---|
| `entity_type_name` | `item` |
| `suppress_replication` | `true` |
| `is_backup_required` | `false` |
| `enable_replication_from` | `kNDFS` |
| `is_evictable` | `false` |
| `retain_attributes_on_deletion` | `false` |
| `track_attribute_changes` | `false` |
| `enable_pulsehd_collection` | `false` |

### Attributes (Metrics)

All fields are registered with `is_attribute: true` (queryable, filterable, stored as attributes).

| Attribute | IDF Name | Data Type | Description | Example Values |
|---|---|---|---|---|
| Item ID | `item_id` | `int64` | Unique numeric identifier | 1 – 220 |
| Item Name | `item_name` | `string` | Display name | `"test item 0"`, `"TYPE2 Item 1"` |
| Item Type | `item_type` | `int64` | Enum: 1 = TYPE1, 2 = TYPE2 | `1`, `2` |
| Description | `description` | `string` | Item description | `"test item description 0"` |
| External ID | `ext_id` | `string` | UUID (matches `entity_id`) | `"550e8400-..."` |
| Quantity | `quantity` | `int64` | Stock count | 10 – 200 (random) |
| Price | `price` | `double` | Item price | 10.0 – 50.0 (random) |
| Is Active | `is_active` | `bool` | Active flag | `true` / `false` |
| Priority | `priority` | `int64` | Priority level | 1 – 5 (random) |
| Status | `status` | `string` | Current status | `"pending"`, `"active"`, `"inactive"` |
| Int64 List | `int64_list` | `int64_list` | List of integers | `[150, 25, 80]` (3 random values) |

### Data Distribution

| Subset | Count | `item_type` | `item_id` Range | Naming Pattern |
|---|---|---|---|---|
| TYPE1 | 100 | `1` | 1 – 100 | `"test item {i}"` |
| TYPE2 | 20 | `2` | 201 – 220 | `"TYPE2 Item {i}"` |

**Entity ID Generation:** Deterministic UUIDs using `uuid5(NAMESPACE_DNS, "nexus-item-{i}")` so re-runs overwrite existing data.

---

## Entity Type: `item_stats`

Stats entity representing time-series metrics for items. 120 entities are created (one-to-one mapping with items).

### Entity Type Configuration

| Property | Value |
|---|---|
| `entity_type_name` | `item_stats` |
| `suppress_replication` | `true` |
| `is_backup_required` | `false` |
| `enable_replication_from` | `kNDFS` |
| `is_evictable` | `false` |
| `retain_attributes_on_deletion` | `false` |
| `track_attribute_changes` | `false` |
| `enable_pulsehd_collection` | `false` |

### Attributes and Metrics

| Attribute | IDF Name | `is_attribute` | Data Type | Storage | Description |
|---|---|---|---|---|---|
| Item External ID | `item_ext_id` | `true` | `string` | Attribute | Foreign key linking to `item.ext_id` |
| Age | `age` | `false` | `int64` | TSDB (time-series) | Age metric, constant per entity (1–10) |
| Heart Rate | `heart_rate` | `false` | `int64` | TSDB (time-series) | Heart rate (60–120 bpm, random per data point) |
| Food Intake | `food_intake` | `false` | `double` | TSDB (time-series) | Food intake (100–500, random per data point) |

### Attribute vs. Time-Series

| `is_attribute` | Storage | Queryable via `$filter` | Supports `$expand` time-series | Use Case |
|---|---|---|---|---|
| `true` | IDF attribute store | Yes (direct `eq`, `ne`, etc.) | No | Static fields, JOIN keys |
| `false` | TSDB (time-series DB) | Limited (lambda `any()` with `eq`/`in` only) | Yes (`$startTime`, `$endTime`, `$statType`, `$samplingInterval`) | Metrics collected over time |

### Time-Series Data Points

| Property | Value |
|---|---|
| Data points per metric | 100 |
| Time range | Last 30 days from script execution |
| Interval | ~7.2 hours between data points |
| Push method | `PutMetricData` API (batches of 10 entities) |

### Data Creation Flow

```
1. RegisterEntityTypes("item_stats")     → Register entity type
2. RegisterMetricTypes(...)              → Register attribute + TSDB metrics
3. UpdateEntity(...)                     → Create entity with attribute (item_ext_id)
4. PutMetricData(...)                    → Push time-series data points to TSDB
```

**Entity ID Generation:** Deterministic UUIDs using `uuid5(NAMESPACE_DNS, "nexus-item-stats-{i}")`.

---

## JOIN Relationship

The `item` and `item_stats` entities are linked via a foreign key for `$expand` queries:

| Property | Value |
|---|---|
| Left entity | `item` |
| Right entity | `item_stats` |
| Left key | `ext_id` |
| Right key | `item_ext_id` |
| Cardinality | One-to-one (each item has exactly one item_stats) |
| GraphQL join type | `left_outer_join` |

**OData Query Example:**
```
GET /config/items?$expand=itemStats($startTime=...;$endTime=...;$statType=AVG;$samplingInterval=25200)
```

**Generated GraphQL:**
```graphql
query {
  item(args: {page_size:10, page_offset:0}) {
    item_id, item_name, item_type, description, ext_id,
    quantity, price, is_active, priority, status, int64_list,
    item_stats(args: {
      left_column: ext_id,
      right_column: item_ext_id,
      condition_operator: equal,
      join_type: left_outer_join,
      interval_start_ms: ...,
      interval_end_ms: ...,
      downsampling_interval_secs: 25200
    }) {
      age(sampling:AVG, timeseries:true),
      heart_rate(sampling:AVG, timeseries:true),
      food_intake(sampling:AVG, timeseries:true),
      _entity_id_
    }
  }
  filtered_entity_count, total_entity_count
}
```

---

## OData to IDF Property Mapping

### `item` Entity

| OData Property (camelCase) | IDF Column (snake_case) | Proto Field |
|---|---|---|
| `itemId` | `item_id` | `item_id` |
| `itemName` | `item_name` | `item_name` |
| `itemType` | `item_type` | `item_type` (enum: TYPE1=2001, TYPE2=2002) |
| `description` | `description` | `description` |
| `extId` | `ext_id` | `ext_id` |
| `quantity` | `quantity` | `quantity` |
| `price` | `price` | `price` |
| `isActive` | `is_active` | `is_active` |
| `priority` | `priority` | `priority` |
| `status` | `status` | `status` |
| `int64List` | `int64_list` | `int64_list` |

### `item_stats` Entity

| OData Property (camelCase) | IDF Column (snake_case) | Proto Field | Stats Flag |
|---|---|---|---|
| `itemExtId` | `item_ext_id` | `item_ext_id` | — |
| `age` | `age` | `age` (IntegerTimeValuePair[]) | `stats: true` |
| `heartRate` | `heart_rate` | `heart_rate` (IntegerTimeValuePair[]) | `stats: true` |
| `foodIntake` | `food_intake` | `food_intake` (DoubleTimeValuePair[]) | `stats: true` |

---

## Running the Script

### First-time setup (register + populate)

```bash
# Copy to PC
scp -O ntnx-api-golang-nexus-pc/register_create_item_idf.py \
  nutanix@${PC_IP}:~/ntnx-api-golang-nexus-pc/

# Run on PC
ssh nutanix@${PC_IP}
cd ~/ntnx-api-golang-nexus-pc
python3 register_create_item_idf.py
```

### Expected Output

```
============================================================
Starting IDF registration for item + item_stats...
============================================================

--- Registering and populating 'item' table ---
Registered 'item' entity type
Registered 'item' metric types (attributes)

Creating 120 items in IDF (100 TYPE1 + 20 TYPE2)...
  Created 20/100 TYPE1 items...
  Created 40/100 TYPE1 items...
  ...
  Creating TYPE2 items...
  Created 10/20 TYPE2 items...
  Created 20/20 TYPE2 items...
Successfully created 120 items (100 TYPE1 + 20 TYPE2)

--- Registering and populating 'item_stats' table ---
Registered 'item_stats' entity type
Registered 'item_stats' metric types
  Created 20/120 item_stats entities...
  ...
Pushing time-series metric data to TSDB...
  Pushed metrics for entities 1-10/120
  ...
Successfully created 120 item_stats entities with TSDB time-series data

============================================================
All registrations complete!
  - 120 items created (100 TYPE1 + 20 TYPE2)
  - 120 item_stats created (1-to-1 mapping with items)
  - item_stats linked via 'item_ext_id' -> item 'ext_id'
============================================================
```

> **Note:** The script uses deterministic UUIDs, so re-running it safely overwrites existing data without creating duplicates.
