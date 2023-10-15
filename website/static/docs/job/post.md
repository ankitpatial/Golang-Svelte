## POST: Job
Endpoint to post job data is `https://www.roofixstaging.com/external/job` and requires [oauth](/docs/oauth) token to talk.

```http request
POST https://www.roofixstaging.com/external/job
Content-Type: application/json
Authorization: Bearer 1Lg1DRDO/+bwv0YjapAnDfJpnPy6igQmdUpEY7CfB1u3DSFBzNvMsFwEU196doIEddmPO...

{
  "homeOwner": {
    "firstName": "Peeter",
    "lastName": "Pan",
    "streetNumber": "548",
    "streetName": "West 22nd Street",
    "city": "New York",
    "state": "New York",
    "zip": "10011",
    "latitude": 40.747812,
    "longitude": -74.007014
  },
  "salesRep": {
    "companyName": "co name",
    "firstName": "Rep",
    "lastName": "Name",
    "email": "rep@domain.com",
    "mobile": "123 123 1234"
  },
  "currentMaterial": "Clay Tile",
  "newRoofingMaterial": "Concrete Tile",
  "lowSlope": true,
  "currentMaterialLowSlope": "Rolled Roofing",
  "newRoofingMaterialLowSlope": "Coating",
  "redeck": true,
  "layers": 3,
  "layer2Material": "Rolled Roofing",
  "layer3Material": "Wood Shakes",
  "webhookURL": "https://domain.com/url-to-post-job-status-update",
  "measurementType": "PrimaryStructureOnly"
}
```
[check this](/docs/job/data-structure) for JSON Body data validation rules

**successful response** example

```json
{
  "data": {
    "jobId": "9b356243-439a-4bf6-b8ea-2565e26f3f88"
  }
}
```

Bad Request **failed data validation** example

```json
{
  "error": "data validation errors",
  "validationErrors": [
    {
      "field": "homeOwner.firstName",
      "validation": "required",
      "message": "required data field"
    },
    {
      "field": "homeOwner.lastName",
      "validation": "required",
      "message": "required data field"
    },
    {
      "field": "salesRep.email",
      "validation": "email",
      "message": "'s' is not valid email address"
    },
    {
      "field": "company.name",
      "validation": "required",
      "message": "required data field"
    },
    {
      "field": "currentMaterial",
      "validation": "valid",
      "message": "'Some Value' is not allowed"
    },
    {
      "field": "layers",
      "validation": "oneof",
      "message": "value must be one of '1 2 3'."
    }
  ]
}

```
