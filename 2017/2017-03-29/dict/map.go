package dict

type goMap map[int]string

func NewMap() Dict {
	return goMap(map[int]string{})
}

func (m goMap) Get(key int) string {
	return m[key]
}

func (m goMap) Set(key int, value string) {
	m[key] = value
}

func (m goMap) Delete(key int) {
	delete(m, key)
}
