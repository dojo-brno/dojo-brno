package pliskova

import "testing"

func TestPliskova(t *testing.T) {
	tests := []struct {
		FirstPlayerName, SecondPlayerName string
		ballWinner                        string
		score                             Score
	}{
		{
			FirstPlayerName:  "",
			SecondPlayerName: "",
			ballWinner:       "",
			score:            Score{},
		},
		{
			FirstPlayerName:  "Pliskova",
			SecondPlayerName: "Muguruza",
			ballWinner:       "",
			score:            Score{FirstPlayerName: "Pliskova", SecondPlayerName: "Muguruza"},
		},
		{
			FirstPlayerName:  "Pliskova",
			SecondPlayerName: "Muguruza",
			ballWinner:       "F",
			score:            Score{FirstPlayerName: "Pliskova", SecondPlayerName: "Muguruza", FirstPlayerBalls: "Fifteen", SecondPlayerBalls: "Love"},
		},
		{
			FirstPlayerName:  "Pliskova",
			SecondPlayerName: "Muguruza",
			ballWinner:       "FS",
			score:            Score{FirstPlayerName: "Pliskova", SecondPlayerName: "Muguruza", FirstPlayerBalls: "Fifteen", SecondPlayerBalls: "Fifteen"},
		},
	}
	for _, tt := range tests {
		if got := CalculateScore(tt.FirstPlayerName, tt.SecondPlayerName, tt.ballWinner); got != tt.score {
			t.Errorf("CalculateScore(%q, %q, %q):%v WANT: %v", tt.FirstPlayerName, tt.SecondPlayerName, tt.ballWinner, got, tt.score)
		}
	}
}
