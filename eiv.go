package main

import (
		"ilmalaiva/core"
		_ "image/png"
		"log"

		"github.com/hajimehoshi/ebiten/v2"
		"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var IMAGES = map[string]*ebiten.Image {}

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

type Game struct { }

func (g *Game) Update() error {
		return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
		bg_o := &ebiten.DrawImageOptions{}
		bg_o.GeoM.Scale(2, 2.3)
		screen.DrawImage(IMAGES["bg"], bg_o)
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
