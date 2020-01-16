package ant

func Step(board Board, ant Ant) (Board, Ant) {

	//nextAnt := ant

	if ant.direction == 3 {
		return Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, Black, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
			},
			Ant{3, 3, 6}
	}

	return Board{
			{White, White, White, White, White},
			{White, White, White, White, White},
			{White, White, Black, White, White},
			{White, White, White, White, White},
			{White, White, White, White, White},
		},
		Ant{
			3, 2, 3,
		}
}

func AntStep(ant Ant, color Color) (Ant, Color) {
	return Ant{3, 2, 3}, White
}
