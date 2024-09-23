package i18ne

import (
	"errors"
)

// UnimplementedError is a generic error that should be thrown by
// functions that haven't been implemented.
func UnimplementedError() (err error) {
	err = errors.New("Unimplemented")
	return
}

// InvalidArgumentError is a generic error that should be thrown by
// functions that receive invalid arguments.
func InvalidArgumentError() (err error) {
	err = errors.New("Invalid argument")
	return
}
