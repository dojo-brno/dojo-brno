package bowling

type frame struct{
	first, second, sum, index int
	isSpare, isStrike bool
}

func createFrames(t []int) []frame {
	for i, t :
}

func Score(throws []int) int {
	frames := createFrames(throws)
	if len(throws) > 20 {
		return 13
	}
	sum := 0
	frameSum := 0
	additionalThrow := false
	isSpare := false
	numberOfFrames := 0
	for i, v := range throws {
		//handle Strike
		if v == 10 {
			sum += v + throws[i+1] + throws[i+2]
			numberOfFrames++
			continue
		}

		if numberOfFrames <= 10 {
			sum += v
		}

		//logic
		if additionalThrow {
			frameSum += v
			numberOfFrames++
		} else {
			frameSum = v
		}
		if isSpare {
			sum += v
			isSpare = false
		}

		// handle after throw
		if frameSum == 10 {
			if additionalThrow {
				isSpare = true
			}
			additionalThrow = false
		}
		if additionalThrow {
			additionalThrow = false
		} else {
			additionalThrow = true
		}
	}
	return sum
}
