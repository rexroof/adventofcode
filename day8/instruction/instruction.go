package instruction

type Instruction struct {
	Operation string // acc jmp or nop
	Argument  int    // +XX or -XX
	Executed  int    // times executed
}

func New(operation string, argument int) Instruction {
	return Instruction{operation, argument, 0}
}

func (i Instruction) Exec() Instruction {
	i.Executed += 1
	return i
}

//method for returning operation and offset value?
