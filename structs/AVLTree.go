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

	var heightMax int
	if heightL > heightR {
		heightMax = heightL
	} else {
		heightMax = heightR
	}
	node.height = heightMax + 1

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
