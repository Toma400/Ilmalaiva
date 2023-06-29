package core

import (
    "github.com/hajimehoshi/ebiten/v2"
)

// CAESAR - module for all core game information required
const RES_X = 960
const RES_Y = 720
const TITLE = "Ilmalaiva"

var TEXTURES = map[string]string {
  "bg": "assets/bg/bg_blue.png",
  "pl": "assets/main/player.png",
}

var MODIFIER = map[string]*ebiten.DrawImageOptions {
  "bg": SetOptions(true),
  "pl": SetOptions(true),
}

func SetOptions(scale bool) *ebiten.DrawImageOptions {
    opt := &ebiten.DrawImageOptions{}
    if scale == true { opt.GeoM.Scale(2, 2.3) }
    return opt
}
