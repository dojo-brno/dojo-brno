package gameoflife

type Cell struct {
	state     State
	neighbors []*Cell
}

type State bool

const (
	Dead State = false
	Live State = true
)

func NextGen(cell *Cell) *Cell {
	nextGen := &Cell{Dead, nil}
	for range cell.neighbors {
		nextGen.neighbors = append(nextGen.neighbors, &Cell{Dead, []*Cell{nextGen}})
	}
	return nextGen
}
