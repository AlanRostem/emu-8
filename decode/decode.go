package decode

import "github.com/AlanRostem/mu-8/mu8"

type Argument struct {
	NibblePosition uint8
	NibbleSize     uint8
	Value          uint16
}

type DecodeInfo struct {
	Class     Class  // First nibble which is can be the same for several opcodes
	Idenitity uint16 // Usually the last few nibbles of the opcode
	Args      []uint16
}

func Decode(instruction mu8.DByte) DecodeInfo {
	nibbles := instruction.GetNibbles()
	class := Class(nibbles[0])
	// classInfo := Classes[class]
	// for i, argSize := range classInfo.Args {
	// 	nibIdx := i + 1
	// }
	return DecodeInfo{
		Class: class,
	}
}
