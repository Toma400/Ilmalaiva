package core

import (
    "github.com/hajimehoshi/ebiten/v2"
		"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
type Table struct {
//    Stove      [] Coord
//    Generators [] Coord
    Walls      [] Coord
    PlayerPos  Coord
}

const AIR = '_'          // air tile
const BGT = '░'          // background texture for BGR
const PST = 'P'          // player starting position
var   BGR = []rune{'E',  // tiles that require backgrounds below
                   '@',
                   '#',
                   'P'}
var   CLT = []rune{'║',  // tiles that have collisions
                   '═',
                   '╚',
                   '╝',
                   '╔',
                   '╗',}
//var   STV = '@'          // tile of stove
var MAPTEXTURES = map[rune]string {
    '║': "assets/tiles/wall_vertical.png",
    '═': "assets/tiles/wall_horizontal.png",
    '╚': "assets/tiles/wall_bottom_left.png",
    '╝': "assets/tiles/wall_bottom_right.png",
    '╔': "assets/tiles/wall_top_left.png",
    '╗': "assets/tiles/wall_top_right.png",
    '¤': "assets/tiles/test_tile.png",
    'E': "assets/tiles/wall_door.png",
    '░': "assets/tiles/floor_wood.png",
    '▒': "assets/tiles/floor_carpet.png",
    '@': "assets/tiles/special_stove.png",
    '#': "assets/tiles/special_generator.png",
}
//var   COL   = ReadFLines("core/table/collisions.ilcmp") // collision map
var   MAP   = ReadFLines("core/table/textures.iltmp")   // texture map
const TILE  = 16          // tile std resolution (16px)
var   TABLE = GetTable()  // table
var   TILES = InitTable() // textures to use

func DrawTable(screen *ebiten.Image) {
    x := 0      // coords analysed
    y := 0
    for _, line := range MAP {
        for _, tile := range line {
            if tile != AIR {
                if contains(BGR, tile) {
                    screen.DrawImage(TILES[BGT], SetOptions(0, Coord{x, y}))
                }
                if tile != PST {
                    screen.DrawImage(TILES[tile], SetOptions(0, Coord{x, y}))
                }
            }
            x += TILE
        }
        y += TILE
        x = 0
    }
}

func GetTable() Table {
    //var gen [] Coord
    var wll [] Coord
    var pps Coord

    x := 0    // coords analysed
    y := 0
    for _, line := range MAP {
        for _, tile := range line {
            if contains(CLT, tile) {
                var tpc = CollisionBox(Coord{x-TILE/2, y-TILE},
                                       Coord{x+TILE/2, y+TILE/3})
                wll = MergeCollisionBoxes(wll, tpc)
            } else if tile == PST {
                pps = Coord{x, y}
            }
            x += TILE
        }
        y += TILE
        x = 0
    }

    return Table{ Walls: wll, PlayerPos: pps }
}

func InitTable() map[rune]*ebiten.Image {
    var maptx = map[rune]*ebiten.Image {}
    for key, value := range MAPTEXTURES {
        var err error
        var img *ebiten.Image
        img, _, err = ebitenutil.NewImageFromFile(value)
        if err != nil {
          // put error there one day
        }
        maptx[key] = img
    }
    return maptx
}

func contains(s []rune, str rune) bool {
	for _, v := range s {
  		if v == str {
  		    return true
  		}
	}
	return false
}
