package varChannel

import "sync"

type Element struct {
	next, prev *Element
	list       *List
	Value      interface{}
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
	l    sync.Mutex
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func NewList() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}

func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e

	return e
}

func (l *List) Remove(e *Element) interface{} {
	l.l.Lock()
	defer l.l.Unlock()
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

func (l *List) PushFront(v interface{}) *Element {
	l.l.Lock()
	defer l.l.Unlock()
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List) PushBack(v interface{}) *Element {
	l.l.Lock()
	defer l.l.Unlock()
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	l.l.Lock()
	defer l.l.Unlock()
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	l.l.Lock()
	defer l.l.Unlock()
	if mark.list != l {
		return nil
	}

	return l.insertValue(v, mark)
}

func (l *List) MoveToFront(e *Element) {
	l.l.Lock()
	defer l.l.Unlock()
	if e.list != l || l.root.next == e {
		return
	}

	l.move(e, &l.root)
}

func (l *List) MoveToBack(e *Element) {
	l.l.Lock()
	defer l.l.Unlock()
	if e.list != l || l.root.prev == e {
		return
	}

	l.move(e, l.root.prev)
}

func (l *List) MoveBefore(e, mark *Element) {
	l.l.Lock()
	defer l.l.Unlock()
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

func (l *List) MoveAfter(e, mark *Element) {
	l.l.Lock()
	defer l.l.Unlock()
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

func (l *List) PushBackList(other *List) {
	l.l.Lock()
	defer l.l.Unlock()
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

func (l *List) PushFrontList(other *List) {
	l.l.Lock()
	defer l.l.Unlock()
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}
