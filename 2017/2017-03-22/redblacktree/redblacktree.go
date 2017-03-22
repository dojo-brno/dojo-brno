package redblacktree

import "github.com/ocdogan/rbt"

type Dict interface {
	Set(int, string)
	Get(int) string
	Delete(int)
}

type ocdogan_rbt struct {
	rbt *rbt.RbTree
}

func (d *ocdogan_rbt) Set(key int, value string) {
	rbtKey := rbt.IntKey(key)
	d.rbt.Insert(&rbtKey, value)
}

func (d *ocdogan_rbt) Get(key int) string {
	rbtKey := rbt.IntKey(key)
	value, _ := d.rbt.Get(&rbtKey)
	ret, _ := value.(string)
	return ret
}

func (d *ocdogan_rbt) Delete(key int) {
	rbtKey := rbt.IntKey(key)
	d.rbt.Delete(&rbtKey)
}

func NewOcdoganTree() Dict {
	return &ocdogan_rbt{rbt.NewRbTree()}
}

func NewGoMapTree() Dict {
	return &gomap{}
}
