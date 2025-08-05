package btree

import (
	"errors"
	"fmt"
	"log"
	"slices"
)

const maxKeysPerNode = 16

type bTree struct {
	indexes  []int
	children []*bTree
	father   *bTree
}

func (r *bTree) insertIndex(index int) error {
	last_elem := len(r.indexes) - 1
	childrens := len(r.children) > 0
	indexes := len(r.indexes) > 0
	if childrens && indexes {
		if index > r.indexes[last_elem] {
			r.children[last_elem+1].insertIndex(index)
			return nil
		}

		if index < r.indexes[0] {
			r.children[0].insertIndex(index)
			return nil
		}
	}

	if len(r.indexes) < maxKeysPerNode {
		r.indexes = append(r.indexes, index)
		slices.Sort(r.indexes)
		return nil
	}

	if len(r.indexes) >= maxKeysPerNode {
		log.Println("split")
		r.indexes = append(r.indexes, index)
		slices.Sort(r.indexes)
		lenth := len(r.indexes)
		left_node := r.indexes[:lenth/2-1]
		right_node := r.indexes[lenth/2+1:]
		tree := &bTree{}
		tree.indexes = []int{r.indexes[lenth/2]}
		tree.children = append(tree.children, &bTree{indexes: left_node, father: tree})
		tree.children = append(tree.children, &bTree{indexes: right_node, father: tree})
		*r = *tree
		return nil
	}

	return errors.New("No space")
}

// TODO: I would rather it be in order. However I'm not sure at the moment how
// to approach it
func printTree(root bTree) {
	for _, el := range root.children {
		printTree(*el)
	}
	fmt.Println(root.indexes, len(root.indexes))
}
