package obstacles

type point struct {
	X, Y int
}

func FindPath(waypoints []point) []point {
	var path []point
	path = append(path, waypoints[0])
	for counter := 1; counter < len(waypoints); counter++ {
		path = findPathBetweenTwoPoints(path, waypoints[counter-1], waypoints[counter])

	}
	return path
}

func findPathBetweenTwoPoints(path []point, start, end point) []point {
	var steps int
	xSteps, xInc := calculateStepsAndIncrement(start.X, end.X)
	ySteps, yInc := calculateStepsAndIncrement(start.Y, end.Y)
	if xSteps > ySteps {
		steps = xSteps
	} else {
		steps = ySteps
	}
	xPoint := start.X
	yPoint := start.Y
	for counter := 1; counter <= steps; counter++ {
		if xPoint != end.X {
			xPoint += xInc
		}
		if yPoint != end.Y {
			yPoint += yInc
		}
		path = append(path, point{xPoint, yPoint})
	}
	return path
}

func calculateStepsAndIncrement(start, end int) (int, int) {
	steps := end - start
	inc := 1
	if steps < 0 {
		inc = -1
		steps = -steps
	}
	return steps, inc
}
