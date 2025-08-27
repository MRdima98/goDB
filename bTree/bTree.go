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

// Inserts indexes and keeps the tree in order
func (r *bTree) insertIndex(index int) error {
	if len(r.children) == 0 {
		if len(r.indexes) < maxKeysPerNode {
			r.indexes = append(r.indexes, index)
			slices.Sort(r.indexes)
		} else {
			*r = *splitRoot(r, index)
		}
		return nil
	}

	for i, el := range r.indexes {
		if index <= el {
			if len(r.children[i].indexes) < maxKeysPerNode {
				r.children[i].insertIndex(index)
			} else {

				splitChild(r, index, i)
			}
			return nil
		}

	}

	last_child := r.children[len(r.children)-1]

	if len(last_child.indexes) < maxKeysPerNode {
		last_child.insertIndex(index)
	} else {
		splitChild(r, index, len(r.children)-1)
	}
	return nil
}

// TODO: I would rather it be in order. However I'm not sure at the moment how
// to approach it
func printTree(root bTree) {
	fmt.Println(root.indexes, len(root.indexes))
}

func printHello() {
	fmt.Println("hello")
}

func splitRoot(root *bTree, index int) *bTree {
	root.indexes = append(root.indexes, index)
	slices.Sort(root.indexes)
	left_child := &bTree{indexes: root.indexes[:midKey]}
	right_child := &bTree{indexes: root.indexes[midKey+1:]}

	if len(root.children) > 0 {
		left_child.children = root.children[:midKey]
		right_child.children = root.children[midKey+1:]
	}

	tree := &bTree{
		indexes:  []int{root.indexes[midKey]},
		children: []*bTree{left_child, right_child},
	}

	return tree
}

func splitChild(node *bTree, index int, i int) {
	overflown_node := node.children[i].indexes
	overflown_node = append(overflown_node, index)
	slices.Sort(overflown_node)
	left_node := overflown_node[:midKey]
	middle := overflown_node[midKey]
	right_node := overflown_node[midKey+1:]
	node.children[i].indexes = left_node
	node.children = append(node.children, &bTree{indexes: right_node})
	var tmp = map[int]*bTree{}
	for _, el := range node.children {
		tmp[el.indexes[len(el.indexes)-1]] = el
	}
	var tmp2 = []int{}
	for k := range tmp {
		tmp2 = append(tmp2, k)
	}
	slices.Sort(tmp2)
	for i, k := range tmp2 {
		node.children[i] = tmp[k]
	}

	if (len(node.indexes)) >= maxKeysPerNode {
		new_tree := splitRoot(node, index)
		*node = *new_tree
	} else {
		node.indexes = append(node.indexes, middle)
		slices.Sort(node.indexes)
	}
}
