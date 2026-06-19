# yael

![Go](https://github.com/snicol/yael/workflows/test/badge.svg)

Yet Another Error Library, used for structured sentinel errors. Zero dependencies.

## Usage

```go
// Create errors
err := yael.New("permission_denied").WithMeta("user", 123)

// Wrap one or more reasons
notAdminErr := yael.New("not_admin")
noRoleErr   := yael.New("no_role")
err = err.WithReasons(notAdminErr, noRoleErr)

// Satisfies Go's error interface
fmt.Println(err) // permission_denied

// Sentinel matching by code (errors.Is traverses all reasons)
errors.Is(err, notAdminErr) // true
errors.Is(err, noRoleErr)   // true

// JSON representable
{
    "code": "permission_denied",
    "meta": { "user": 123 },
    "reasons": [
        { "code": "not_admin" },
        { "code": "no_role" }
    ]
}
```

## Built-in error codes

| Constant             | Code                  | HTTP |
|----------------------|-----------------------|------|
| `BadRequest`         | `bad_request`         | 400  |
| `Unauthorized`       | `unauthorized`        | 401  |
| `Forbidden`          | `forbidden`           | 403  |
| `NotFound`           | `not_found`           | 404  |
| `MethodNotAllowed`   | `method_not_allowed`  | 405  |
| `Conflict`           | `conflict`            | 409  |
| `NotImplemented`     | `not_implemented`     | 501  |
| `ServiceUnavailable` | `service_unavailable` | 503  |

Use `yael.StatusCode(e)` to map an error to its HTTP status code.
