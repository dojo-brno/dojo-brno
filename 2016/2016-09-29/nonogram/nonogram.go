package nonogram

type Problem struct {
	Vertical   [][]int
	Horizontal [][]int
}

type Solution [][]bool

func isSolution(p Problem, s Solution) bool {
	out := p.Vertical[0][0] * p.Horizontal[0][0]
        if s[0][0] && !s[0][0] &&  
	return (out == 0) == !s[0][0]
}
