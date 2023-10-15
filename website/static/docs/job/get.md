### GET Job Status

fetch already submitted job status

```http request
GET https://www.roofixstaging.com/external/job/b9b0652a...
Authorization: Bearer 1Lg1DRDO/+bwv0YjapAnDfJpnPy6igQmdUpEY7CfB1u3DSFBzNvMsFwEU196doIEddmPO...

```

successful response example

```json
{
  "data": {
    "status": "CREATED",
    "createdAt": "2022-11-01T10:42:23Z",
    "updatedAt": "2022-11-01T10:42:23Z",
    "statusHistory": [
      {
        "status": "CREATED",
        "createdAt": "2022-11-01T10:42:23Z"
      }
    ]
  }
}
```

**Bad Request**, invalid jobID is passed on

```json
{
  "error": "invalid jobID: b9b0652a..."
}
```

**Bad Request**, jobID does not belong to API User

```json
{
  "error": "access denied to jobID: b9b0652a..."
}
```

**Server Error** is not supposed to happen but may happen if there is some issue with DB connection

```json
{
  "error": "failed to find Job info"
}
```
