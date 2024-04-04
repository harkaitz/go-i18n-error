package i18ne

import (
	"errors"
)

// JoinErrors returns nil if errl is empty, or joins the errors
// with newlines.
func JoinErrors(errl []error) (err error) {
	if len(errl) == 0 {
		return nil
	}
	var s string
	for _, err = range errl {
		s += err.Error() + "\n"
	}
	err = errors.New(s)
	return
}

