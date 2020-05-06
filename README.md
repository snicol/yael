# yael

![Go](https://github.com/snicol/yael/workflows/Go/badge.svg)

Yet Another Error Library

> Yael (Hebrew: יעל‎, pronounced [jaˈʔel]; also spelled Jael) is a female given
name, from the Hebrew meaning "Nubian Ibex".

Similar code is used across a lot of my side projects so I decided to
standardise it.

## Usage

```go
// Errors with metadata
err := yael.New("permission_denied").WithMeta("user", 123)

// Wrapping errors
notAdminErr := yael.New("not_admin")
err = err.WithReason(notAdminErr)

// Satisfies Go's error interface
fmt.Println(err) // permission_denied

// Also Go 1.13+ error wrapping
errors.Is(err, notAdminErr) // true

// JSON representable:
{
    "code": "permission_denied",
    "meta": {
        "user": 123
    },
    "reason": {
        "code": "not_admin"
    }
}
```
