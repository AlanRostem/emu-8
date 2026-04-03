package instruction

import "github.com/AlanRostem/mu-8/mu8"

const singleInstKey = 0xFF

type instructionTable map[mu8.DByte]Instruction

func newInstructionTable() instructionTable {
	return make(map[mu8.DByte]Instruction)
}

func NewSingularInstructionMap(inst Instruction) instructionTable {
	return map[mu8.DByte]Instruction{singleInstKey: inst}
}

func (im instructionTable) IsSingle() bool {
	_, ok := im[singleInstKey]
	return ok
}

func (im instructionTable) Single() Instruction {
	if !im.IsSingle() {
		panic("cannot get categorized instruction when map is singluar")
	}
	return im[singleInstKey]
}

func (im instructionTable) Get(id mu8.DByte) Instruction {
	if im.IsSingle() {
		panic("cannot get multiple instructions for singular map")
	}
	return im[id]
}

func (im instructionTable) Add(id mu8.DByte, inst Instruction) {
	if im.IsSingle() {
		panic("cannot add more than one instruction to a singular map")
	}
	im[id] = inst
}
