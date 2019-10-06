package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type player struct {
	Image         *ebiten.Image
	ImageOpts     *ebiten.DrawImageOptions
	MovementSpeed float64
	X             float64
	Y             float64

	FrameCounter     uint
	Frames           []*ebiten.Image
	FrameIndex       uint
	AnimIndex        uint
	CurrentAnimation []uint

	idAnim []uint
	mdAnim []uint
	iuAnim []uint
	muAnim []uint
	ilAnim []uint
	mlAnim []uint
	irAnim []uint
	mrAnim []uint
}

var p *player

func init() {
	p = &player{}
	var err error
	p.Image, _, err = ebitenutil.NewImageFromFile("player.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	p.ImageOpts = &ebiten.DrawImageOptions{}

	w, h := p.Image.Size()
	cols := w / 24
	rows := h / 24
	p.Frames = make([]*ebiten.Image, cols*rows)
	for i := range p.Frames {
		x := (i % (w / 24)) * 24
		y := (i / (w / 24)) * 24

		p.Frames[i] = p.Image.SubImage(image.Rect(x, y, x+24, y+24)).(*ebiten.Image)
	}

	p.idAnim = []uint{0}
	p.mdAnim = []uint{1, 0, 2, 0}
	p.iuAnim = []uint{3}
	p.muAnim = []uint{4, 3, 5, 3}
	p.ilAnim = []uint{6}
	p.mlAnim = []uint{7, 6, 8, 6}
	p.irAnim = []uint{9}
	p.mrAnim = []uint{10, 9, 11, 9}

	p.CurrentAnimation = p.idAnim
	p.MovementSpeed = 2
}

var vx float64
var vy float64

func playerUpdate(screen *ebiten.Image) {

	if vx > 0 {
		p.CurrentAnimation = p.irAnim
	} else if vx < 0 {
		p.CurrentAnimation = p.ilAnim
	} else if vy > 0 {
		p.CurrentAnimation = p.idAnim
	} else if vy < 0 {
		p.CurrentAnimation = p.iuAnim
	}

	vx = 0
	vy = 0

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		vx = p.MovementSpeed
		vy = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		vx = -p.MovementSpeed
		vy = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		vy = -p.MovementSpeed
		vx = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		vy = p.MovementSpeed
		vx = 0
	}
	if vx > 0 {
		p.CurrentAnimation = p.mrAnim
	} else if vx < 0 {
		p.CurrentAnimation = p.mlAnim
	} else if vy > 0 {
		p.CurrentAnimation = p.mdAnim
	} else if vy < 0 {
		p.CurrentAnimation = p.muAnim
	}

	p.X += vx
	p.Y += vy

	p.ImageOpts.GeoM.Reset()
	p.ImageOpts.GeoM.Translate(p.X, p.Y)
	p.ImageOpts.GeoM.Translate(cameraX, cameraY)
	if p.FrameCounter%uint(48/len(p.CurrentAnimation)) == 0 {
		p.AnimIndex++
	}
	p.AnimIndex = uint(int(p.AnimIndex) % len(p.CurrentAnimation))
	p.FrameIndex = p.CurrentAnimation[p.AnimIndex]
	screen.DrawImage(p.Frames[p.FrameIndex], p.ImageOpts)
	p.FrameCounter++
}
