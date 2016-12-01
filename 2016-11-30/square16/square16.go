package square16

type Point struct {
	X, Y int
}
type Solution struct {
	Black, White []Point
}

func IsSolution(s Solution) bool {
	if perimeter(s.White) == 16 && perimeter(s.Black) == 16 {
		return true
	}
	return false
}

func perimeter(p []Point) int {
	if len(p) == 0 {
		return 0
	}
	var r int
	for i := range p[1:] {
		r += distance(p[i+1], p[i])
	}
	r += distance(p[len(p)-1], p[0])
	return r
}

func distance(p1, p2 Point) int {
	return abs(p1.X - p2.X) + abs(p1.Y - p2.Y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
