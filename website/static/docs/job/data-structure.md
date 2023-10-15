## POST: JOB Data

Following is full json structure to POST a job info.

```json
{
  "homeOwner": {
    "firstName": "",
    "lastName": "",
    "streetNumber": "",
    "streetName": "",
    "city": "",
    "state": "",
    "zip": "",
    "latitude": 0.0,
    "longitude": 0.0
  },
  "salesRep": {
    "companyName": "",
    "firstName": "",
    "lastName": "",
    "email": "",
    "mobile": ""
  },
  "currentMaterial": "",
  "newRoofingMaterial": "",
  "lowSlope": true,
  "currentMaterialLowSlope": "",
  "newRoofingMaterialLowSlope": "",
  "redeck": true,
  "layers": 3,
  "layer2Material": "",
  "layer3Material": "",
  "webhookURL": "",
  "measurementType": ""
}
```

## Owner Info

All fields are required

```json
 {
  "firstName": "required",
  "lastName": "required",
  "streetNumber": "required",
  "streetName": "required",
  "city": "required",
  "state": "required",
  "zip": "required",
  "latitude": "must be valid latitude",
  "longitude": "must be valid longitude"
}
```

## Sales Rep info

All fields are required

```json
{
  "companyName": "required",
  "firstName": "required",
  "lastName": "required",
  "email": "required & valid email",
  "mobile": "required"
}
``` 

## Other Fields info

- `currentMaterial` is **required**. Allowed values are:
    - 3-Tab Asphalt Shingles
    - Architectural Asphalt Shingles
    - Clay Tile
    - Concrete Tile
    - Wood Shakes
    - Metal Shakes
    - Metal Tile
    - Standing Seam Metal
    - Slate
    - Metal R-Panel Exposed Fastener
    - Low Slope Only
- `newRoofingMaterial` is **required** in most of the case but if **currentMaterial = 'Low Slope Only'** then it must be
  empty. Allowed values are:
    - Best Value - Architectural
    - More Expensive - Premium Architectural
    - Standing Seam Metal
    - Concrete Tile
    - Clay Tile (Barrel)
- `lowSlope` is **required** and must be **true** or **false**. It must be true if **currentMaterial = 'Low Slope Only'
  **
- `currentMaterialLowSlope` is **required** if **lowSlope = true** else it must be an empty string. Allowed values are:
    - Rolled Roofing
    - Asphalt Shingles
    - Black Membrane
    - White Membrane
    - Gravel BUR
- `newRoofingMaterialLowSlope` is **required** if **lowSlope = true** else it must be an empty string. Allowed values
  are:
    - Modified Bitumen
    - Coating
- `redeck` is **required** and must be **true** or **false**
- `layers` is **required** and must be **1** or **2** or **3**
- `layer2Material` is **required** only if **layers** value is **2** or **3**. Allowed values are:
    - 3-Tab Asphalt Shingles
    - Architectural Asphalt Shingles
    - Asphalt Shingles
    - Wood Shakes
    - Rolled Roofing
    - Black Membrane
    - White Membrane
    - Low Slope Only
- `layer3Material` is **required** only if **layers** value is **3**. Allowed values are:
    - 3-Tab Asphalt Shingles
    - Architectural Asphalt Shingles
    - Asphalt Shingles
    - Wood Shakes
    - Rolled Roofing
    - Black Membrane
    - White Membrane
    - Low Slope Only
- `webhookURL` is **optional** but any given value must be in a valid url format. Webhook URL will be used as webhook to
  **POST** job status updates.
- `measurementType` is **required**. Allowed values are:
    - PrimaryStructureOnly
    - PrimaryPlusDetachedGarage
    - AllStructuresOnParcel
