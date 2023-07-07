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
// - ogarnij napis na końcu (koniec gry), w razie czego też podpowiedź z klawiszami
// - napis nie powinien być debugiem

// - player pos on tilemap
// - config file with keyboard...
// - ...and pickable levels??? :O

var IMAGES = map[string]*ebiten.Image {}
// STOVE //
var STVXY  = core.CollisionBox(core.Coord{29.5*core.TILE,   5.5*core.TILE},
											 				 core.Coord{30.5*core.TILE,   6.5*core.TILE})
// GENERATORS //
var GEN1XY = core.CollisionBox(core.Coord{15.5*core.TILE,   7.5*core.TILE},
											 				 core.Coord{16.5*core.TILE,   8.5*core.TILE})
var GEN1E  = 150 // energy stored
var GEN1EC = 150 // energy cap
var GEN2XY = core.CollisionBox(core.Coord{15.5*core.TILE,   9.5*core.TILE},
											 				 core.Coord{16.5*core.TILE,  10.5*core.TILE})
var GEN2E  = 150 // energy stored
var GEN2EC = 150 // energy cap
var GEN3XY = core.CollisionBox(core.Coord{29.5*core.TILE,  10.5*core.TILE},
 											 				 core.Coord{30.5*core.TILE,  11.5*core.TILE})
var GEN3E  = 150 // energy stored
var GEN3EC = 150 // energy cap
// WALLS //
var WLEFT  = core.CollisionBox(core.Coord{14.5*core.TILE,   0.5*core.TILE},
															core.Coord{15.5*core.TILE,   12.5*core.TILE})
var WRIGHT = core.CollisionBox(core.Coord{30.5*core.TILE,   0.5*core.TILE},
															core.Coord{31.5*core.TILE,   12.5*core.TILE})
var WUP    = core.CollisionBox(core.Coord{14.5*core.TILE,   0.5*core.TILE},
															core.Coord{31.5*core.TILE,    1.5*core.TILE})
var WDOWN  = core.CollisionBox(core.Coord{14.5*core.TILE,  11.0*core.TILE},
															core.Coord{31.5*core.TILE,   12.5*core.TILE})
var WMIDL  = core.CollisionBox(core.Coord{14.5*core.TILE,   8.0*core.TILE},
															core.Coord{22.75*core.TILE,   9.5*core.TILE})
var WMIDR  = core.CollisionBox(core.Coord{23.25*core.TILE,  8.0*core.TILE},
															core.Coord{31.5*core.TILE,    9.5*core.TILE})
var WALLXY = core.MergeCollisionBoxes(WLEFT, WRIGHT, WUP, WDOWN, WMIDL, WMIDR)
// OTHER DATA //
var PLXY   = core.Coord{260, 80}            // player coordinates
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
					if core.Collide(WALLXY, PLXY) {
							PLXY.X += SPD
					}
			}
			if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
					PLXY.X += SPD
					if core.Collide(WALLXY, PLXY) {
							PLXY.X += -SPD
					}
			}
			if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
					PLXY.Y += -SPD
					if core.Collide(WALLXY, PLXY) {
							PLXY.Y += SPD
					}
			}
			if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
					PLXY.Y += SPD
					if core.Collide(WALLXY, PLXY) {
							PLXY.Y += -SPD
					}
			}
			// BOOST //
			if ebiten.IsKeyPressed(ebiten.KeyQ) {
					if BCLD == 0 {
						var BDQ = 1 // boost downgrade qualifier
						switch {
							case CNS < 2:  BDQ  = 0 // ..1 CNS do not do anything
							case CNS == 2: SPD  = 2;   BTIME = 400; BDQ = CNS
							case CNS == 3: SPD  = 2;   BTIME = 600; BDQ = CNS
							case CNS == 4: GSPD = 0.5; BTIME = 300; BDQ = CNS
							case CNS > 4:  GSPD = 0.5; BTIME = 400; BDQ = CNS - 4
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
			if core.Collide(STVXY, PLXY) && ebiten.IsKeyPressed(ebiten.KeySpace) && TIME < 1000 {
					if FUEL > 0 {
						CNS  += 1
					}
					TIME += float64(FUEL)
					PTS  += FUEL
					FUEL -= FUEL
			}
			// GENERATORS //
			if core.Collide(GEN1XY, PLXY) && ebiten.IsKeyPressed(ebiten.KeySpace) {
					FUEL += GEN1E
					GEN1E = 0
			}
			if core.Collide(GEN2XY, PLXY) && ebiten.IsKeyPressed(ebiten.KeySpace) {
					FUEL += GEN2E
					GEN2E = 0
			}
			if core.Collide(GEN3XY, PLXY) && ebiten.IsKeyPressed(ebiten.KeySpace) {
					FUEL += GEN3E
					GEN3E = 0
			}

			// GENERAL EVENTS //
			TIME += -GSPD
			GEN1E = core.Generator(GEN1E, GEN1EC)
			GEN2E = core.Generator(GEN2E, GEN2EC)
			GEN3E = core.Generator(GEN3E, GEN3EC)
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
					GEN1E = 150
					GEN2E = 150
					GEN3E = 150
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
