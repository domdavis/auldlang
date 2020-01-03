package interpreter

import (
	"errors"
	"fmt"
	"strings"
)

// Script is a type that can be run by the interpreter.
type Script interface {
	// Run the script.
	Run()
}

type script struct {
	debug bool
	line  int
	lines []string
	Memory
}

// ErrInvalidScript is returned if the Auld script is invalid (currently only
// if it's empty)
var ErrInvalidScript = errors.New("invalid Auld script")

// New Script type that can be interpreted.
func New(code string) (Script, error) {
	s := &script{Memory: NewMemory()}
	lines := strings.Split(code, "\n")

	if len(lines) == 0 {
		return s, ErrInvalidScript
	}

	s.lines = lines

	return s, nil
}

func (s *script) Run() {
	for s.line < len(s.lines) {
		i := s.next()

		i.Execute()
	}

	fmt.Println(s.Memory)
	display()
}

//nolint: gocyclo
func (s *script) next() Instruction {
	fmt.Printf("%d: %s\n", s.line, s.lines[s.line])

	i := &instruction{
		function:   noop,
		terminator: noop,
	}
	line := s.lines[s.line]

	s.line++

	if len(line) == 0 {
		return i
	}

	command := line[:len(line)-1]

	switch line[len(line)-1:] {
	case "?":
		i.terminator = func() { input(s.Memory); s.Memory.Next(); s.dump() }
	case "!":
		i.terminator = func() { s.Memory.Next(); s.dump() }
	case ";":
		i.terminator = func() { s.Memory.Previous(); s.dump() }
	case ",":
		i.terminator = func() { s.Memory.Add(1); s.dump() }
	case ".":
		i.terminator = func() { s.Memory.Add(-1); s.dump() }
	default:
		command = line
	}

	if arg := tokenize("Happy", command); arg >= 0 {
		i.function = func() { s.Memory.Malloc(arg) }
	} else if arg := tokenize("Should auld acquaintance be forgot", command); arg >= 0 {
		i.function = func() { s.rpt() }
	} else if arg := tokenize("We'll", command); arg >= 0 {
		i.function = func() { s.Add(-arg); s.dump() }
	} else if arg := tokenize("And", command); arg >= 0 {
		i.function = func() { s.Add(arg); s.dump() }
	} else if arg := tokenize("Frae", command); arg >= 0 {
		i.function = func() { s.Move(arg); s.dump() }
	} else if arg := tokenize("Sin auld lang syne", command); arg >= 0 {
		i.function = func() { output(s.Value()); s.Move(arg); s.dump() }
	} else if arg := tokenize("For auld lang syne", command); arg >= 0 {
		i.function = func() { output(s.Value()); s.Move(-arg); s.dump() }
	} else if arg := tokenize("We", command); arg >= 0 {
		i.function = func() { s.jmp("But", arg); s.dump() }
	} else if arg := tokenize("But", command); arg >= 0 {
		i.function = func() { s.rtn("We", arg); s.dump() }
	} else if arg := tokenize("Kevlin", command); arg >= 0 {
		s.debug = true
		s.dump()
	} else {
		panic(fmt.Sprintf("syntax error on line %d: %s", s.line-1, command))
	}

	return i
}

func (s *script) rpt() {
	s.dump()
	for s.Memory.Value() != 0 {
		i := s.next()
		s.line--
		i.Execute()
	}

	s.line++
}

func (s *script) jmp(keyword string, condition int) {
	if s.Memory.Value() >= condition {
		return
	}

	for i := s.line; i < len(s.lines); i++ {
		if arg := tokenize(keyword, s.lines[i]); arg >= 0 {
			s.line = i + 1
			return
		}
	}

	s.line = len(s.lines)
}

func (s *script) rtn(keyword string, condition int) {
	if s.Memory.Value() >= condition {
		return
	}

	for i := s.line; i > 0; i-- {
		if arg := tokenize(keyword, s.lines[i]); arg >= 0 {
			s.line = i
			return
		}
	}

	s.line = 0
}

func (s *script) dump() {
	if s.debug {
		fmt.Println(s.Memory)
	}
}

func tokenize(keyword, line string) int {
	token := strings.ToLower(keyword) + " "
	input := strings.ToLower(strings.ReplaceAll(line, ",", " "))

	switch {
	case strings.HasPrefix(input, token):
		return len(input) - len(token)
	case input == strings.ToLower(keyword):
		return len(input) - len(keyword)
	default:
		return -1
	}
}
