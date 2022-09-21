package argtool

import "fmt"

type InvalidArgumentError struct {
	Name   string
	Reason string
	Err    error
}

func (e *InvalidArgumentError) Error() string {
	if len(e.Reason) > 0 {
		return fmt.Sprintf("invalid argument '%s'; %s", e.Name, e.Reason)
	}
	return fmt.Sprintf("invalid argument '%s'", e.Name)
}

// Unwrap returns the underlying error.
func (e *InvalidArgumentError) Unwrap() error { return e.Err }
