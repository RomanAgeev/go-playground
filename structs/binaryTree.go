package structs

type BTNode struct {
	Data  interface{}
	Left  *BTNode
	Right *BTNode
}

func NewBTNode(data interface{}) *BTNode {
	return &BTNode{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}
