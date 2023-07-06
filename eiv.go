package main

import (
		"ilmalaiva/core"
		_ "image/png"
		"log"

		"github.com/hajimehoshi/ebiten/v2"
		"github.com/hajimehoshi/ebiten/v2/inpututil"
		"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var IMAGES = map[string]*ebiten.Image {}
var PLXY   = core.Coord{0, 0}             // player coordinates

func init() {
		for key, value := range core.TEXTURES {
				var err error
				var img *ebiten.Image
				img, _, err = ebitenutil.NewImageFromFile(value)
				if err != nil {
						log.Fatal(err)
				}
				IMAGES[key] = img
		}
}

type Game struct {
		keys []ebiten.Key
}

func (g *Game) Update() error {
		g.keys = inpututil.AppendPressedKeys(g.keys[:0])
		return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
		screen.DrawImage(IMAGES["bg"], core.MODIFIER["bg"])
		screen.DrawImage(IMAGES["pl"], core.SetOptions(true, PLXY))
		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
				PLXY.X += -1
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
				PLXY.X +=  1
		}
		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
				PLXY.Y += -1
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
				PLXY.Y +=  1
		}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
		return core.RES_X, core.RES_Y
}

func main() {
		ebiten.SetWindowSize(core.RES_X, core.RES_Y)
		ebiten.SetWindowTitle(core.TITLE)
		if err := ebiten.RunGame(&Game{}); err != nil {
				log.Fatal(err)
		}
}
