package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	countList int
	frontItem *ListItem
	backItem  *ListItem
}

func (l *list) Len() int {
	return l.countList
}

func (l *list) Front() *ListItem {
	return l.frontItem
}

func (l *list) Back() *ListItem {
	return l.backItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	var n ListItem
	n.Value = v
	if l.frontItem != nil {
		l.frontItem.Prev = &n
		n.Next = l.frontItem
	} else {
		l.backItem = &n
	}
	l.frontItem = &n
	l.countList++
	return &n
}

func (l *list) PushBack(v interface{}) *ListItem {
	var b ListItem
	b.Value = v
	if l.backItem != nil {
		l.backItem.Next = &b
		b.Prev = l.backItem
	} else {
		l.frontItem = &b
	}
	l.backItem = &b
	l.countList++
	return &b
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil && i.Next != nil {
		if i.Prev.Next != i || i.Next.Prev != i {
			return
		}
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	if l.backItem == i {
		if i.Prev != nil {
			i.Prev.Next = nil
		}
		l.backItem = l.backItem.Prev
	}
	if l.frontItem == i {
		if i.Next != nil {
			i.Next.Prev = nil
		}
		l.frontItem = l.frontItem.Next
	}
	l.countList--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
