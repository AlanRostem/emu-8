package instruction

const singleInstKey = 0xFF

type InstructionTable map[uint8]Instruction

func NewInstructionTable() InstructionTable {
	return make(map[uint8]Instruction)
}

func NewSingularInstructionMap(inst Instruction) InstructionTable {
	return map[uint8]Instruction{singleInstKey: inst}
}

func (im InstructionTable) IsSingle() bool {
	_, ok := im[singleInstKey]
	return ok
}

func (im InstructionTable) Single() Instruction {
	if !im.IsSingle() {
		panic("cannot get categorized instruction when map is singluar")
	}
	return im[singleInstKey]
}

func (im InstructionTable) Get(cat uint8) Instruction {
	if im.IsSingle() {
		panic("cannot get multiple instructions for singular map")
	}
	return im[cat]
}

func (im InstructionTable) Add(subcat uint8, inst Instruction) {
	if im.IsSingle() {
		panic("cannot add more than one instruction to a singular map")
	}
	im[subcat] = inst
}
