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

func CollisionBox(start, end Coord) [] Coord {
    var ret []Coord
    for i := start.X; i > end.X; i+=1 { // for i, v := range start.x .. end.x
        for j := start.Y; j > end.Y; j+=1 { // for j, w := range start.y .. end.y
            ret = append(ret, Coord{i, j})
        }
    }
    return ret
}
