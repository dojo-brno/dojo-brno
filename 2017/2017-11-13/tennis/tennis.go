package pliskova

type Score struct {
	FirstPlayerName, SecondPlayerName   string
	FirstPlayerBalls, SecondPlayerBalls string
	FirstPlayerGames, SecondPlayerGames int
	FirstPlayerSets, SecondPlayerSets   [3]int
}

func (s Score) withFirstPlayerBalls(firstPlayerBalls string) Score {
	s.FirstPlayerBalls = firstPlayerBalls
	return s
}

const (
	firstPlayerAbbrev  = "F"
	secondPlayerAbbrev = "S"
)

func CalculateScore(firstPlayerName, secondPlayerName, ballWinnerSequence string) Score {
	if ballWinnerSequence == "FS" {
		return Score{FirstPlayerName: firstPlayerName, SecondPlayerName: secondPlayerName, FirstPlayerBalls: "Fifteen", SecondPlayerBalls: "Fifteen"}
	}
	if ballWinnerSequence == firstPlayerAbbrev {
		return Score{FirstPlayerName: firstPlayerName, SecondPlayerName: secondPlayerName, FirstPlayerBalls: "Fifteen", SecondPlayerBalls: "Love"}
	}
	return Score{FirstPlayerName: firstPlayerName, SecondPlayerName: secondPlayerName}
}

type Game struct {
	firstPlayerBalls, secondPlayerBalls int
}

func (g *Game) PlayBall(winner string) Game {
	if winner == "F" {
		g.firstPlayerBalls++
	}

	if winner == "S" {
		g.secondPlayerBalls++
	}

	//fmt.error("Should not ever happen");
	return *g
}

func (g *Game) PlayerString(player int) string {
	var playerBalls int
	if player == 1 {
		playerBalls = g.firstPlayerBalls
	}
	if player == 2 {
		playerBalls = g.secondPlayerBalls
	}

		switch player {
		case 0:
			return "Love"
		case 1:
			return "Fifteen"
		case 2:
			return "Thirty"
		case 3:
			return "Forty"

		}
	}
	retrun ""
}
