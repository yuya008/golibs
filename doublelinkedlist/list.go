package doublelinkedlist

import (
	_ "log"
)

const (
	After    = 1
	Before   = 2
	Forward  = 1
	Backward = 2
)

type Compartor func(*Node, *Node) int
type EqualFunc func(*Node, *Node) bool

type List interface {
	// 往链表左侧添加元素
	Lput(interface{})
	// 往链表右侧添加元素
	Rput(interface{})
	// 获取指定索引位置的元素
	Get(int) interface{}
	// 从链表左侧弹出元素
	Lpop() interface{}
	// 从链表右侧弹出元素
	Rpop() interface{}
	// 快捷放入元素(右)
	Put(interface{})
	// 快捷弹出元素(左)
	Pop() interface{}
	// 删除一个元素
	Remove(interface{})
	// 获得链表总长
	Len() int
	// 返回相应对象的索引位置
	IndexOf(a interface{}) int
	// 获得链表的一个子链表的拷贝
	SubList(int, int) List
	// 获得链表的一个切片
	// Slice(int, int) List
	// 在指定位置插入一个元素
	Insert(p, i int, a interface{})
	// 清除链表
	Clear()
	// 获得链表头
	Iter(int) Iterator
	// 设置相等器
	SetEq(eq EqualFunc)
	// 设置比较器
	SetComp(comp Compartor)
}

type Iterator interface {
	Next() *Node
	HasNext() bool
}

type iter struct {
	curNode *Node
	mode    int
	n       int
}

type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

type doublelinkedlist struct {
	n    int
	head *Node
	tail *Node
	eq   EqualFunc
	comp Compartor
}

func NewList() List {
	return &doublelinkedlist{
		eq: defaultEq,
	}
}

func createNode(a interface{}) *Node {
	if a == nil {
		return nil
	}
	return &Node{
		Value: a,
	}
}

func (i *iter) Next() *Node {
	if i.n <= 0 {
		return nil
	}
	node := i.curNode
	if i.mode == Forward {
		i.curNode = i.curNode.prev
	} else if i.mode == Backward {
		i.curNode = i.curNode.next
	}
	i.n--
	return node
}

func (i *iter) HasNext() bool {
	if i.n <= 0 {
		return false
	}
	return i.curNode != nil
}

func defaultEq(n1 *Node, n2 *Node) bool {
	return n1.Value == n2.Value
}

func (self *doublelinkedlist) findNodeByIndex(i int) *Node {
	var index int = 0
	for node := self.head; node != nil; node = node.next {
		if i == index {
			return node
		}
		index++
	}
	return nil
}

func (self *doublelinkedlist) Iter(mode int) Iterator {
	iter := &iter{
		mode: mode,
		n:    self.n,
	}
	if mode == Forward {
		iter.curNode = self.tail
	} else if mode == Backward {
		iter.curNode = self.head
	} else {
		panic("unknow mode")
	}
	return iter
}

func (self *doublelinkedlist) findNodeByNode(n *Node) *Node {
	for node := self.head; node != nil; node = node.next {
		if self.eq(node, n) {
			return node
		}
	}
	return nil
}

func (self *doublelinkedlist) Lput(a interface{}) {
	node := createNode(a)
	if node == nil {
		return
	}
	if self.n > 0 {
		node.next = self.head
		node.prev = nil
		self.head.prev = node
		self.head = node
	} else {
		self.head = node
		self.tail = node
		node.prev = nil
		node.next = nil
	}
	self.n++
}

func (self *doublelinkedlist) Rput(a interface{}) {
	node := createNode(a)
	if self.n > 0 {
		node.prev = self.tail
		node.next = nil
		self.tail.next = node
		self.tail = node
	} else {
		self.head = node
		self.tail = node
		node.prev = nil
		node.next = nil
	}
	self.n++
}

func (self *doublelinkedlist) Put(a interface{}) {
	self.Rput(a)
}

func (self *doublelinkedlist) Pop() interface{} {
	return self.Lpop()
}

func (self *doublelinkedlist) IndexOf(a interface{}) int {
	index := -1
	witchnode := createNode(a)
	for node := self.head; node != nil; node = node.next {
		index++
		if self.eq(node, witchnode) {
			return index
		}
	}
	return -1
}

func (l *doublelinkedlist) Get(i int) interface{} {
	return l.findNodeByIndex(i).Value
}

func (self *doublelinkedlist) Lpop() interface{} {
	if self.n == 0 {
		return nil
	}
	node := self.head
	self.n--
	if self.n > 0 {
		self.head = node.next
		self.head.prev = nil
	} else {
		self.head = nil
		self.tail = nil
	}
	node.next = nil
	node.prev = nil
	return node.Value
}

func (self *doublelinkedlist) Rpop() interface{} {
	if self.n == 0 {
		return nil
	}
	node := self.tail
	self.n--
	if self.n > 0 {
		self.tail = node.prev
		self.tail.next = nil
	} else {
		self.tail = nil
		self.head = nil
	}
	node.next = nil
	node.prev = nil
	return node.Value
}

func (self *doublelinkedlist) Remove(a interface{}) {
	node := self.findNodeByNode(createNode(a))
	if node == nil {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		self.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		self.tail = node.prev
	}
	if node.Value != nil {
		node.Value = nil
	}
	self.n--
}

func (l *doublelinkedlist) Len() int {
	return l.n
}

func (l *doublelinkedlist) SetEq(eq EqualFunc) {
	l.eq = eq
}

func (l *doublelinkedlist) SetComp(comp Compartor) {
	l.comp = comp
}

func (l *doublelinkedlist) Insert(p, i int, a interface{}) {
	if p != After && p != Before {
		return
	}
	node := l.findNodeByIndex(i)
	if node == nil {
		if l.Len() == 0 {
			l.Put(a)
		}
		return
	}
	newnode := createNode(a)
	if newnode == nil {
		return
	}
	if p == After {
		newnode.prev = node
		newnode.next = node.next
		if node.next == nil {
			l.tail = newnode
		} else {
			node.next.prev = newnode
		}
		node.next = newnode
	} else if p == Before {
		newnode.prev = node.prev
		newnode.next = node
		if node.prev == nil {
			l.head = newnode
		} else {
			node.prev.next = newnode
		}
		node.prev = newnode
	}
	l.n++
}

func (l *doublelinkedlist) Clear() {
	l.n = 0
	l.head = nil
	l.tail = nil
}

func (self *doublelinkedlist) SubList(fromIndex int, toIndex int) List {
	if fromIndex < 0 || fromIndex >= toIndex || toIndex < 0 {
		return NewList()
	}
	list := NewList()
	var index int
	for node := self.head; node != nil; node = node.next {
		if index >= fromIndex {
			list.Put(node.Value)
		}
		index++
		if index == toIndex {
			break
		}
	}
	return list
}

/*
func (self *doublelinkedlist) Slice(fromIndex int, toIndex int) List {
	if fromIndex < 0 || fromIndex >= toIndex || toIndex < 0 {
		return NewList()
	}
	var (
		index   int
		n       int
		infirst bool = true
		list         = &doublelinkedlist{
			eq: self.eq,
		}
	)
	for node := self.head; node != nil; node = node.next {
		if index >= fromIndex {
			if infirst {
				list.head = node
				infirst = false
			}
			n++
		}
		index++
		if index == toIndex {
			list.tail = node
			break
		}
	}
	if n != 0 {
		list.n = n
	}
	if toIndex >= self.n-1 {
		list.tail = self.tail
	}
	return list
}
*/
