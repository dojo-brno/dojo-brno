package pacman

import "testing"

func TestPacmanMoving(t *testing.T) {
	tests := []struct {
		testName  string
		world     World
		userInput int
		want      World
	}{
		{
			testName: "test move Pacman right",
			world: World{
				Pacman: Piece{
					X:         1,
					Y:         1,
					Direction: 3,
				},
				Width:  4,
				Height: 4,
			},
			userInput: -1,
			want: World{
				Pacman: Piece{
					X:         2,
					Y:         1,
					Direction: 3,
				},
				Width:  4,
				Height: 4,
			},
		},
		{
			testName: "test move Pacman right #2",
			world: World{
				Pacman: Piece{
					X:         2,
					Y:         2,
					Direction: 3,
				},
				Width:  4,
				Height: 4,
			},
			userInput: -1,
			want: World{
				Pacman: Piece{
					X:         3,
					Y:         2,
					Direction: 3,
				},
				Width:  4,
				Height: 4,
			},
		},
		{
			testName: "test move Pacman down",
			world: World{
				Pacman: Piece{
					X:         2,
					Y:         2,
					Direction: 6,
				},
				Width:  4,
				Height: 4,
			},
			userInput: -1,
			want: World{
				Pacman: Piece{
					X:         2,
					Y:         1,
					Direction: 6,
				},
				Width:  4,
				Height: 4,
			},
		},
		{
			testName: "test move Pacman left",
			world: World{
				Pacman: Piece{
					X:         1,
					Y:         1,
					Direction: 9,
				},
				Width:  4,
				Height: 4,
			},
			userInput: -1,
			want: World{
				Pacman: Piece{
					X:         0,
					Y:         1,
					Direction: 9,
				},
				Width:  4,
				Height: 4,
			},
		},
		{
			testName: "test move Pacman up",
			world: World{
				Pacman: Piece{
					X:         1,
					Y:         1,
					Direction: 0,
				},
				Width:  4,
				Height: 4,
			},
			userInput: -1,
			want: World{
				Pacman: Piece{
					X:         1,
					Y:         2,
					Direction: 0,
				},
				Width:  4,
				Height: 4,
			},
		},
		{
			testName: "test move Pacman hits into southern border",
			world: World{
				Pacman: Piece{
					X:         1,
					Y:         0,
					Direction: 6,
				},
				Width:  4,
				Height: 4,
			},
			userInput: -1,
			want: World{
				Pacman: Piece{
					X:         1,
					Y:         0,
					Direction: 6,
				},
				Width:  4,
				Height: 4,
			},
		},
		{
			testName: "test move Pacman hits into northen border",
			world: World{
				Pacman: Piece{
					X:         3,
					Y:         3,
					Direction: 0,
				},
				Width:  4,
				Height: 4,
			},
			userInput: -1,
			want: World{
				Pacman: Piece{
					X:         3,
					Y:         3,
					Direction: 0,
				},
				Width:  4,
				Height: 4,
			},
		},
		{
			testName: "test move Pacman hits into western border",
			world: World{
				Pacman: Piece{
					X:         0,
					Y:         3,
					Direction: 9,
				},
				Width:  4,
				Height: 4,
			},
			userInput: -1,
			want: World{
				Pacman: Piece{
					X:         0,
					Y:         3,
					Direction: 9,
				},
				Width:  4,
				Height: 4,
			},
		}, {
			testName: "test move Pacman hits into eastern border",
			world: World{
				Pacman: Piece{
					X:         3,
					Y:         3,
					Direction: 3,
				},
				Width:  4,
				Height: 4,
			},
			userInput: -1,
			want: World{
				Pacman: Piece{
					X:         3,
					Y:         3,
					Direction: 3,
				},
				Width:  4,
				Height: 4,
			},
		}, {
			testName: "Test move Pacman hits into wall:",
			world: World{
				Pacman: Piece{
					X:         3,
					Y:         3,
					Direction: 3,
				},
				Width:  4,
				Height: 4,
			},
			userInput: -1,
			want: World{
				Pacman: Piece{
					X:         3,
					Y:         3,
					Direction: 3,
				},
				Width:  4,
				Height: 4,
			},
		},
	}

	for _, tt := range tests {
		if got := Tick(tt.world, tt.userInput); got != tt.want {
			t.Errorf("%q\n Tick(%v, %v): %v, but we WANT: %v", tt.testName, tt.world, tt.userInput, got, tt.want)
		}
	}
}
