package structs

import (
	"github.com/RomanAgeev/playground/utils"
)

type AVLNode struct {
	Key    int
	Left   *AVLNode
	Right  *AVLNode
	height int
}

func NewAVLNode(key int) *AVLNode {
	return &AVLNode{
		Key:    key,
		Left:   nil,
		Right:  nil,
		height: 1,
	}
}

func (node *AVLNode) InsertAVL(key int) *AVLNode {
	if node == nil {
		return NewAVLNode(key)
	}

	if key < node.Key {
		node.Left = node.Left.InsertAVL(key)
	} else if key > node.Key {
		node.Right = node.Right.InsertAVL(key)
	} else {
		return node
	}

	heightL := node.Left.getHeight()
	heightR := node.Right.getHeight()

	node.height = utils.MaxInt(heightL, heightR) + 1

	switch balance := node.getBalance(); {

	case balance > 1 && key < node.Left.Key:
		return rightRotate(node)

	case balance < -1 && key > node.Right.Key:
		return leftRotate(node)

	case balance > 1 && key > node.Left.Key:
		node.Left = leftRotate(node.Left)
		return rightRotate(node)

	case balance < -1 && key < node.Right.Key:
		node.Right = rightRotate(node.Right)
		return leftRotate(node)

	default:
		return node
	}
}

func (node *AVLNode) DeleteAVL(key int) *AVLNode {
	if node == nil {
		return node
	}

	if key < node.Key {
		node.Left = node.Left.DeleteAVL(key)
	} else if key > node.Key {
		node.Right = node.Right.DeleteAVL(key)
	} else {
		if node.Left != nil && node.Right != nil {
			node.Left = node.replaceMinMax(node.Left)
		} else if node.Left != nil {
			node = node.Left
		} else if node.Right != nil {
			node = node.Right
		} else {
			return nil
		}
	}

	heightL := node.Left.getHeight()
	heightR := node.Right.getHeight()

	node.height = utils.MaxInt(heightL, heightR) + 1

	switch balance := node.getBalance(); {

	case balance > 1 && node.Left.getBalance() >= 0:
		return rightRotate(node)

	case balance > 1 && node.Left.getBalance() < 0:
		node.Left = leftRotate(node.Left)
		return rightRotate(node)

	case balance < -1 && node.Right.getBalance() <= 0:
		return leftRotate(node)

	case balance < -1 && node.Right.getBalance() > 0:
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

func (root *AVLNode) replaceMinMax(node *AVLNode) *AVLNode {
	if node.Right != nil {
		node.Right = root.replaceMinMax(node.Right)
		return node
	}

	if node.Left != nil {
		node.Left = root.replaceMinMax(node.Left)
		return node
	}

	root.Key = node.Key

	return nil
}

func leftRotate(node *AVLNode) *AVLNode {
	right := node.Right
	temp := right.Left

	right.Left = node
	node.Right = temp

	node.updateHeight()
	right.updateHeight()

	return right
}

func rightRotate(node *AVLNode) *AVLNode {
	left := node.Left
	temp := left.Right

	left.Right = node
	node.Left = temp

	node.updateHeight()
	left.updateHeight()

	return left
}

func (node *AVLNode) updateHeight() {
	if node == nil {
		return
	}

	node.height = utils.MaxInt(node.Left.getHeight(), node.Right.getHeight()) + 1
}

func (node *AVLNode) getHeight() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *AVLNode) getBalance() int {
	if node == nil {
		return 0
	}
	return node.Left.getHeight() - node.Right.getHeight()
}
