package instruction

import (
	"fmt"

	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

// ldVxByte executes opcode "6xkk", aka "LD Vx, byte"
func ldVxByte(args []mu8.DByte, sys *system.System) {
	x := args[0]
	kk := args[1]
	sys.Registers.GeneralPurpose[x] = mu8.Byte(kk)
	fmt.Printf("LD V%d, 0x%02X\n", x, kk)
}

// ldIAddr executes opcode "Annn", aka "LD I, addr"
func ldIAddr(args []mu8.DByte, sys *system.System) {
	addr := args[0]
	sys.Registers.Index = addr
	fmt.Printf("LD I, 0x%03X\n", addr)
}
