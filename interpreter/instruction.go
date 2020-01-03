package interpreter

// Instruction holds an Auld instruction and it's optional terminator.
type Instruction interface {
	// Execute the instruction and it's terminator.
	Execute()
}

type instruction struct {
	function   func()
	terminator func()
}

func (i *instruction) Execute() {
	i.function()
	i.terminator()
}

func noop() {}
