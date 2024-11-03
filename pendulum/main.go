package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth = 640
	screenHeight = 480
)

type Vector struct {
	X, Y float64
}

type Game struct {
	image image.Image
	base Vector
	length float64
	size float64
	angle float64

	aAcc float64
	aVel float64
}

func (g *Game) Update() error {
	dc := gg.NewContext(screenWidth, screenHeight)


	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(1)

	dc.DrawCircle(g.base.X, g.base.Y, 2)
	dc.Stroke()

	pendulumXLine := g.base.X + (g.length - g.size) * math.Sin(g.angle)
	pendulumYLine := g.base.Y + (g.length - g.size) * math.Cos(g.angle)

	dc.DrawLine(g.base.X, g.base.Y, pendulumXLine, pendulumYLine)
	dc.Stroke()

	pendulumX := g.base.X + g.length * math.Sin(g.angle)
	pendulumY := g.base.Y + g.length * math.Cos(g.angle)

	dc.DrawCircle(pendulumX, pendulumY, g.size)
	dc.Stroke()

	g.image = dc.Image()

	g.aAcc = -.01 * math.Sin(g.angle)
	g.angle += g.aVel
	g.aVel += g.aAcc

	g.aVel *= 0.99

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	screen.DrawImage(ebiten.NewImageFromImage(g.image), nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight 
}


func main() {
	game := &Game{}

	game.base.X = screenWidth / 2
	game.base.Y = 50
	game.angle = math.Pi/4
	game.length = 200
	game.size = 20

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello world!!")

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}
}
