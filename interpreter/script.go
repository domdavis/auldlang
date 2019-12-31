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
		i.Terminator()
	}
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
		i.terminator = func() { input(s.Memory); s.Memory.Next() }
	case "!":
		i.terminator = s.Memory.Next
	case ";":
		i.terminator = s.Memory.Previous
	case ",":
		i.terminator = func() { s.Memory.Add(1) }
	case ".":
		i.terminator = func() { s.Memory.Add(-1) }
	default:
		command = line
	}

	if arg := tokenize("Happy", command); arg >= 0 {
		i.function = func() { s.Memory.Malloc(arg) }
	} else if arg := tokenize("Should auld acquaintance be forgot", command); arg >= 0 {
		i.function = func() { s.rpt() }
	} else if arg := tokenize("We'll", command); arg >= 0 {
		i.function = func() { s.Add(arg) }
	} else if arg := tokenize("And", command); arg >= 0 {
		i.function = func() { s.Add(-arg) }
	} else if arg := tokenize("Frae", command); arg >= 0 {
		i.function = func() { s.Move(arg) }
	} else if arg := tokenize("Sin auld lang syne", command); arg >= 0 {
		i.function = func() { output(s.Value()); s.Move(arg) }
	} else if arg := tokenize("For auld lang syne", command); arg >= 0 {
		i.function = func() { output(s.Value()); s.Move(-arg) }
	} else if arg := tokenize("We", command); arg >= 0 {
		i.function = func() { s.jmp("But", arg) }
	} else if arg := tokenize("But", command); arg >= 0 {
		i.function = func() { s.rtn("We", arg) }
	} else {
		panic(fmt.Sprintf("syntax error on line %d: %s", s.line-1, command))
	}

	return i
}

func (s *script) rpt() {
	for s.Memory.Value() != 0 {
		i := s.next()
		s.line--
		i.Execute()
		i.Terminator()
	}
}

func (s *script) jmp(keyword string, condition int) {
	if s.Memory.Value() >= condition {
		return
	}

	for i := s.line + 1; i < len(s.lines); i++ {
		if strings.HasPrefix(s.lines[i], keyword) {
			s.line = i
			return
		}
	}

	s.line = len(s.lines)
}

func (s *script) rtn(keyword string, condition int) {
	if s.Memory.Value() <= condition {
		return
	}

	for i := s.line - 1; i > 0; i-- {
		if strings.HasPrefix(s.lines[i], keyword) {
			s.line = i
			return
		}
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
