package core

import (
    "os"
    "log"
    "bufio"
)

type Coord struct {        // operates on px
    X, Y int
}

func ParseCell(pos int, ctx int, sign string) int {
    var svc = ctx / 100
    return int(float32(pos) * float32(svc))
}

func ReadFLines(path string) [] string {
    f, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    var ret []string

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        ret = append(ret, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        // put error here one day
    }
    return ret
}

func CollisionBox(start, end Coord) [] Coord {
    var ret []Coord
    for i := start.X; i < end.X; i+=1 {     // for i, v := range start.x .. end.x
        for j := start.Y; j < end.Y; j+=1 { // for j, w := range start.y .. end.y
            ret = append(ret, Coord{i, j})
        }
    }
    return ret
}

func MergeCollisionBoxes(box ...[] Coord) [] Coord {
    var ret []Coord
    for ij := range box {
        ret = append(ret, box[ij]...)
    }
    return ret
}

func Collide(rect [] Coord, crd Coord) bool {
    for _, v := range rect {
        if v == crd {
            return true
        }
    }
    return false
}
