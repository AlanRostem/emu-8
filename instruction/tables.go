package instruction

import "github.com/AlanRostem/mu-8/mnemonic"

var MapCategory6 = NewSingularInstructionMap(ldVxByte)
var MapCategoryA = NewSingularInstructionMap(ldIAddr)

// TableAll contains instruction maps mapped to a category number.
var TableAll = map[mnemonic.Category]InstructionMap{
	mnemonic.Category6: MapCategory6,
	mnemonic.CategoryA: MapCategoryA,
}
