package main

var (
	cameraX float64
	cameraY float64
)

func init() {
	cameraX = 320 / 2
	cameraY = 240 / 2
}

func cameraUpdate() {
	cameraX = -p.X + 320/2 - 12
	cameraY = -p.Y + 240/2 - 12
}
