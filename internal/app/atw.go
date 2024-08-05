package app

import (
	"context"
	"io"

	"github.com/mortedecai/all-text-wrestling/internal/atwerrors"
)

// ATW is the application interface for the all text wrestling game.
type ATW interface {
	io.StringWriter
	Start() error
	Stop() error
}

// atwVersion is a non-exported type to ensure only implemented app versions are supported.
type atwVersion string

const (
	ATW_CONSOLE = atwVersion("console")
)

func New(_ atwVersion, _ context.Context, _ any) (ATW, error) {
	return nil, atwerrors.NotYetImplemented
}
