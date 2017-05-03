package bowling

type Game [12]Frame

type Frame [2]int

const (
	GameLength = 10
)

func IsStrike(throw int) bool {
	return throw == 10
}

func Score(game Game) int {
	var score int
	for i := 0; i < GameLength; i++ {
		fScore := frameScore(game[i], game[i+1], game[i+2], i)

		score += fScore
	}
	return score
}

func frameScore(frame, next, nextnext Frame, i int) int {
	firstThrow := frame[0]
	frameScore := firstThrow + frame[1]
	// Is it a spare or strike? Add next throw.
	if frameScore == 10 {
		frameScore += next[0]
	}
	// Is it a strike? Add the next throw after the next one.
	if IsStrike(firstThrow) {
		if next[0] == 10 && i != GameLength-1 {
			frameScore += nextnext[0]
		} else {
			frameScore += next[1]
		}
	}
	return frameScore
}
