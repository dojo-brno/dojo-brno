package pacman

type Position struct {
	X, Y int
}

type World struct {
	Pacman        Piece
	Width, Height int
	Wall          map[Position]bool
}

type Piece struct {
	X, Y      int
	Direction int
}

func Tick(w World, userInput int) World {
	resWorld := w
	switch w.Pacman.Direction {
	case 0:
		if w.Pacman.Y < w.Height-1 {
			resWorld.Pacman.Y += 1
		}
	case 3:
		if w.Pacman.X < w.Width-1 {
			resWorld.Pacman.X += 1
		}
	case 6:
		if w.Pacman.Y > 0 {
			resWorld.Pacman.Y -= 1
		}
	case 9:
		if w.Pacman.X > 0 {
			resWorld.Pacman.X -= 1
		}
	}

	return resWorld
}
