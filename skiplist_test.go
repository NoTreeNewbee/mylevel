package mylevel

import (
	"strings"
	"testing"
)

type testComparator struct{}

func (t *testComparator) Compare(x, y []byte) int {
	return strings.Compare(string(x), string(y))
}

func TestSkipList_Insert(t *testing.T) {
	skipList := NewSkipList(&testComparator{})
	skipList.Insert([]byte("1"))
	skipList.Insert([]byte("2"))
	skipList.Insert([]byte("3"))
	skipList.Insert([]byte("4"))
	skipList.Insert([]byte("31"))
}

func assertIterValidAndEqual(t *testing.T, it *skipListIter, val []byte) {
	if !it.Valid() {
		t.Fail()
		return
	}
	if string(it.Key()) != string(val) {
		t.Fail()
		return
	}
}

func TestSkipListIter_Next(t *testing.T) {
	skipList := NewSkipList(&testComparator{})
	skipList.Insert([]byte("1"))
	skipList.Insert([]byte("2"))
	skipList.Insert([]byte("3"))
	skipList.Insert([]byte("4"))
	skipList.Insert([]byte("31"))
	it := skipList.Iterator()
	assertIterValidAndEqual(t, it, []byte("1"))
	it.Next()
	assertIterValidAndEqual(t, it, []byte("2"))
	it.Next()
	assertIterValidAndEqual(t, it, []byte("3"))
	it.Next()
	assertIterValidAndEqual(t, it, []byte("31"))
	it.Next()
	assertIterValidAndEqual(t, it, []byte("4"))
	it.Next()
	if it.Valid() {
		t.Fail()
	}
}

func TestSkipListIter_Prev(t *testing.T) {
	skipList := NewSkipList(&testComparator{})
	skipList.Insert([]byte("1"))
	skipList.Insert([]byte("2"))
	skipList.Insert([]byte("3"))
	skipList.Insert([]byte("4"))
	skipList.Insert([]byte("31"))
	it := skipList.Iterator()
	it.Seek([]byte("4"))
	assertIterValidAndEqual(t, it, []byte("4"))
	it.Prev()
	assertIterValidAndEqual(t, it, []byte("31"))
	it.Prev()
	assertIterValidAndEqual(t, it, []byte("3"))
	it.Prev()
	assertIterValidAndEqual(t, it, []byte("2"))
	it.Prev()
	assertIterValidAndEqual(t, it, []byte("1"))
	it.Prev()
	if it.Valid() {
		t.Fail()
	}
}

func TestSkipListIter_Seek(t *testing.T) {
	skipList := NewSkipList(&testComparator{})
	skipList.Insert([]byte("1"))
	skipList.Insert([]byte("2"))
	skipList.Insert([]byte("3"))
	skipList.Insert([]byte("4"))
	skipList.Insert([]byte("31"))
	it := skipList.Iterator()
	it.Seek([]byte("3"))
	assertIterValidAndEqual(t, it, []byte("3"))
	it.Seek([]byte("39"))
	assertIterValidAndEqual(t, it, []byte("4"))
	it.Seek([]byte("41"))
	if it.Valid() {
		t.Fail()
		return
	}
}