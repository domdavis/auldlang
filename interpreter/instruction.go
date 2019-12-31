package interpreter

// Instruction holds an Auld instruction and it's optional terminator.
type Instruction interface {
	// Execute the instruction.
	Execute()

	// Terminator will execute the termination instruction, if there is one.
	Terminator()
}

type instruction struct {
	function   func()
	terminator func()
}

func (i *instruction) Execute() {
	i.function()
}

func (i *instruction) Terminator() {
	i.terminator()
}

func noop() {}
