package main
// COPYRIGHT NOTICE -------------------------------
// Tomasz Stępień 2023 (C), All Rights Reserved
// ------------------------------------------------

import (
		"ilmalaiva/core"
		_ "image/png"
		"strconv"
		"log"
		"fmt"

		"github.com/hajimehoshi/ebiten/v2"
		// "github.com/hajimehoshi/ebiten/v2/text"
		"github.com/hajimehoshi/ebiten/v2/inpututil"
		"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
// TODO:
// - better text not being debug
// - keyboard keys to change
// - option menu in game to change things
// - some readme on how to create maps
// - animated player sprite (?)

var IMAGES = map[string]*ebiten.Image {}
var WALLS  = core.TABLE.Walls               // walls
var STVS   = core.TABLE.Stoves              // stoves
var PLXY   = core.TABLE.PlayerPos           // player coordinates
var GENS   = core.TABLE.Generators          // generators (collisions)
var GENSD  = core.TABLE.GeneratorsD         // generators (data)
var GENSF  = core.InitGeneratorFuel(GENSD)  // generators (fuel)
var SKXY   = core.Coord{0, 0}               // sky coordinates
var FBXY   = core.Coord{900, 450}           // fuel bar coordinates
var FLXY   = core.Coord{904, 478}           // fuel level coordinates
var TIME   = 1000.0
var FUEL   = 0
var BCLD   = 0       // boost cooldown
var BTIME  = 0       // boost time
var GSPD   = 1.0     // game speed (countdown)
var SPD    = 1       // player speed
var CNS    = 0       // consecutive adding
var PTS    = 0

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
	if TIME > 0 {
			// drawing
			screen.DrawImage(IMAGES["bg"], core.SetOptions(0, SKXY))
			// text.Draw(screen, "Fuel ")
			core.DrawTable(screen)
			screen.DrawImage(IMAGES["pl"], core.SetOptions(0, PLXY))
			screen.DrawImage(IMAGES["fb"], core.SetOptions(2, FBXY))
			switch {
				case TIME > 900.0: screen.DrawImage(IMAGES["ff"], core.SetOptions(2, FLXY))
				case TIME > 800.0: screen.DrawImage(IMAGES["f9"], core.SetOptions(2, FLXY))
				case TIME > 700.0: screen.DrawImage(IMAGES["f8"], core.SetOptions(2, FLXY))
				case TIME > 600.0: screen.DrawImage(IMAGES["f7"], core.SetOptions(2, FLXY))
				case TIME > 500.0: screen.DrawImage(IMAGES["f6"], core.SetOptions(2, FLXY))
				case TIME > 400.0: screen.DrawImage(IMAGES["f5"], core.SetOptions(2, FLXY))
				case TIME > 300.0: screen.DrawImage(IMAGES["f4"], core.SetOptions(2, FLXY))
				case TIME > 200.0: screen.DrawImage(IMAGES["f3"], core.SetOptions(2, FLXY))
				case TIME > 100.0: screen.DrawImage(IMAGES["f2"], core.SetOptions(2, FLXY))
				case TIME > 0:   screen.DrawImage(IMAGES["f1"], core.SetOptions(2, FLXY))
				case TIME == 0:  screen.DrawImage(IMAGES["f0"], core.SetOptions(2, FLXY))
			}
			// set "second batch" which will be walls and all things that are should be
			// above player (door wall for example), simply whatever should cover player
			// (usually is collision-related, not floor-like)

			// controls
			if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
					PLXY.X += -SPD
					if core.Collide(WALLS, PLXY) {
							PLXY.X += SPD
					}
			}
			if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
					PLXY.X += SPD
					if core.Collide(WALLS, PLXY) {
							PLXY.X += -SPD
					}
			}
			if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
					PLXY.Y += -SPD
					if core.Collide(WALLS, PLXY) {
							PLXY.Y += SPD
					}
			}
			if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
					PLXY.Y += SPD
					if core.Collide(WALLS, PLXY) {
							PLXY.Y += -SPD
					}
			}
			// BOOST //
			if ebiten.IsKeyPressed(ebiten.KeyQ) {
					if BCLD == 0 {
						var BDQ = 1 // boost downgrade qualifier
						switch {
							case CNS < 2:  BDQ  = 0 // ..1 CNS do not do anything
							case CNS == 2: SPD  = 2;   BTIME += 450; BDQ = CNS
							case CNS == 3: SPD  = 2;   BTIME += 750; BDQ = CNS     // 1.5 + bonus
							case CNS == 4: GSPD = 0.5; BTIME += 500; BDQ = CNS
							case CNS > 4:  GSPD = 0.5; BTIME += 600; BDQ = CNS - 4 // CNS4 + small bonus
							// yes, CNS 6 let you make double boost - slow down game + boost yourself
						}
						CNS += -BDQ
						BCLD = 100
					}
			}
			// SKY MOVING //
			if SKXY.X > -2029 { // -2040
					SKXY.X += -1
			} else {
					SKXY.X = -15 // can be dynamic, so FPS related
					// look: https://github.com/hajimehoshi/ebiten/blob/0db860b5dd64948a0907f012ef8418519084e051/internal/clock/clock.go#L55
			}
			// STOVE //
			if core.Collide(STVS, PLXY) && ebiten.IsKeyPressed(ebiten.KeySpace) && TIME < 1000 {
					if FUEL > 0 {
						CNS  += 1
					}
					TIME += float64(FUEL)
					PTS  += FUEL
					FUEL -= FUEL
			}
			// GENERATORS //
			if core.Collide(GENS, PLXY) && ebiten.IsKeyPressed(ebiten.KeySpace) {
					for _, g := range GENSD {
							if core.Collide(g.Pos, PLXY) {
									FUEL += GENSF[g.Id]
									GENSF[g.Id] = 0
							}
					}
			}

			// GENERAL EVENTS //
			TIME += -GSPD
			for _, g := range GENSD {
					GENSF[g.Id] = core.RunGenerator(GENSF[g.Id], g.Cap)
			}
			ebitenutil.DebugPrint(screen, fmt.Sprintf("Fuel: %s | Boost: %s", strconv.Itoa(FUEL), strconv.Itoa(CNS)))
			// BOOST MANAGEMENTS //
			if BTIME > 1 {
					BTIME -= 1
			} else if BTIME == 1 {
					BTIME = 0
					GSPD = 1.0
					SPD  = 1
			}
			if BCLD > 0 {
					BCLD -= 1
			}
	} else {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("You couldn't keep with airship fuel! Your score: %s", strconv.Itoa(PTS)))
			if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
					for _, g := range GENSD {
							GENSF[g.Id] = 150
					}
					FUEL = 0
					PTS  = 0
					CNS  = 0
					SPD  = 1
					GSPD = 1.0
					TIME = 1000.0
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
