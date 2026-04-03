package instruction

import (
	"github.com/AlanRostem/mu-8/mu8"
	"github.com/AlanRostem/mu-8/system"
)

const Count = 35
const MaxNibbles = 4

type Instruction func(args []mu8.DByte, sys *system.System)
