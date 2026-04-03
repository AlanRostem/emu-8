package system

import "github.com/AlanRostem/mu-8/mu8"

const MemorySize = mu8.Uint12Max

type MemoryBank struct {
	data [MemorySize]mu8.Byte
}

func newMemoryBank() *MemoryBank {
	return &MemoryBank{
		data: [MemorySize]mu8.Byte{},
	}
}

func (m *MemoryBank) Write(addr mu8.Uint12, value mu8.Byte) {
	m.data[addr.Value()] = value
}

func (m *MemoryBank) Read(addr mu8.Uint12) mu8.Byte {
	return m.data[addr.Value()]
}
