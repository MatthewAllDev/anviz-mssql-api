# API Reference

Base URL:

```text
http://localhost:8080/api/v1
```

All endpoints require:

```http
Authorization: Bearer <api-key>
```

## Common Behavior

### Scopes

| Scope | Permissions |
| --- | --- |
| `read` | `GET` endpoints. |
| `crud` | All `GET`, `POST`, `PUT`, and `DELETE` endpoints. |

### Pagination

List endpoints accept:

| Query parameter | Default | Maximum | Description |
| --- | --- | --- | --- |
| `limit` | `100` | `1000` | Number of rows to return. |
| `offset` | `0` | - | Number of rows to skip. |

Responses include:

```http
X-Limit: 100
X-Offset: 0
```

### Error Responses

```json
{
  "error": "Invalid API key"
}
```

Common statuses:

| Status | Meaning |
| --- | --- |
| `400` | Invalid query parameter, path parameter, or JSON body. |
| `401` | Missing, malformed, or invalid API key. |
| `403` | API key scope is insufficient. |
| `404` | Record was not found for update/delete. |
| `500` | Internal server or database error. Details are logged server-side. |

### Partial Updates

`PUT` endpoints update only fields present in the JSON body. Unknown fields and read-only fields return `400`.

### JSON Field Omission

Response models use `omitempty`. Fields with empty strings, numeric zero values, `false`, or `null` relations may be omitted from JSON responses. Timestamp fields can still appear as Go's zero time, `0001-01-01T00:00:00Z`, when the database value is empty. The examples below show representative response shapes, but real responses only include fields that are serialized for the current database row.

Successful update responses use this shape:

```json
{
  "message": "User updated"
}
```

Successful delete responses use this shape:

```json
{
  "message": "User deleted"
}
```

## Users

### User Object

Public response fields:

```json
{
  "id": "1001",
  "code": "E1001",
  "name": "Employee 1001",
  "sex": "M",
  "departmentId": 1,
  "department": {
    "id": 1,
    "name": "Main",
    "supDepartmentId": 1
  },
  "employDate": "2026-01-15T00:00:00Z",
  "duty": "Engineer",
  "isAtt": true,
  "isOverTime": true,
  "isRest": true,
  "mgFlag": 3
}
```

Hidden response fields include password, birthday, phone, mobile, ID card, address, native place, card number, images, and admin group.

Optional public fields that may appear when non-empty: `nation`, `educated`, `polity`, `specialty`, `remark`, `userFlag`, `groupId`, `classFlag`.

### List Users

```http
GET /api/v1/users
```

Query parameters:

| Parameter | Description |
| --- | --- |
| `user_id` | Optional user ID filter. |
| `department_id` | Optional department ID filter. |
| `limit` | Optional page size. |
| `offset` | Optional page offset. |

Example:

```bash
curl -H "Authorization: Bearer read-key" \
  "http://localhost:8080/api/v1/users?department_id=1&limit=50&offset=0"
```

Response `200`:

```json
[
  {
    "id": "1001",
    "code": "E1001",
    "name": "Employee 1001",
    "sex": "M",
    "departmentId": 1,
    "department": {
      "id": 1,
      "name": "Main"
    },
    "employDate": "2026-01-15T00:00:00Z",
    "duty": "Engineer",
    "isAtt": true
  }
]
```

### Create User

```http
POST /api/v1/users
```

Example:

```bash
curl -X POST "http://localhost:8080/api/v1/users" \
  -H "Authorization: Bearer crud-key" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1001",
    "code": "E1001",
    "name": "Employee 1001",
    "sex": "M",
    "departmentId": 1,
    "nation": "RU",
    "employDate": "2026-01-15T00:00:00Z",
    "duty": "Engineer",
    "isAtt": true
  }'
```

Response `201`:

```json
{
  "id": "1001",
  "code": "E1001",
  "name": "Employee 1001",
  "sex": "M",
  "departmentId": 1,
  "nation": "RU",
  "employDate": "2026-01-15T00:00:00Z",
  "duty": "Engineer",
  "isAtt": true
}
```

### Update User

```http
PUT /api/v1/users/{id}
```

Allowed fields:

`code`, `name`, `sex`, `departmentId`, `nation`, `employDate`, `duty`, `educated`, `polity`, `specialty`, `isAtt`, `isOverTime`, `isRest`, `remark`, `mgFlag`, `userFlag`, `groupId`, `classFlag`

Example:

```bash
curl -X PUT "http://localhost:8080/api/v1/users/1001" \
  -H "Authorization: Bearer crud-key" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Employee 1001 Updated",
    "departmentId": 2,
    "isAtt": false
  }'
```

Response `200`:

```json
{
  "message": "User updated"
}
```

### Delete User

```http
DELETE /api/v1/users/{id}
```

Example:

```bash
curl -X DELETE "http://localhost:8080/api/v1/users/1001" \
  -H "Authorization: Bearer crud-key"
```

Response `200`:

```json
{
  "message": "User deleted"
}
```

## Departments

### Department Object

```json
{
  "id": 1,
  "name": "Main"
}
```

`supDepartmentId` and `supDepartment` are included only when present.

### List Departments

```http
GET /api/v1/departments
```

Example:

```bash
curl -H "Authorization: Bearer read-key" \
  "http://localhost:8080/api/v1/departments?limit=100"
```

Response `200`:

```json
[
  {
    "id": 1,
    "name": "Main"
  }
]
```

### Create Department

```http
POST /api/v1/departments
```

Example:

```bash
curl -X POST "http://localhost:8080/api/v1/departments" \
  -H "Authorization: Bearer crud-key" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 10,
    "name": "Production Line",
    "supDepartmentId": 1
  }'
```

Response `201`:

```json
{
  "id": 10,
  "name": "Production Line",
  "supDepartmentId": 1
}
```

### Update Department

```http
PUT /api/v1/departments/{id}
```

Allowed fields:

`name`, `supDepartmentId`

Example:

```bash
curl -X PUT "http://localhost:8080/api/v1/departments/10" \
  -H "Authorization: Bearer crud-key" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Production Line A"
  }'
```

Response `200`:

```json
{
  "message": "Department updated"
}
```

### Delete Department

```http
DELETE /api/v1/departments/{id}
```

Example:

```bash
curl -X DELETE "http://localhost:8080/api/v1/departments/10" \
  -H "Authorization: Bearer crud-key"
```

Response `200`:

```json
{
  "message": "Department deleted"
}
```

## Devices

### Device Object

```json
{
  "id": 1,
  "name": "Warehouse Gate",
  "linkMode": 1,
  "port": 5010,
  "clientNumber": 1,
  "baudRate": 115200,
  "floorId": 1,
  "machineType": 2011
}
```

`commPassword` is hidden from JSON responses.

Optional public fields that may appear when non-empty: `ip`, `recStatus`, `deviceType`, `deviceFlag`, `timezone`.

### List Devices

```http
GET /api/v1/devices
```

Example:

```bash
curl -H "Authorization: Bearer read-key" \
  "http://localhost:8080/api/v1/devices?limit=100"
```

Response `200`:

```json
[
  {
    "id": 1,
    "name": "Warehouse Gate",
    "linkMode": 1,
    "port": 5010,
    "clientNumber": 1,
    "baudRate": 115200,
    "floorId": 1,
    "machineType": 2011
  }
]
```

### Create Device

```http
POST /api/v1/devices
```

Example:

```bash
curl -X POST "http://localhost:8080/api/v1/devices" \
  -H "Authorization: Bearer crud-key" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "name": "Warehouse Gate",
    "linkMode": 1,
    "ip": "192.168.1.50",
    "port": 5010,
    "clientNumber": 1,
    "baudRate": 115200,
    "timezone": "Europe/Berlin"
  }'
```

Response `201`:

```json
{
  "id": 1,
  "name": "Warehouse Gate",
  "linkMode": 1,
  "ip": "192.168.1.50",
  "port": 5010,
  "clientNumber": 1,
  "baudRate": 115200,
  "timezone": "Europe/Berlin"
}
```

### Update Device

```http
PUT /api/v1/devices/{id}
```

Allowed fields:

`name`, `linkMode`, `ip`, `port`, `clientNumber`, `baudRate`, `recStatus`, `floorId`, `machineType`, `deviceType`, `deviceFlag`, `timezone`

Example:

```bash
curl -X PUT "http://localhost:8080/api/v1/devices/1" \
  -H "Authorization: Bearer crud-key" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Warehouse Gate A",
    "ip": "192.168.1.51"
  }'
```

Response `200`:

```json
{
  "message": "Device updated"
}
```

### Delete Device

```http
DELETE /api/v1/devices/{id}
```

Example:

```bash
curl -X DELETE "http://localhost:8080/api/v1/devices/1" \
  -H "Authorization: Bearer crud-key"
```

Response `200`:

```json
{
  "message": "Device deleted"
}
```

## Records

### Record Object

```json
{
  "id": 100,
  "userId": "1001",
  "user": {
    "id": "1001",
    "name": "Employee 1001",
    "departmentId": 1
  },
  "checkTime": "2026-05-04T08:30:00Z",
  "deviceId": 1,
  "device": {
    "id": 1,
    "name": "Warehouse Gate",
    "port": 5010
  },
  "identificationCode": 19,
  "openDoorFlag": true,
  "mask": 2
}
```

`checkType` is included when the related row can be loaded, including rows whose `checkTypeId` is `0`. `user` and `device` are included when their related rows can be loaded.

Optional public fields that may appear when non-zero/non-empty: `workTypeId`, `isChecked`, `isExported`, `temperature`, `whyNoOpen`.

### List Records

```http
GET /api/v1/records
```

Query parameters:

| Parameter | Description |
| --- | --- |
| `record_id` | Optional record ID filter. |
| `user_ids` | Optional comma-separated user IDs. |
| `start_time` | Optional RFC3339 lower bound for `checkTime`. |
| `end_time` | Optional RFC3339 upper bound for `checkTime`. |
| `limit` | Optional page size. |
| `offset` | Optional page offset. |

Example:

```bash
curl -H "Authorization: Bearer read-key" \
  "http://localhost:8080/api/v1/records?user_ids=1001,1002&start_time=2026-05-04T00:00:00Z&end_time=2026-05-04T23:59:59Z&limit=100"
```

Response `200`:

```json
[
  {
    "id": 100,
    "userId": "1001",
    "user": {
      "id": "1001",
      "name": "Employee 1001",
      "departmentId": 1
    },
    "checkTime": "2026-05-04T08:30:00Z",
    "deviceId": 1,
    "device": {
      "id": 1,
      "name": "Warehouse Gate",
      "port": 5010
    },
    "openDoorFlag": true,
    "mask": 2
  }
]
```

### List New Records Since ID

```http
GET /api/v1/records/since
```

Use this endpoint for incremental synchronization from Anviz into another database. It returns records with `id > after_id`, ordered by `id ASC`, so the consumer can store the largest received `id` as the next checkpoint.

Query parameters:

| Parameter | Description |
| --- | --- |
| `after_id` | Required last processed record ID. |
| `limit` | Optional page size. |
| `offset` | Optional page offset. Usually keep `0` for checkpoint-based sync. |

Example:

```bash
curl -H "Authorization: Bearer read-key" \
  "http://localhost:8080/api/v1/records/since?after_id=121000&limit=500"
```

Response `200`:

```json
[
  {
    "id": 121001,
    "userId": "1001",
    "checkTime": "2026-05-04T08:30:00Z",
    "deviceId": 1,
    "device": {
      "id": 1,
      "name": "Warehouse Gate"
    },
    "openDoorFlag": true
  },
  {
    "id": 121002,
    "userId": "1002",
    "checkTime": "2026-05-04T08:32:00Z",
    "deviceId": 1,
    "device": {
      "id": 1,
      "name": "Warehouse Gate"
    },
    "openDoorFlag": true
  }
]
```

Response `400` when `after_id` is missing or invalid:

```json
{
  "error": "after_id is required"
}
```

### Create Record

```http
POST /api/v1/records
```

Example:

```bash
curl -X POST "http://localhost:8080/api/v1/records" \
  -H "Authorization: Bearer crud-key" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 100,
    "userId": "1001",
    "checkTime": "2026-05-04T08:30:00Z",
    "checkTypeId": 1,
    "deviceId": 1,
    "workTypeId": 0,
    "identificationCode": 0,
    "isChecked": false,
    "isExported": false,
    "openDoorFlag": true,
    "temperature": 36.6,
    "whyNoOpen": 0,
    "mask": 0
  }'
```

Response `201`:

```json
{
  "id": 100,
  "userId": "1001",
  "checkTime": "2026-05-04T08:30:00Z",
  "checkTypeId": 1,
  "deviceId": 1,
  "openDoorFlag": true,
  "temperature": 36.6
}
```

### Update Record

```http
PUT /api/v1/records/{id}
```

Allowed fields:

`userId`, `checkTime`, `checkTypeId`, `deviceId`, `workTypeId`, `identificationCode`, `isChecked`, `isExported`, `openDoorFlag`, `temperature`, `whyNoOpen`, `mask`

Example:

```bash
curl -X PUT "http://localhost:8080/api/v1/records/100" \
  -H "Authorization: Bearer crud-key" \
  -H "Content-Type: application/json" \
  -d '{
    "isExported": true,
    "temperature": 36.7
  }'
```

Response `200`:

```json
{
  "message": "Record updated"
}
```

### Delete Record

```http
DELETE /api/v1/records/{id}
```

Example:

```bash
curl -X DELETE "http://localhost:8080/api/v1/records/100" \
  -H "Authorization: Bearer crud-key"
```

Response `200`:

```json
{
  "message": "Record deleted"
}
```

## Check Types

### Check Type Object

```json
{
  "id": 1,
  "char": "I",
  "name": "In"
}
```

### List Check Types

```http
GET /api/v1/check_types
```

Example:

```bash
curl -H "Authorization: Bearer read-key" \
  "http://localhost:8080/api/v1/check_types"
```

Response `200`:

```json
[
  {
    "id": 1,
    "char": "I",
    "name": "In"
  }
]
```

### Update Check Type

```http
PUT /api/v1/check_types/{id}
```

Allowed fields:

`char`, `name`

Example:

```bash
curl -X PUT "http://localhost:8080/api/v1/check_types/1" \
  -H "Authorization: Bearer crud-key" \
  -H "Content-Type: application/json" \
  -d '{
    "char": "I",
    "name": "Check in"
  }'
```

Response `200`:

```json
{
  "message": "Check type updated"
}
```
