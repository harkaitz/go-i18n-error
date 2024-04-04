package i18ne

// UnimplementedError is a generic error that should be thrown by
// functions that haven't been implemented.
func UnimplementedError() (err error) {
	err = newError(l("Unimplemented"))
	return
}

// InvalidArgumentError is a generic error that should be thrown by
// functions that receive invalid arguments.
func InvalidArgumentError() (err error) {
	err = newError(l("Invalid argument"))
	return
}
