package btree

import (
	"errors"
	"slices"
)

const nrKeys = 16

type bTree struct {
	indexes       [nrKeys]int
	children      []*bTree
	indexZeroUsed bool
}

func (r *bTree) insertIndex(index int) error {
	for i, el := range r.indexes {
		if el == 0 {
			r.indexes[i] = index
			arr := r.indexes[:]
			slices.Sort(arr)
			r.indexes = [16]int(arr)
			return nil
		}
	}

	return errors.New("No space")
}
