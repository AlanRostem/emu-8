package instruction

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

// ldVxByte executes opcode "6xkk", aka "LD Vx, byte"
func ldVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	sys.Registers.GeneralPurpose[x] = mu8.Byte(kk)
}

// ldIAddr executes opcode "Annn", aka "LD I, addr"
func ldIAddr(args []mu8.DByte, sys *system.System) {
	addr := args[0]
	sys.Registers.Index = addr
}
