package btree

import (
	"math/rand"
	"slices"
	"testing"
)

func TestInserOneIndex(t *testing.T) {
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

func TestOverflowShouldCreateNewNode(t *testing.T) {
	var expected []int
	expected = append(expected, 0)
	tree := &bTree{}
	tree.insertIndex(0)
	tmp := 0
	for range 15 {
		tmp += rand.Intn(100)
		expected = append(expected, tmp)
		tree.insertIndex(tmp)
	}
	tree.insertIndex(1)
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
