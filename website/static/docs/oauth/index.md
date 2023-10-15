## Obtain Access Token

To obtain access token

```http request
POST https://www.roofixstaging.com/token
Content-Type: application/json

{
  "grant_type": "password",
  "username": "api-username",
  "password": "api-user-password"
} 
```

sample response:

```json
{
  "access_token": "Ey8m5Z321cjZ3RWxysDPrDjAneghbsFEsSBADEqhjWgm3DvMsFwEU1SFBzN96fYRRKI6...",
  "refresh_token": "QZOvQgEgnZ+h4mT6n023JU67FqpTbTux3RzBsT48BTru3Dv4pPBBZhGYd3tdqadcFPI...",
  "token_type": "Bearer",
  "expires_in": 604800,
  "properties": null
}

```

**expires_in** is in seconds

## Refresh Token

get refresh token

```http request
POST https://www.roofixstaging.com/token
Content-Type: application/json

{
  "grant_type": "refresh_token",
  "refresh_token": "QZOvQgEgnZ+h4mT6n023JU67FqpTbTux3RzBsT48BTru3Dv4pPBBZhGYd3tdqadcFPI..."
}

```

sample response:

```json
{
  "access_token": "1Lg1DRDO/+bwv0YjapAnDfJpnPy6igQmdUpEY7CfB1u3DSFBzNvMsFwEU196doIEddmPO...",
  "refresh_token": "OKo4l0mK6qsIFVbfrA1y4N8aKzG2tkuvaCR63DviBrE4NL4pPBBZhGYd3tdqadcFPI...",
  "token_type": "Bearer",
  "expires_in": 604800,
  "properties": null
}
```
