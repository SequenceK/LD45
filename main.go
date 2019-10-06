package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	cameraUpdate()
	ebitenutil.DebugPrint(screen, "Hello, World!")
	mapUpdate(screen)
	playerUpdate(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 4, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
