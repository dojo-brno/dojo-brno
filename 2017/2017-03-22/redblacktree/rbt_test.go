package redblacktree

import "testing"

func Test(t *testing.T) {
	implementations := []func() Dict{
		NewOcdoganTree,
		NewGoMapTree,
	}

	for _, impl := range implementations {
		tree := impl()
		tree.Set(1, "foo")
		if got := tree.Get(1); got != "foo" {
			t.Errorf("tree.Get(1) = %v, want %v", got, "foo")
		}
		tree.Delete(1)
		if got := tree.Get(1); got != "" {
			t.Errorf("tree.Get(1) = %v, want %v", got, "")
		}
	}
}
