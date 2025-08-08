package btree

import (
	"errors"
	"fmt"
	"slices"
)

const maxKeysPerNode = 16

// TODO: No father, instead try doing iteractevly
type bTree struct {
	indexes  []int
	children []*bTree
	father   *bTree
}

func (r *bTree) insertIndex(index int) error {
	if len(r.indexes) == 0 {
		r.indexes = append(r.indexes, index)
		return nil
	}

	if len(r.indexes) >= maxKeysPerNode && r.father == nil {
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

	if len(r.indexes) >= maxKeysPerNode && r.father != nil {
		r.indexes = append(r.indexes, index)
		slices.Sort(r.indexes)
		lenth := len(r.indexes)
		// left_node := r.indexes[:lenth/2-1]
		// right_node := r.indexes[lenth/2+1:]
		r.father.indexes = append(r.father.indexes, r.indexes[lenth/2])
		fmt.Println(r.father.indexes, &r.father)
		fmt.Println(r.father.indexes, &r.father)

		// slices.Sort(r.father.indexes)
		// r.father.children = append(r.father.children, &bTree{indexes: left_node, father: r.father})
		// r.father.children = append(r.father.children, &bTree{indexes: right_node, father: r.father})
		// var children_in_order [16]*bTree
		// for _, child := range r.father.children {
		// 	for i, index := range r.father.indexes {
		// 		if child.indexes[len(child.indexes)-1] < index {
		// 			children_in_order[i] = child
		// 		}
		// 	}
		// 	if child.indexes[len(child.indexes)-1] >= r.father.indexes[len(r.father.indexes)-1] {
		// 		children_in_order[len(r.father.indexes)] = child
		// 	}
		// }
		// for _, el := range r.father.children {
		// 	fmt.Println(el)
		// }
		// r.father.children = children_in_order[:]
		return nil
	}

	for i, el := range r.indexes {
		if len(r.children) == 0 {
			r.indexes = append(r.indexes, index)
			slices.Sort(r.indexes)
			return nil
		}

		if index <= el && len(r.children) > i {
			r.children[i].insertIndex(index)
			return nil
		}

		if i == len(r.indexes)-1 {
			r.children[i+1].insertIndex(index)
			return nil
		}
	}

	return errors.New("No space")
}

// TODO: I would rather it be in order. However I'm not sure at the moment how
// to approach it
func printTree(root bTree) {
	if root.father == nil {
		fmt.Print("root -> ")
	}
	fmt.Println(root.indexes, len(root.indexes))
	for _, el := range root.children {
		printTree(*el)
	}
}
