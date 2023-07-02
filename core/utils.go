package core

import (
    "github.com/hajimehoshi/ebiten/v2"
)

type Coord struct {
    X, Y int
}

func ParseCell(pos int, ctx int, sign string) int {
    var svc = ctx / 100
    return int(float32(pos) * float32(svc))
}

func Contains(ks []ebiten.Key, k ebiten.Key) bool {
	for _, kc := range ks {
		if kc == k {
			return true
		}
	}
	return false
}
