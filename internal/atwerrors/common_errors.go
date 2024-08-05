package atwerrors

import (
	"errors"
	"fmt"
)

var (
	// NotYetImplemented is a common placeholder error to facilitate simple TDD.
	NotYetImplemented = errors.New("not yet implemented")
)

// Embed creates  an embedded error string in a specific format.
func Embed(err error, msg string) error {
	return fmt.Errorf("%w: %s", err, msg)
}
