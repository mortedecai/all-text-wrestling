package console

import (
	"context"
	"io"
	"os"

	"github.com/mortedecai/all-text-wrestling/internal/atwerrors"
)

type ConsoleATW struct {
	ctx context.Context
	io.StringWriter
	in *os.File
}

func New(ctx context.Context, _ any) (*ConsoleATW, error) {
	return &ConsoleATW{
		ctx:          ctx,
		StringWriter: os.Stdout,
		in:           os.Stdin,
	}, nil
}

func (c *ConsoleATW) Start() error {
	return atwerrors.Embed(atwerrors.NotYetImplemented, "Start() error")
}

func (c *ConsoleATW) Stop() error {
	return atwerrors.Embed(atwerrors.NotYetImplemented, "Stop() error")
}
