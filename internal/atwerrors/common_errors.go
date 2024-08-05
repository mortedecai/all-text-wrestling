package atwerrors

import "errors"

var (
	// NotYetImplemented is a common placeholder error to facilitate simple TDD.
	NotYetImplemented = errors.New("not yet implemented")
)
