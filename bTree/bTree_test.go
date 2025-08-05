package btree

import (
	"math/rand"
	"slices"
	"testing"
)

const maxIndex = 100
const minIndex = 0

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
	tree.insertIndex(maxIndex + 1)
	tree.insertIndex(minIndex)
	expected = append(expected, maxIndex+1)
	expected = append([]int{minIndex}, expected...)
	expected_root := []int{expected[9]}
	expected_left_son := expected[:8]
	expected_right_son := expected[10:]

	if !slices.Equal(expected_root, tree.indexes) {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected_root, tree.indexes)
	}

	if !slices.Equal(expected_left_son, tree.children[0].indexes) {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected_left_son, tree.children[0].indexes)
	}

	if !slices.Equal(expected_right_son, tree.children[1].indexes) {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected_right_son, tree.children[1].indexes)
	}

}

func oneDepthTree(nodes_num int) (bTree, []int) {
	tree := &bTree{}
	expected := []int{}
	var tmp int
	for range nodes_num {
		tmp = rand.Intn(maxIndex)
		tree.insertIndex(tmp)
		expected = append(expected, tmp)
	}
	slices.Sort(expected)

	return *tree, expected
}
