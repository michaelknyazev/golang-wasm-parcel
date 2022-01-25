package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 320

	frameWidth  = 65
	frameHeight = 65
)

var (
	characterImage *ebiten.Image
)

type Game struct {
	stanceY int
	stanceX int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0xf0, 0xf0, 0xf0, 0xf0})

	sx, sy := g.stanceX*frameWidth, g.stanceY*frameHeight
	frame := characterImage.SubImage(image.Rect(sx, sy, sx+frameHeight, sy+frameWidth)).(*ebiten.Image)

	screen.DrawImage(frame, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	var err error
	characterImage, _, err = ebitenutil.NewImageFromFile("./assets/sprite.png")

	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{0, 0}); err != nil {
		log.Fatal(err)
	}
}
