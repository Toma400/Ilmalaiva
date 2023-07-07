package core

import (
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

// function that takes square-ish element and makes matrix of all pixel positions
// a way to collect all such matrixes into one bigger matrix
// function that checks if currently player isn't touching that position?
