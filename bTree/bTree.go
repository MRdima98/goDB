package btree

import (
	"errors"
	"slices"
)

const maxKeysPerNode = 16

type bTree struct {
	indexes       []int
	children      []*bTree
	indexZeroUsed bool
}

func (r *bTree) insertIndex(index int) error {
	if len(r.indexes) < maxKeysPerNode {
		r.indexes = append(r.indexes, index)
		slices.Sort(r.indexes)
		return nil
	}

	return errors.New("No space")
}
