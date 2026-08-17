
type CountSquares struct {
	PointMap map[[2]int]int
}

func Constructor() CountSquares {
	return CountSquares{PointMap: make(map[[2]int]int)}
}

func (this *CountSquares) Add(point []int)  {
	this.PointMap[[2]int{point[0], point[1]}]++
}

func (this *CountSquares) Count(point []int) int {
	x, y := point[0], point[1]
	res := 0

	for p, freq := range this.PointMap{
		px, py := p[0], p[1]
		if abs(px - x) == abs(py-y) && x != px && y != py{
			corner1 := this.PointMap[[2]int{x, py}]
			corner2 := this.PointMap[[2]int{px, y}]
			res += freq * corner1 * corner2
		}
	}

	return res
}

func abs(x int)int{
	if x < 0{
		return -x
	}
	return x
}
