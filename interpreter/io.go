package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var buffer = strings.Builder{}

func input(memory Memory) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	scanner.Scan()
	memory.Add(len(scanner.Text()))
}

// output the ASCII value of the character if it's between 32 and 126, otherwise
// print the value in square brackets.
func output(value int) {
	c := fmt.Sprintf("[%d]", value)

	if value < 0 {
		value = -value
	}

	value %= 127

	switch value {
	case 0, 9, 10, 15:
		c = string(value)
	}

	if value >= 32 {
		c = string(value)
	}

	buffer.WriteString(c)
}

func display() {
	fmt.Println(buffer.String())
}
