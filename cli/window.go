package cli

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct{}

func NewWindow() *Window {
	return &Window{}
}

func (w *Window) Run() {
	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("MU-8")
	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}
}

func (w *Window) Update() error {
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	screen.Set(5, 10, color.Black)
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 32
}
