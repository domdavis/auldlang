package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/domdavis/auldlang/interpreter"
)

// ErrUsage is returned when there are incorrect arguments.
var ErrUsage = errors.New("usage: auld <file>")

// Exit codes
const (
	FatalUsage = iota + 1
	FatalRead
	FatalInterpret
	FatalSyntax
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			exit(FatalSyntax, fmt.Errorf("%v", r))
		}
	}()

	if len(os.Args) != 2 {
		exit(FatalUsage, ErrUsage)
	} else if b, err := ioutil.ReadFile(os.Args[1]); err != nil {
		exit(FatalRead, fmt.Errorf("failed to read %s: %w", os.Args[1], err))
	} else if i, err := interpreter.New(string(b)); err != nil {
		exit(FatalInterpret, fmt.Errorf("failed to start interpreter: %w", err))
	} else {
		i.Run()
	}
}

func exit(code int, err error) {
	if code > 0 {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}

	os.Exit(code)
}
