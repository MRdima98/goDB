package btree

import (
	"math/rand"
	"slices"
	"testing"
)

const biggestIndex = 100

func TestInsertOneIndex(t *testing.T) {
	tree := &bTree{}
	err := tree.insertIndex(5)

	if err != nil || !slices.Contains(tree.indexes[:], 5) {
		t.Errorf(`pushIndex(5) should return 5 instead it returns: %v`, err)
	}
}

func TestNodePreservesOrder(t *testing.T) {
	expected := []int{3, 5, 7}
	tree := &bTree{}
	err := tree.insertIndex(3)
	err = tree.insertIndex(7)
	err = tree.insertIndex(5)

	if err != nil || !slices.Equal(tree.indexes[:], expected) {
		t.Errorf(`Expected order: %v\n got instead: %v`, expected, tree.indexes)
	}
}

func TestOverflowShouldCreateNewNodeIfRoot(t *testing.T) {
	tree, expected := oneDepthTree(15)
	tree.insertIndex(0)
	tree.insertIndex(1)
	expected = append(expected, 0)
	expected = append(expected, 1)
	slices.Sort(expected)
	middle := 17 / 2

	if expected[middle] != tree.indexes[0] {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected[middle], tree.indexes[0])
	}

	if !slices.Equal(expected[:middle-1], tree.children[0].indexes) {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected[:middle], tree.children[0].indexes)
	}

	if !slices.Equal(expected[middle+1:], tree.children[1].indexes) {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected[middle+1:], tree.children[1].indexes)
	}
}

func TestExtraKeysShouldGoToChildren(t *testing.T) {
	tree, expected := oneDepthTree(17)
	tree.insertIndex(biggestIndex + 1)
	expected = append(expected, biggestIndex+1)
	root := []int{expected[8]}
	// left_son := expected[:7]
	// middle_son := expected[9:16]
	// right_son := expected[18:]
	// t.Log(root, left_son, middle_son, right_son)
	t.Log(tree.indexes)
	t.Log(tree.children[0])
	t.Log(tree.children[1])

	if !slices.Equal(root, tree.indexes) {
		t.Errorf(`Expected array: %v
		 got instead: %v`, root, tree.indexes)
	}
}

func oneDepthTree(nodes_num int) (bTree, []int) {
	tree := &bTree{}
	expected := []int{}
	var tmp int
	for range nodes_num {
		tmp = rand.Intn(100)
		tree.insertIndex(tmp)
		expected = append(expected, tmp)
	}
	slices.Sort(expected)

	return *tree, expected
}
