package interpreter

// Memory map for an Auld script.
type Memory interface {
	// Add the value to the current memory cell. A negative value results in
	// subtraction.
	Add(value int)

	// Value of the current cell
	Value() int

	// Next memory cell is now the current memory cell.
	Next()

	// Previous memory cell is now the current memory cell.
	Previous()

	// Move the pointer n cells. Use a negative value of n to move the pointer
	// backward. The memory will wrap if the pointer under or overflows.
	Move(n int)

	// Malloc the memory. This is a destructive operation which will allocate a
	// new chunk of memory, discarding the contents of the old memory location
	// and resetting the pointer to 0.
	Malloc(size int)
}

type memory struct {
	ptr   int
	cells []int
}

// NewMemory returns an empty, 1 cell memory structure. Calls on the memory
// structure are safe but it is intended that the memory be resized using
// Malloc straight away.
func NewMemory() Memory {
	return &memory{cells: make([]int, 1)}
}

func (m *memory) Add(value int) {
	m.cells[m.ptr] += value
}

func (m *memory) Value() int {
	return m.cells[m.ptr]
}

func (m *memory) Next() {
	m.Move(1)
}

func (m *memory) Previous() {
	m.Move(-1)
}

func (m *memory) Move(n int) {
	size := len(m.cells)

	m.ptr += n

	for m.ptr >= size {
		m.ptr -= size
	}

	for m.ptr < 0 {
		m.ptr += size
	}
}

func (m *memory) Malloc(size int) {
	if size < 1 {
		panic("malloc failed")
	}

	m.cells = make([]int, size)
	m.ptr = 0
}
