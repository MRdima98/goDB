package btree

import (
	"fmt"
	"slices"
)

const maxKeysPerNode = 16
const midKey = 8

type bTree struct {
	indexes  []int
	children []*bTree
}

func (r *bTree) insertIndex(index int) error {
	if len(r.children) == 0 {
		if len(r.indexes) < maxKeysPerNode {
			r.indexes = append(r.indexes, index)
			slices.Sort(r.indexes)
		} else {
			fmt.Println("overflow")
			*r = *splitRoot(r, index)
			printTree(*r)
		}
		return nil
	}

	for i, el := range r.indexes {
		if index <= el && len(r.children[i].indexes) <= maxKeysPerNode {
			fmt.Println("in")
			r.children[i].insertIndex(index)
			return nil
		} else {
			fmt.Println("out")
			overflown_node := r.children[i].indexes
			overflown_node = append(overflown_node, index)
			ov_len := len(overflown_node)
			left_node := overflown_node[:ov_len/2-1]
			middle := overflown_node[ov_len/2]
			right_node := r.indexes[ov_len/2+1:]
			r.indexes = append(r.indexes, middle)
			slices.Sort(r.indexes)
			r.children[i].indexes = left_node
			r.children = append(r.children, &bTree{indexes: right_node})
			var tmp = map[int]*bTree{}
			for _, el := range r.children {
				tmp[el.indexes[len(el.indexes)-1]] = el
			}
			var tmp2 = []int{}
			for k := range tmp {
				tmp2 = append(tmp2, k)
			}
			slices.Sort(tmp2)
			for i, k := range tmp2 {
				r.children[i] = tmp[k]
			}
		}

	}

	last_child := r.children[len(r.children)-1]

	if len(last_child.indexes) <= maxKeysPerNode {
		fmt.Println("pizza")
		last_child.insertIndex(index)
	} else {
		fmt.Println("hit 2")
		overflown_node := last_child.indexes
		overflown_node = append(overflown_node, index)
		ov_len := len(overflown_node)
		left_node := overflown_node[:ov_len/2-1]
		middle := overflown_node[ov_len/2]
		right_node := r.indexes[ov_len/2+1:]
		r.indexes = append(r.indexes, middle)
		slices.Sort(r.indexes)
		last_child.indexes = left_node
		r.children = append(r.children, &bTree{indexes: right_node})
		var tmp = map[int]*bTree{}
		for _, el := range r.children {
			tmp[el.indexes[len(el.indexes)-1]] = el
		}
		var tmp2 = []int{}
		for k := range tmp {
			tmp2 = append(tmp2, k)
		}
		slices.Sort(tmp2)
		for i, k := range tmp2 {
			r.children[i] = tmp[k]
		}
	}
	return nil
}

// TODO: I would rather it be in order. However I'm not sure at the moment how
// to approach it
func printTree(root bTree) {
	fmt.Println(root.indexes, len(root.indexes))
	for _, el := range root.children {
		printTree(*el)
	}
}

func splitRoot(root *bTree, index int) *bTree {
	root.indexes = append(root.indexes, index)
	slices.Sort(root.indexes)
	left_child := &bTree{indexes: root.indexes[:midKey]}
	right_child := &bTree{indexes: root.indexes[midKey+1:]}

	tree := &bTree{
		indexes:  []int{root.indexes[midKey]},
		children: []*bTree{left_child, right_child},
	}

	return tree
}

// func splitChild(node *bTree, index int) {
// 	node.indexes = append(node.indexes, index)
// 	left_node := node.indexes[:maxKeysPerNode/2-1]
// 	middle := node.indexes[maxKeysPerNode/2]
// 	right_node := node.indexes[maxKeysPerNode/2+1:]
// 	node.indexes = append(node.indexes, middle)
// 	slices.Sort(node.indexes)
// 	node.children[i].indexes = left_node
// 	r.children = append(r.children, &bTree{indexes: right_node})
// 	var tmp = map[int]*bTree{}
// 	for _, el := range r.children {
// 		tmp[el.indexes[len(el.indexes)-1]] = el
// 	}
// 	var tmp2 = []int{}
// 	for k := range tmp {
// 		tmp2 = append(tmp2, k)
// 	}
// 	slices.Sort(tmp2)
// 	for i, k := range tmp2 {
// 		r.children[i] = tmp[k]
// 	}
// }
