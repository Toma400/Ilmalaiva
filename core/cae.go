package core

import (
    "math/rand"

    "github.com/hajimehoshi/ebiten/v2"
)

// CAESAR - module for all core game information required
const TR_X  = 60           // tile res
const TR_Y  = 45           // tile res
const RES_X = TR_X * TILE  // 960px
const RES_Y = TR_Y * TILE  // 720px
const TITLE = "Ilmalaiva"

var TEXTURES = map[string]string {
  "bg": "assets/bg/bg_big.png",
  "pl": "assets/main/player.png",
  "fb": "assets/main/fuel_bar.png",
  "ff": "assets/main/fuel_full.png",
  "f9": "assets/main/fuel_9.png",
  "f8": "assets/main/fuel_8.png",
  "f7": "assets/main/fuel_7.png",
  "f6": "assets/main/fuel_6.png",
  "f5": "assets/main/fuel_5.png",
  "f4": "assets/main/fuel_4.png",
  "f3": "assets/main/fuel_3.png",
  "f2": "assets/main/fuel_2.png",
  "f1": "assets/main/fuel_1.png",
  "f0": "assets/main/fuel_0.png",
}

func SetOptions(scale int, movq Coord) *ebiten.DrawImageOptions {
    opt := &ebiten.DrawImageOptions{}
    if scale != 0            { opt.GeoM.Scale(float64(scale), float64(scale)) }
    if !(movq == Coord{0,0}) { opt.GeoM.Translate(float64(movq.X),
                                                  float64(movq.Y)) }
    return opt
}

func RunGenerator(e int, cap int) int {
    if e < cap/2 {
        e += 2
    } else if e < cap {
        e += 1
    }
    return e
}

// deprecated as it is not really passable to work with
func HardcoreGenerator(e int, cap int) int {
    if e < cap/3 {              // < 0.3 energy cap
        e += 2
    } else if e < cap/2 {       // < 0.5 energy cap
        e += 1
    } else if e < cap + cap/4 { // 0.5 .. 1.25 energy cap
        if rand.Intn(100) >= 50 {
            e += 1
        }
    } else { // overpowering
        e = int(e/2)
    }
    return e
}

// function that takes square-ish element and makes matrix of all pixel positions
// a way to collect all such matrixes into one bigger matrix
// function that checks if currently player isn't touching that position?
