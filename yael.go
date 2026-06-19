// Package yael provides a structured error type that conforms to the errors.Wrapper
// and error interfaces.
package yael

// E defines the standard yael error type which conforms to errors.Wrapper
// and error interfaces.
type E struct {
	Code    string         `json:"code"`
	Reasons []*E           `json:"reasons,omitempty"`
	Meta    map[string]any `json:"meta,omitempty"`
}

// New returns a new instance of E with code set.
func New(code string) *E {
	return &E{
		Code: code,
	}
}

// WithMeta sets key/value metadata on the error directly.
func (e *E) WithMeta(key string, value any) *E {
	if e.Meta == nil {
		e.Meta = make(map[string]any, 1)
	}

	e.Meta[key] = value

	return e
}

// WithReasons sets one or more underlying error reasons.
func (e *E) WithReasons(reasons ...*E) *E {
	e.Reasons = append(e.Reasons, reasons...)

	return e
}

// Error conforms to the standard error interface.
func (e E) Error() string {
	return e.Code
}

// Is reports whether e matches target by comparing error codes.
func (e *E) Is(target error) bool {
	if e == nil {
		return false
	}

	t, ok := target.(*E)
	if !ok {
		return false
	}

	return e.Code == t.Code
}

// Unwrap returns the list of reasons so errors.Is and errors.As traverse them.
func (e *E) Unwrap() []error {
	if e == nil || len(e.Reasons) == 0 {
		return nil
	}

	errs := make([]error, len(e.Reasons))
	for i, r := range e.Reasons {
		errs[i] = r
	}

	return errs
}
