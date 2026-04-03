package instruction

import "github.com/AlanRostem/mu-8/decode"

var TableClass6 = NewSingularInstructionMap(ldVxByte)
var TableClassA = NewSingularInstructionMap(ldIAddr)

// TableAll contains instruction maps mapped to a class number.
var TableAll = map[decode.Class]InstructionTable{
	decode.Class6: TableClass6,
	decode.ClassA: TableClassA,
}
