package bowling

import "fmt"

const (
	NumberOfFrames = 10
	maxFrames      = 12
	allPinsDown    = 10
)

type game struct {
	normalGame [NumberOfFrames]frame
	bonusFrame frame
}

type frame struct {
	fst, snd          int
	isStrike, isSpare bool
}

type TooFewThrowsError int

func (e TooFewThrowsError) Error() string {
	return fmt.Sprintf("Too Few Throws: %v", int(e))
}

func isValid(throws []int) error {
	if len(throws) == 0 {
		return TooFewThrowsError(len(throws))
	}
	if len(throws)<20 && throws[0] != 10 {
		return TooFewThrowsError(len(throws))
	}
	return nil
}

func Score(throws []int) (int, error) {
	if err := isValid(throws); err != nil {
		return 0, err
	}
	frames := buildFrames(throws)
	return frameScore(frames), nil
}

func buildFrames(throws []int) [maxFrames]frame {
	var frames [maxFrames]frame
	i := 0
	for len(throws) > 0 {
		frames[i], throws = buildFrame(throws)
		i++
	}
	return frames
}

func buildFrame(throws []int) (frame, []int) {
	if throws[0] == allPinsDown {
		return frame{throws[0], 0, true, false}, throws[1:]
	}
	if len(throws) == 1 {
		return frame{throws[0], 0, false, false}, []int{}
	} else {
		isSpare := throws[0]+throws[1] == allPinsDown
		return frame{throws[0], throws[1], false, isSpare}, throws[2:]
	}
}

func frameScore(frames [maxFrames]frame) int {
	score := 0
	for i := 0; i < NumberOfFrames; i++ {
		score += frames[i].fst + frames[i].snd
		if frames[i].isStrike {
			score += getStrikeBonus(frames, i)
		}
		if frames[i].isSpare {
			score += frames[i+1].fst
		}
	}
	return score
}

func getStrikeBonus(frames [maxFrames]frame, i int) int {
	if frames[i+1].isStrike {
		return frames[i+1].fst + frames[i+2].fst
	} else {
		return frames[i+1].fst + frames[i+1].snd
	}
}
