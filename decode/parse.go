package decode

import (
	"github.com/AlanRostem/mu-8/instruction"
	"github.com/AlanRostem/mu-8/mnemonic"
	"github.com/AlanRostem/mu-8/mu8"
)

type ParseInfo struct {
	// Nibbles is an array of the full opcode separated into each 4bit value
	Nibbles     [4]uint8
	Instruction instruction.Instruction
}

func (pi *ParseInfo) Category() mnemonic.Category {
	return mnemonic.Category(pi.Nibbles[0])
}

func ParseOpcode(opcode mu8.DByte) ParseInfo {
	info := ParseInfo{}
	for i := range uint8(4) {
		// going in reverse so that nibbles are read from left to right
		info.Nibbles[4-i] = opcode.Nibble(i)
	}
	category := mnemonic.Category(info.Nibbles[0])
	instructionMap, ok := instruction.TableAll[category]
	if !ok {
		panic("instruction category not found")
	}
	if instructionMap.IsSingle() {
		info.Instruction = instructionMap.Single()
	} else {
		// TODO: get instruction based on subcategory
	}
	return info
}
