package system

type System struct {
	Registers *RegisterFile
	Memmory   *MemoryBank
}

func New() *System {
	return &System{
		Registers: newRegisterFile(),
		Memmory:   newMemoryBank(),
	}
}
