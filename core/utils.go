package core

import (
    "os"
    "bufio"
    "gopkg.in/yaml.v3"
)

type Coord struct {        // operates on px
    X, Y int
}
type Config struct {       // config file
    // KEYS struct {
    //     MoveLeft  string `yaml:"move_left"`
    //     MoveRight string `yaml:"move_right"`
    //     MoveUp    string `yaml:"move_up"`
    //     MoveDown  string `yaml:"move_down"`
    //     Boost     string
    //     Activate  string
    // } `yaml:"KEYS"`
    MAPS struct {
        Background string
        Map        string
    } `yaml:"MAPS"`
}
var CFG = ReadConfig()

func ParseCell(pos int, ctx int, sign string) int {
    var svc = ctx / 100
    return int(float32(pos) * float32(svc))
}

func ReadFLines(path string) [] string {
    f, err := os.Open(path)
    if err != nil {
        // put error here one day
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

func MergeGenerators(gen ...[] Generator) [] Generator {
    var ret []Generator
    for ij := range gen {
        ret = append(ret, gen[ij]...)
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

func ReadConfig() Config {
    f, err := os.ReadFile("ilmalaiva.yaml")
    if err != nil {
        // put error here one day
    }
    ret := Config{}

    err = yaml.Unmarshal([]byte(string(f)), &ret)
    if err != nil {
        // put error here one day
    }
    return ret
}
