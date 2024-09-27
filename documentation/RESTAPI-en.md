# REST API Documentation

<img align="right" width="128px" src="../.media/homburg.png">

## Summary

| USER Endpoint                | Request | PGA required |
|:----------------------------:|:-------:|:-----:|
| [/api/register](#register)   | POST    | - |

# USER endpoints

<h3 id="register">
:book:&nbsp;&nbsp/api/register
</h3>

The authentication scheme is PGA (Pretty Good Authentication).  
We register a public key instead of an id and password.

| Field       | Type    | Required |
|:-----------:|:-------:|:--------:|
| algorithm   | string  | no | 
| public-key  | string  | yes |

```json
{
  "message": "not implemented yet"
}
```

| Code | Status         | Message |
|:----:|:--------------:|:--------|
| 501  | Not Implemented | not implemented yet |

## TODO

Bugs:

- None.
