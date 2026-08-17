type Point struct{
	x, y int
}
type CountSquares struct {
	ptsCount map[Point]int
	pts []Point
}

func Constructor() CountSquares {
	return CountSquares{
		ptsCount: make(map[Point]int),
		pts: make([]Point, 0),
	}
}

func (this *CountSquares) Add(point []int)  {
   p := Point{point[0], point[1]} 
   this.ptsCount[p]++
   this.pts = append(this.pts, p)
}

func (this *CountSquares) Count(point []int) int {
	px, py := point[0], point[1]
	res := 0

	for _, pt := range this.pts{
		if abs(py-pt.y) != abs(px-pt.x) || px == pt.x || py == pt.y{
			continue
		}
		p1 := Point{pt.x, py}
		p2 := Point{px, pt.y}

		res += this.ptsCount[p1] * this.ptsCount[p2]
	}

	return res
}

func abs(x int)int{
	if x < 0{
		return -x
	}
	return x
}
