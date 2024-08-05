package main

import (
	"context"
	"github.com/mortedecai/all-text-wrestling/internal/app"
)

func prepare() (app.ATW, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	a, _ := app.New(app.ATW_CONSOLE, ctx, nil)
	return a, cancel
}

func main() {
	atw, cancel := prepare()
	// TODO: {mortedecai} Remove this once actual app-iness happens.
	_, _ = atw.WriteString("Hello\n")
	cancel()
}
