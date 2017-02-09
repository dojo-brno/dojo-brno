package lca_tree

import "testing"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) LCA(node1, node2 *Node) *Node {
	if node1.Right == node2 {
		return node1
	}
	if node1 != node2 {
		return n
	}

	return node1
}

func TestLCA_Branchy(t *testing.T) {
	root := &Node{
		Value: 1,
		Left:  &Node{Value: 2},
		Right: &Node{
			Value: 3,
			Left:  &Node{Value: 66},
			Right: &Node{Value: -42},
		},
	}
	node1 := root.Right
	node2 := root.Right.Right
	want := root.Right
	got := root.LCA(node1, node2)
	if got != want {
		t.Errorf("LCA(%v, %v) = %+v want %+v", node1, node2, got, want)
	}
}

func TestLCA_42(t *testing.T) {
	root := &Node{
		Value: 1,
		Left:  &Node{Value: 2},
		Right: &Node{
			Value: 3,
			Right: &Node{Value: -42},
		},
	}
	node1 := root.Right
	node2 := root.Right.Right
	want := root.Right
	got := root.LCA(node1, node2)

	if got != want {
		t.Errorf("LCA(%v, %v) = %+v want %+v", node1, node2, got, want)
	}

}

func TestLCATreeWithTwoNodes(t *testing.T) {
	root := &Node{
		Value: 1,
		Left:  &Node{Value: 2},
		Right: &Node{Value: 3},
	}
	node1 := root.Left
	node2 := root.Right
	want := root
	for {
		got := root.LCA(node1, node2)
		if got != want {
			t.Errorf("LCA(%v, %v) = %+v want %+v", node1, node2, got, want)
		}
		break
	}
}

func TestLCA(t *testing.T) {
	var tree *Node = &Node{Value: 1}
	node1 := tree
	node2 := tree
	lca := tree.LCA(node1, node2)
	want := tree
	if lca != want {
		t.Errorf("tree.LCA(%v, %v) = (%v) want (%v)", node1, node2, lca, want)
	}
}

func TestLCATreeWithMultipleNode(t *testing.T) {
	two := &Node{Value: 2}
	var tree *Node = &Node{Value: 1, Left: two}
	node1 := two
	node2 := two
	lca := tree.LCA(node1, node2)
	want := two
	if lca != want {
		t.Errorf("tree.LCA(%v, %v) = %+v want %+v", node1, node2, lca, want)
	}
}
