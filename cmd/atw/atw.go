package main

import (
	"os"
)

type atwApp struct {
	// TODO: {mortedecai} Move this to a more appropriate location in future.

	out *os.File
}

func New(out *os.File) *atwApp {
	return &atwApp{out: out}
}

func (atw *atwApp) Write(msg string) (n int, err error) {
	// TODO: {mortedecai} Delete this once more testable code is in place.
	return atw.out.WriteString(msg)
}

func main() {
	atw := New(os.Stdout)
	atw.Write("Hello")
}
