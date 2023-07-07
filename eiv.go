package main

import (
		"ilmalaiva/core"
		_ "image/png"
		"log"
		"fmt"

		"github.com/hajimehoshi/ebiten/v2"
		"github.com/hajimehoshi/ebiten/v2/inpututil"
		"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var IMAGES = map[string]*ebiten.Image {}
var GENXY  = core.CollisionBox(core.Coord{29.5*core.TILE, 5.5*core.TILE}, // generator coordinates
											 				 core.Coord{30.5*core.TILE, 6.5*core.TILE})
var WALLXY = core.CollisionBox(core.Coord{14.5*core.TILE, 0.5*core.TILE}, // walls
															core.Coord{15.5*core.TILE, 12.5*core.TILE})
var PLXY   = core.Coord{260, 80}            // player coordinates
var SKXY   = core.Coord{0, 0}               // sky coordinates
var TIME   = 1000

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
		// drawing
		screen.DrawImage(IMAGES["bg"], core.SetOptions(false, SKXY))
		core.DrawTable(screen)
		screen.DrawImage(IMAGES["pl"], core.SetOptions(false, PLXY))
		// set "second batch" which will be walls and all things that are should be
		// above player (door wall for example), simply whatever should cover player
		// (usually is collision-related, not floor-like)

		// controls
		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
				PLXY.X += -1
				if core.Collide(WALLXY, PLXY) {
						PLXY.X += 1
				}
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
				PLXY.X +=  1
				if core.Collide(WALLXY, PLXY) {
						PLXY.X += -1
				}
		}
		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
				PLXY.Y += -1
				if core.Collide(WALLXY, PLXY) {
						PLXY.Y += 1
				}
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
				PLXY.Y +=  1
				if core.Collide(WALLXY, PLXY) {
						PLXY.Y += -1
				}
		}
		if SKXY.X > -2029 { // -2040
				SKXY.X += -1
		} else {
				SKXY.X = -15 // can be dynamic, so FPS related
				// look: https://github.com/hajimehoshi/ebiten/blob/0db860b5dd64948a0907f012ef8418519084e051/internal/clock/clock.go#L55
		}
		if TIME > 0 {
				fmt.Print("Time: ", TIME, "\n")
				TIME += -1
				if core.Collide(GENXY, PLXY) {
						TIME += 500
						fmt.Print("Player collides with generator!\n")
				}
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
