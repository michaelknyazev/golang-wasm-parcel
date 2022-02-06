package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 320

	frameWidth  = 64
	frameHeight = 64

	speed = 5
)

var (
	characterImage *ebiten.Image
)

type Game struct {
	stanceY int
	stanceX int

	posY int
	posX int
}

func (g *Game) Update() error {

	_keyMoveUp := inpututil.IsKeyJustPressed(ebiten.KeyW)
	_keyMoveDown := inpututil.IsKeyJustPressed(ebiten.KeyS)
	_keyMoveLeft := inpututil.IsKeyJustPressed(ebiten.KeyA)
	_keyMoveRight := inpututil.IsKeyJustPressed(ebiten.KeyD)

	// Move Control
	if _keyMoveUp {
		g.posY -= speed

		if g.stanceY == 4 && g.stanceX == 0 {
			g.stanceX = 1
		} else {
			g.stanceX = 0
		}

		g.stanceY = 4

		if g.posY <= 0 {
			g.posY = 0
		}
	}

	if _keyMoveDown {
		g.posY += speed

		if g.stanceY == 10 && g.stanceX == 0 {
			g.stanceX = 2
		} else {
			g.stanceX = 0
		}

		g.stanceY = 10

		if g.posY >= screenHeight {
			g.posY = screenHeight
		}
	}

	if _keyMoveLeft {
		g.posX -= speed

		if g.stanceY == 9 && g.stanceX == 0 {
			g.stanceX = 1
		} else {
			g.stanceX = 0
		}

		g.stanceY = 9

		if g.posX <= 0 {
			g.posX = 0
		}
	}

	if _keyMoveRight {
		g.posX += speed

		if g.stanceY == 11 && g.stanceX == 0 {
			g.stanceX = 1
		} else {
			g.stanceX = 0
		}

		g.stanceY = 11

		if g.posX >= screenWidth {
			g.posX = screenWidth
		}
	}
	// End Move Control

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0xf0, 0xf0, 0xf0, 0xf0})

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(g.posX), float64(g.posY))

	sx, sy := g.stanceX*frameWidth, g.stanceY*frameHeight
	frame := characterImage.SubImage(image.Rect(sx, sy, sx+frameHeight, sy+frameWidth)).(*ebiten.Image)

	screen.DrawImage(frame, op)
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

	if err := ebiten.RunGame(&Game{0, 0, 0, 0}); err != nil {
		log.Fatal(err)
	}
}
