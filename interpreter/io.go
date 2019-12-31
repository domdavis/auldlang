package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var buffer = strings.Builder{}

// Output the contents of the output buffer. Typically this is done once a
// script is run.
func Output() string {
	return buffer.String()
}

func input(memory Memory) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	scanner.Scan()
	memory.Add(len(scanner.Text()))
}

// output the ASCII value of the character if it's between 32 and 126, otherwise
// print "?".
func output(value int) {
	c := "?"

	if value < 0 {
		value = -value
	}

	value %= 127

	if value >= 32 {
		c = string(value)
	}

	buffer.WriteString(fmt.Sprintf("[%d '%s']", value, c))
}
