package cli

import (
	"fmt"

	"github.com/AlanRostem/mu-8/decode"
	"github.com/AlanRostem/mu-8/mu8"
)

func Run() {
	instruction := mu8.DByte(0x8125)
	info := decode.Decode(instruction)
	fmt.Printf("Raw instruction: 0x%04X\n", instruction)
	fmt.Printf("Class: %X\n", info.Class)
	fmt.Printf("Identity: %X\n", info.Idenitity)
	fmt.Printf("Args:\n")
	for _, arg := range info.Args {
		fmt.Printf("\t%X\n", arg)
	}
}
