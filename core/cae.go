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
}

var MODIFIER = map[string]*ebiten.DrawImageOptions {
  "bg": SetOptions(true, Coord{0,0}),
  // "pl": SetOptions(true),
}

func SetOptions(scale bool, movq Coord) *ebiten.DrawImageOptions {
    opt := &ebiten.DrawImageOptions{}
    if scale == true         { opt.GeoM.Scale(2, 2.3)              }
    if !(movq == Coord{0,0}) { opt.GeoM.Translate(float64(movq.X),
                                                  float64(movq.Y)) }
    return opt
}

// function that takes square-ish element and makes matrix of all pixel positions
// a way to collect all such matrixes into one bigger matrix
// function that checks if currently player isn't touching that position?
