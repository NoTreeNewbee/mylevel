package mylevel

import (
	"math/rand"
)

const (
	maxHeight = 12
)

type comparator interface {
	Compare(x, y []byte) int
}

type skipListNode struct {
	key  []byte
	val  []byte
	next []*skipListNode
}

func NewNodeWithRandomHeight(key []byte) *skipListNode {
	const ratio = 4
	height := 1
	for height < maxHeight && rand.Int()%ratio == 0 {
		height += 1
	}
	return &skipListNode{
		key:  key,
		next: make([]*skipListNode, height),
	}
}

func (node *skipListNode) Height() int {
	return len(node.next)
}

func (node *skipListNode) Next(n int) *skipListNode {
	return node.next[n]
}

func (node *skipListNode) SetNext(n int, x *skipListNode) {
	node.next[n] = x
}

type skipList struct {
	head       *skipListNode
	maxHeight  int
	comparator comparator
}

func NewSkipList(comparator comparator) *skipList {
	return &skipList{
		head: &skipListNode{
			key:  nil,
			next: make([]*skipListNode, maxHeight),
		},
		maxHeight:  1,
		comparator: comparator,
	}
}

// 找到大于等于x的节点，并且返回每层x之前的节点
func (s *skipList) findGreaterOrEqual(x []byte) (target *skipListNode, prev [maxHeight]*skipListNode) {
	i := s.maxHeight - 1
	n := s.head
	for i >= 0 {
		next := n.Next(i)
		if next != nil && s.comparator.Compare(next.key, x) < 0 {
			n = next
		} else {
			target = next // 这里虽然每一轮都赋值，但是最终第0层会被赋值为正确的值
			prev[i] = n   // 存下每层的上一个节点
			i -= 1
		}
	}
	return
}

func (s *skipList) findLessThan(x []byte) *skipListNode {
	i := s.maxHeight - 1
	n := s.head
	for i >= 0 {
		next := n.Next(i)
		if next == nil || s.comparator.Compare(next.key, x) >= 0 {
			i -= 1
		} else {
			n = next
		}
	}
	return n
}

func (s *skipList) Insert(key []byte) {
	_, prev := s.findGreaterOrEqual(key)
	n := NewNodeWithRandomHeight(key)
	if n.Height() > s.maxHeight {
		s.maxHeight = n.Height()
	}
	for i := 0; i < n.Height(); i++ {
		if prev[i] == nil {
			prev[i] = s.head
		}
		n.SetNext(i, prev[i].Next(i))
		prev[i].SetNext(i, n)
	}
}

func (s *skipList) Contains(key []byte) bool {
	return false
}

func (s *skipList) Iterator() *skipListIter {
	return &skipListIter{
		skipList: s,
		node:     s.head.Next(0),
	}
}

type skipListIter struct {
	skipList *skipList
	node     *skipListNode
}

func (it *skipListIter) Valid() bool {
	return it.node != nil
}

func (it *skipListIter) Next() {
	if it.node == nil {
		return
	}
	it.node = it.node.Next(0)
}

func (it *skipListIter) Prev() {
	if it.node == nil {
		return
	}
	n := it.skipList.findLessThan(it.node.key)
	if n == it.skipList.head {
		n = nil
	}
	it.node = n
}

func (it *skipListIter) Key() []byte {
	if it.node == nil {
		return nil
	}
	return it.node.key
}

func (it *skipListIter) Seek(key []byte) {
	it.node, _ = it.skipList.findGreaterOrEqual(key)
}
