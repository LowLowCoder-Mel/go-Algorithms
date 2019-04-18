package queue

type ListNode struct {
	v interface{}
	next *ListNode
}

type LinkQueue struct {
	head *ListNode
	tail *ListNode
	length int
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{nil, nil, 0}
}

func (this *LinkQueue) EnQueue(v interface{}) {
	node := &ListNode{v, nil}
	if nil == this.tail {
		this.tail = node
		this.head = node
	} else {
		this.tail.next = node
		this.tail = node
	}

	this.length++
}

func (this *LinkQueue) DeQueue() interface{} {
	if this.head == nil {
		return nil
	}

	v := this.head.v
	this.head = this.head.next
	this.length--
	return v
}
