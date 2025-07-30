package btree

import (
	"slices"
	"testing"
)

func TestPushOneIndex(t *testing.T) {
	tree := &bTree{}
	err := tree.pushIndex(5)

	if err != nil || !slices.Contains(tree.indexes[:], 5) {
		t.Errorf(`pushIndex(5) should return 5 instead it returns: %v`, err)
	}
}
