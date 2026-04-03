package instruction

import "github.com/AlanRostem/mu-8/decode"

var tableClass6 = NewSingularInstructionMap(ldVxByte)
var tableClassA = NewSingularInstructionMap(ldIAddr)

// tableAll contains instruction maps mapped to a class number.
var tableAll = map[decode.Class]instructionTable{
	decode.Class6: tableClass6,
	decode.ClassA: tableClassA,
}
