package stack

type node struct {
	next *node
	val interface{}
}

type LinkedListStack struct {
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (stack *LinkedListStack) IsEmpty() bool {
	if stack.topNode == nil {
		return true
	}

	return false
}

func (stack *LinkedListStack) Push(v interface{}) {
	stack.topNode = &node{next: stack.topNode, val: v}
}

func (stack *LinkedListStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}

	v := stack.topNode
	stack.topNode = stack.topNode.next
	return v
}

func (stack *LinkedListStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.topNode.val
}

func (stack *LinkedListStack) Flush() {
	stack.topNode = nil
}

