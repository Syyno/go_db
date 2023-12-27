package domainErr

import "errors"

var (
	ErrUnknownOperation   = errors.New("Operation is unknown")
	ErrArgsCountIsInvalid = errors.New("Invalid amount of args")
	ErrArgsCountMismatch  = errors.New("Command was recognized, but args count is invalid for requested command")
)
