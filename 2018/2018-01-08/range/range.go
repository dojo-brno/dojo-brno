package dojorange

type Interval struct {
	First, Last     int
	FirstIn, LastIn bool
}

func (this Interval) Contains(n int) bool {
	if n > this.Last || (n == this.Last && !this.LastIn) {
		return false
	}
	if n < this.First || (n == this.First && !this.FirstIn) {
		return false
	}
	return true
}

func (this Interval) GetAllPoints() []int {
	if FirstIn{
		return []int{0}
	}
	return nil
}
