package collections

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

/*Comparable - A function type used to compare entries.*/
type Comparable func(a interface{}, b interface{}) int

type node struct {
	data    interface{}
	balance int
	right   *node
	left    *node
	parent  *node
}

func (n node) String() string {
	return fmt.Sprintf("%v ", n.data)
}

/*AVL_BST - Represents the AVL Binary Search Tree Structure. */
type AVL_BST struct {
	root    node
	Size    uint
	compare Comparable
}

var root node

func NewBst(comparator Comparable) *AVL_BST {
	newNode := node{nil, 0, nil, nil, nil}
	return &AVL_BST{
		root:    newNode,
		Size:    0,
		compare: comparator,
	}
}

func (bst AVL_BST) Add(element interface{}) (bool, error) {
	if root.data == nil {
		root = node{element, 0, nil, nil, nil}
		return true, nil
	}
	if reflect.TypeOf(element) != reflect.TypeOf(root.data) {
		return false, errors.New("can only insert data of the same types")
	}
	return bst.add(root.parent, &root, element), nil
}

func (bst AVL_BST) add(parent *node, root *node, element interface{}) bool {
	if root == nil {
		root = &node{
			parent:  parent,
			balance: 0,
			data:    element,
			left:    nil,
			right:   nil,
		}
		comp := bst.compare(element, root.parent.data)

		if comp > 0 {
			parent.right = root
		} else if comp < 0 {
			parent.left = root
		}

		rebalanceTree(root)
		return true
	}
	//if the new element is greater than the root
	if bst.compare(element, root.data) > 0 {
		return bst.add(root, root.right, element)
	}

	if bst.compare(element, root.data) < 0 { //if the new element is less than the root
		return bst.add(root, root.left, element)
	}

	return false
}

func rebalanceTree(root *node) {
	if root == nil {
		return
	}
	leftBalance, rightBalance := 0, 0

	if root.left == nil {
		leftBalance = -1
	} else {
		leftBalance = root.left.balance
	}

	if root.right == nil {
		rightBalance = -1
	} else {
		rightBalance = root.right.balance
	}

	root.balance = leftBalance - rightBalance
	if math.Abs(float64(root.balance)) > 1 {
		fmt.Printf("%v is unbalanced \n", root.data)
	}
	rebalanceTree(root.parent)
}

func (bst AVL_BST) AddAll(collection Collection) (bool, error) {

	return false, nil
}

func (bst AVL_BST) Clear() {
	root = node{nil, -1, nil, nil, nil}
}

func (bst AVL_BST) Contains(element interface{}) (bool, error) {
	return false, nil
}

func (bst AVL_BST) Remove(element interface{}) (bool, error) {
	return false, nil
}

func (bst AVL_BST) String() string {
	var result string
	bst.inOrderTraversal(&root, &result)
	return result
}

func (bst AVL_BST) inOrderTraversal(root *node, prefix *string) {
	if root != nil {
		/*fmt.Printf("r = %d (%v) \n", root.data, root.balance)
		fmt.Printf("r.r = %d (%v) \n", root.right.data, root.right.balance)
		fmt.Printf("r.r.r = %d (%v) \n", root.right.right.data, root.right.right.balance)*/

		bst.inOrderTraversal(root.left, prefix)
		*prefix += fmt.Sprintf("%v (%d) \n", root.data, root.balance)
		bst.inOrderTraversal(root.right, prefix)
	}
}
