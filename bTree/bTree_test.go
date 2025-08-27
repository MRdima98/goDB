package btree

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"slices"
	"testing"
)

const maxIndex = 10000
const minIndex = 0
const twoChildren = 16 + 1
const threeChildren = 26
const twoLevelTree = 16*(8+1) + 1

func TestInsertOneIndex(t *testing.T) {
	tree := &bTree{}
	err := tree.insertIndex(5)

	if err != nil || !slices.Contains(tree.indexes[:], 5) {
		t.Errorf(`pushIndex(5) should return 5 instead it returns: %v`, err)
	}
}

func TestNodePreservesOrder(t *testing.T) {
	expected := []int{3, 5, 7}
	tree := &bTree{}
	err := tree.insertIndex(3)
	err = tree.insertIndex(7)
	err = tree.insertIndex(5)

	if err != nil || !slices.Equal(tree.indexes[:], expected) {
		t.Errorf(`Expected order: %v\n got instead: %v`, expected, tree.indexes)
	}
}

func TestOverflowShouldCreateNewNodeIfRoot(t *testing.T) {
	tree, expected := genTree(16)
	tree.insertIndex(minIndex)
	expected = append(expected, minIndex)
	slices.Sort(expected)

	if expected[midKey] != tree.indexes[0] {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected[midKey], tree.indexes[0])
	}

	if !slices.Equal(expected[:midKey], tree.children[0].indexes) {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected[:midKey], tree.children[0].indexes)
	}

	if !slices.Equal(expected[midKey+1:], tree.children[1].indexes) {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected[midKey+1:], tree.children[1].indexes)
	}
}

func TestExtraKeysShouldGoToChildren(t *testing.T) {
	tree, expected := genTree(twoChildren)
	tree.insertIndex(maxIndex + 1)
	tree.insertIndex(minIndex)
	expected = append(expected, maxIndex+1)
	expected = append([]int{minIndex}, expected...)
	expected_root := []int{expected[9]}
	expected_left_son := expected[:9]
	expected_right_son := expected[10:]

	if !slices.Equal(expected_root, tree.indexes) {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected_root, tree.indexes)
	}

	if !slices.Equal(expected_left_son, tree.children[0].indexes) {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected_left_son, tree.children[0].indexes)
	}

	if !slices.Equal(expected_right_son, tree.children[1].indexes) {
		t.Errorf(`Expected array: %v
		 got instead: %v`, expected_right_son, tree.children[1].indexes)
	}

}

func TestSibilingsSplit(t *testing.T) {
	tree, expected := genTree(twoChildren)
	tmp := 0
	for range 9 {
		tmp = maxIndex + rand.Intn(maxIndex)
		tree.insertIndex(tmp)
		expected = append(expected, tmp)
	}
	slices.Sort(expected)
	left_node := expected[:midKey]
	center_node := expected[midKey+1 : (midKey*2 + 1)]
	right_node := expected[len(expected)-midKey:]
	var root []int
	root = append(root, expected[midKey])
	root = append(root, expected[midKey*2+1])

	if !slices.Equal(left_node, tree.children[0].indexes) {
		t.Errorf(`Expected left child: %v
		 got instead: %v`, left_node, tree.children[0].indexes)
	}

	if !slices.Equal(center_node, tree.children[1].indexes) {
		t.Errorf(`Expected center child: %v
		 got instead: %v`, center_node, tree.children[1].indexes)
	}

	if !slices.Equal(right_node, tree.children[2].indexes) {
		t.Errorf(`Expected right child: %v
		 got instead: %v`, right_node, tree.children[2].indexes)
	}

	if !slices.Equal(root, tree.indexes) {
		t.Errorf(`Expected root: %v
		 got instead: %v`, root, tree.indexes)
	}

}

func TestRootSplitByChildrenOverflow(t *testing.T) {
	tree := &bTree{}
	threeLevelTree := 161
	for i := range threeLevelTree {
		tree.insertIndex(i)
	}

	if tree.indexes[0] != threeLevelTree/2 {
		t.Errorf(`Expected root: %v
		 got instead: %v`, threeLevelTree/2, tree.indexes[0])
	}

	if len(tree.children) != 2 {
		t.Errorf(`Expected %v children, 
		 got instead: %v`, 2, len(tree.children))
	}

	expected_children := 8
	left_child := tree.children[0].children
	if len(left_child) != expected_children {
		t.Errorf(`Expected left child to have %v children, 
			got instead: %v! Child-> %v`, expected_children, len(left_child), left_child)
	}

	right_child := tree.children[1].children
	if len(right_child) != expected_children+1 {
		t.Errorf(`Expected right child to have %v children, 
		 instead %v! Child -> %v`, expected_children+1, len(right_child), right_child)
	}
}

func TestRootPrint(t *testing.T) {
	noChildren := 10
	tree, exp := genTree(noChildren)
	got := readTreeStdout(printTree, tree)
	expected := readStdout(exp)
	if got != expected {
		t.Errorf(`I should've printed %v,
			but I printed %v instead `, expected, got)
	}
}

func genTree(nodes_num int) (bTree, []int) {
	tree := &bTree{}
	expected := []int{}
	var tmp int
	for range nodes_num {
		tmp = rand.Intn(maxIndex)
		tree.insertIndex(tmp)
		expected = append(expected, tmp)
	}
	slices.Sort(expected)

	return *tree, expected
}

func readTreeStdout(myPrint func(bTree), tree bTree) string {
	// ????
	noIdea := os.Stdout

	// This connects the buffer with my stuff ???
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
	}

	os.Stdout = w
	myPrint(tree)
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)
	// This flushes the buffer otherwise I will keep eating it
	os.Stdout = noIdea

	return buf.String()
}

func readStdout(arr []int) string {
	noIdea := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
	}

	os.Stdout = w
	fmt.Println(arr, len(arr))
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = noIdea

	return buf.String()
}
