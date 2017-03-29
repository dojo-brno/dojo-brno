package dict

import "github.com/ocdogan/rbt"

type ocdoganRbt struct {
	rbt *rbt.RbTree
}

func NewOcdoganRbt() Dict {
	return ocdoganRbt{rbt.NewRbTree()}
}

func (o ocdoganRbt) Get(key int) string {
	k := rbt.IntKey(key)
	value, found := o.rbt.Get(&k)
	if !found {
		return ""
	}
	return value.(string)
}

func (o ocdoganRbt) Set(key int, value string) {
	k := rbt.IntKey(key)
	o.rbt.Insert(&k, value)
}

func (o ocdoganRbt) Delete(key int) {
	k := rbt.IntKey(key)
	o.rbt.Delete(&k)
	return
}
