package cli

import (
	"fmt"

	"github.com/AlanRostem/mu-8/mu8"
)

func Run() {
	db := mu8.DByte(0x6ABC)
	for _, nib := range db.GetNibbles() {
		fmt.Printf("0x%X\n", nib)
	}
}
