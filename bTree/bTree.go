package btree

import (
	"errors"
)

const nrKeys = 16

type bTree struct {
	indexes       [nrKeys]int
	children      []*bTree
	indexZeroUsed bool
}

// TODO: Push an index in the root. Keep track of the index 0 if it's in the tree.
func (r *bTree) pushIndex(index int) error {
	for el, i := range r.indexes {
		if el == 0 {
			r.indexes[i] = index
			return nil
		}
	}

	return errors.New("No space")
}
