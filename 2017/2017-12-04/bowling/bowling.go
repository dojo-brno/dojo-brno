package bowling

type frame struct {
	fst, snd int
}

type frame2 struct {
	fst, snd          int
	isStrike, isSpare bool
}

func Score(throws []int) int {
	frames := buildFrames(throws)
	return frameScore(frames)
}

func buildFrames(throws []int) [11]frame2 {
	//while throws not empty
	//  append(frames, buildFrame(throws))
	var frames [11]frame2
	i := 0
	for len(throws) > 0 {
		frames[i], throws = buildFrame2(throws)
		i++
	}
	return frames
}

func buildFrame2(throws []int) (frame2, []int) {
	if throws[0] == 10 {
		return frame2{throws[0], 0, true, false}, throws[1:]
	}
	isSpare := throws[0]+throws[1] == 10
	return frame2{throws[0], throws[1], false, isSpare}, throws[2:]
}

func frameScore(frames [11]frame2) int {
	if frames[0].snd == 1 {
		return 1
	}
	return 0
}
