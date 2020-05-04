package yael

// E defines the standard yael error type which conforms to errors.Wrapper
// and error interfaces
type E struct {
	Code   string                 `json:"code"`
	Reason *E                     `json:"reason,omitempty"`
	Meta   map[string]interface{} `json:"meta,omitempty"`

	sensitive bool
}

// New returns a new instance of E with code set
func New(code string) *E {
	return &E{
		Code: code,
	}
}

// WithMeta sets key/value metadata on the error directly
func (e *E) WithMeta(key string, value interface{}) *E {
	if e.Meta == nil {
		e.Meta = make(map[string]interface{}, 1)
	}

	e.Meta[key] = value
	return e
}

// WithReason sets the underlying error reason
func (e *E) WithReason(reason *E) *E {
	e.Reason = reason
	return e
}

// Error conforms to the standar error interface
func (e E) Error() string {
	return e.Code
}

// Unwrap conforms to the errors.Wrapper interface
func (e *E) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Reason
}
