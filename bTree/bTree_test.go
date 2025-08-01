package btree

import (
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
	base := [16]int{3, 5, 7}
	expected := base[:]
	slices.Sort(expected)
	tree := &bTree{}
	err := tree.insertIndex(3)
	err = tree.insertIndex(7)
	err = tree.insertIndex(5)

	if err != nil || !slices.Equal(tree.indexes[:], expected) {
		t.Errorf(`Expected order: %v\n got instead: %v`, expected, tree.indexes)
	}
}
