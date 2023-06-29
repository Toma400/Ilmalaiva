package core

func parseCell(pos int, ctx int, sign string) int {
    var svc = ctx / 100
    return int(float32(pos) * float32(svc))
}
