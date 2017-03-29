package dict

type Dict interface {
	Get(int) string
	Set(int, string)
	Delete(int)
}
