package btree

import (
	"errors"
	"log"
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

	if len(r.indexes) >= maxKeysPerNode {
		r.indexes = append(r.indexes, index)
		slices.Sort(r.indexes)
		lenth := len(r.indexes)
		left_node := r.indexes[:lenth/2-1]
		right_node := r.indexes[lenth/2+1:]
		tree := &bTree{}
		tree.indexes = []int{r.indexes[lenth/2]}
		log.Println(tree.indexes)
		tree.children = append(tree.children, &bTree{indexes: left_node})
		tree.children = append(tree.children, &bTree{indexes: right_node})
		*r = *tree
		return nil
	}

	return errors.New("No space")
}
