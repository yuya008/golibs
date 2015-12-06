package doublelinkedlist

import (
	"log"
	"testing"
)

var data = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestLput(t *testing.T) {
	list := NewList()
	for _, e := range data {
		list.Lput(e)
	}
	var i int = len(data) - 1
	for iter := list.Iter(Backward); iter.HasNext(); {
		if val, ok := iter.Next().Value.(int); ok {
			if data[i] != val {
				t.Error("error")
			} else {
				t.Log(val)
			}
		} else {
			t.Error("error type")
		}
		i--
	}
}

func TestRput(t *testing.T) {
	list := NewList()
	for _, e := range data {
		list.Rput(e)
	}
	var i int = 0
	for iter := list.Iter(Backward); iter.HasNext(); {
		if val, ok := iter.Next().Value.(int); ok {
			if data[i] != val {
				t.Error("error")
			} else {
				t.Log(val)
			}
		} else {
			t.Error("error type")
		}
		i++
	}
}

func TestGet(t *testing.T) {
	list := NewList()
	for _, e := range data {
		list.Lput(e)
	}
	if v, ok := list.Get(0).(int); ok {
		if v == 9 {
			t.Log(v)
		} else {
			t.Error("error")
		}
	} else {
		t.Error("error type")
	}
	if v, ok := list.Get(9).(int); ok {
		if v == 0 {
			t.Log(v)
		} else {
			t.Error("error")
		}
	} else {
		t.Error("error type")
	}
	if v, ok := list.Get(5).(int); ok {
		if v == 4 {
			t.Log(v)
		} else {
			t.Error("error")
		}
	} else {
		t.Error("error type")
	}
	if v, ok := list.Get(2).(int); ok {
		if v == 7 {
			t.Log(v)
		} else {
			t.Error("error")
		}
	} else {
		t.Error("error type")
	}
}

func TestLpop(t *testing.T) {
	list := NewList()
	for _, e := range data {
		list.Lput(e)
	}
	var i = len(data) - 1
	for val := list.Lpop(); val != nil; val = list.Lpop() {
		if v, ok := val.(int); ok {
			if v != data[i] {
				t.Error("error")
			} else {
				t.Log(v)
			}
		} else {
			t.Error("error type")
		}
		i--
	}
	if list.Len() != 0 {
		t.Error("len()!=0")
	}
	for _, e := range data {
		list.Rput(e)
	}
	i = 0
	for val := list.Lpop(); val != nil; val = list.Lpop() {
		if v, ok := val.(int); ok {
			if v != data[i] {
				t.Error("error")
			} else {
				t.Log(v)
			}
		} else {
			t.Error("error type")
		}
		i++
	}
	if list.Len() != 0 {
		t.Error("len()!=0")
	}
}

func TestRpop(t *testing.T) {
	list := NewList()
	for _, e := range data {
		list.Lput(e)
	}
	var i = 0
	for val := list.Rpop(); val != nil; val = list.Rpop() {
		if v, ok := val.(int); ok {
			if v != data[i] {
				t.Error("error")
			} else {
				t.Log(v)
			}
		} else {
			t.Error("error type")
		}
		i++
	}
	if list.Len() != 0 {
		t.Error("len()!=0")
	}
	for _, e := range data {
		list.Rput(e)
	}
	i = 0
	for val := list.Lpop(); val != nil; val = list.Lpop() {
		if v, ok := val.(int); ok {
			if v != data[i] {
				t.Error("error")
			} else {
				t.Log(v)
			}
		} else {
			t.Error("error type")
		}
		i++
	}
	if list.Len() != 0 {
		t.Error("len()!=0")
	}
}

func TestPut(t *testing.T) {
	// like Rput
}

func TestPop(t *testing.T) {
	// like Lpop
}

func TestRemove(t *testing.T) {
	list := NewList()
	list.SetEq(eq1)
	for _, val := range data {
		list.Put(val)
	}
	for _, val := range data {
		list.Remove(val)
	}
	if list.Len() != 0 {
		t.Error("error")
	}
}

func TestLen(t *testing.T) {
	list := NewList()
	for _, val := range data {
		list.Put(val)
	}
	if list.Len() != len(data) {
		t.Error("error")
	}
}

func eq1(n1 *Node, n2 *Node) bool {
	var v1, v2 int
	var ok bool
	if v1, ok = n1.Value.(int); !ok {
		log.Panicln("error")
	}
	if v2, ok = n2.Value.(int); !ok {
		log.Panicln("error")
	}
	return v1 == v2
}

func TestIndexOf(t *testing.T) {
	list := NewList()
	list.SetEq(eq1)
	for _, val := range data {
		list.Put(val)
	}
	if list.IndexOf(0) != 0 {
		t.Error("error")
	}
	if list.IndexOf(9) != 9 {
		t.Error("error")
	}
	if list.IndexOf(4) != 4 {
		t.Error("error")
	}
	if list.IndexOf(100) != -1 {
		t.Error("error")
	}
}

func TestSubList(t *testing.T) {
	list := NewList()
	for _, val := range data {
		list.Put(val)
	}
	newlist := list.SubList(0, 3)
	var i int = 0
	for iter := newlist.Iter(Backward); iter.HasNext(); {
		if v, ok := iter.Next().Value.(int); ok {
			if v != data[i] {
				t.Error("error")
			} else {
				t.Log(v)
			}
		} else {
			t.Error("error")
		}
		i++
	}
	newlist = list.SubList(0, 0)
	if newlist.Len() != 0 {
		t.Error("error")
	}
	newlist = list.SubList(9, 5)
	if newlist.Len() != 0 {
		t.Error("error")
	}
	newlist = list.SubList(4, 4)
	if newlist.Len() != 0 {
		t.Error("error")
	}
	newlist = list.SubList(4, 100)
	i = 4
	for iter := newlist.Iter(Backward); iter.HasNext(); {
		if v, ok := iter.Next().Value.(int); ok {
			if v != data[i] {
				t.Error("error")
			} else {
				t.Log(v)
			}
		} else {
			t.Error("error")
		}
		i++
	}
}

/*
func TestSlice(t *testing.T) {
	list := NewList()
	for _, val := range data {
		list.Put(val)
	}
	newslice := list.Slice(0, 8)
	for iter := newslice.Iter(Backward); iter.HasNext(); {
		if v, ok := iter.Next().Value.(int); ok {
			t.Log(v)
		} else {
			t.Error("error")
		}
	}
}*/

func TestInsert(t *testing.T) {
	list := NewList()
	data := []int{4, 5, 6, 7, 34, 65}
	should := []int{4, 65, 34, 7, 6, 5}
	for _, d := range data {
		list.Insert(After, 0, d)
	}
	t.Log("Len() == ", list.Len())
	var i int
	for iter := list.Iter(Backward); iter.HasNext(); i++ {
		if v, ok := iter.Next().Value.(int); ok {
			if should[i] != v {
				t.Error("error")
			}
		} else {
			t.Error("error")
		}
	}
	list.Clear()
	should = []int{10, 65, 34, 7, 6, 5, 4, 100}
	list.Insert(Before, 0, 100)
	list.Insert(Before, 0, 10)
	for _, d := range data {
		list.Insert(Before, 1, d)
	}
	i = 0
	for iter := list.Iter(Backward); iter.HasNext(); i++ {
		if v, ok := iter.Next().Value.(int); ok {
			if should[i] != v {
				t.Error("error")
			}
		} else {
			t.Error("error")
		}
	}
}

func TestClear(t *testing.T) {
	list := NewList()
	data := []string{"a", "b", "c", "d", "e"}
	for _, s := range data {
		list.Put(s)
	}
	if list.Len() != 5 {
		t.Error("error")
	}
	list.Clear()
	if list.Len() != 0 {
		t.Error("error")
	}
}
