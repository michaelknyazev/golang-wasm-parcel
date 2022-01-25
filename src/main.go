package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 320
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0xf0, 0xf0, 0xf0, 0xf0})
	fmt.Println(g.count)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
