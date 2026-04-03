package cli

import (
	"github.com/AlanRostem/mu-8/instruction"
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

func Run() {
	opcode := mu8.DByte(0x6F33)
	sys := system.New()
	exec := instruction.NewExecutor(sys)
	exec.Exec(opcode)
}
