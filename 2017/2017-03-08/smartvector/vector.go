package smartvector

type Vector interface {
	New(n int) Vector
	Set(i, value int)
	Get(i int) int
}
