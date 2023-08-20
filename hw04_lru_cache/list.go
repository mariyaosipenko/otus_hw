package hw04lrucache

import (
	"golang.org/x/exp/slices"
)

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
	allItem   []*ListItem
	frontItem *ListItem
	backItem  *ListItem
}

func (l *list) Len() int {
	return len(l.allItem)
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
	l.allItem = append(l.allItem, &n)
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
	l.allItem = append(l.allItem, &b)
	return &b
}

func (l *list) Remove(i *ListItem) {
	index := slices.IndexFunc(l.allItem, func(c *ListItem) bool { return i == c })
	elm := l.allItem[index]
	if elm.Prev != nil && elm.Next != nil {
		elm.Prev.Next = elm.Next
		elm.Next.Prev = elm.Prev
	}
	if elm.Prev != nil && elm.Next == nil {
		elm.Prev.Next = nil
	}
	if elm.Prev == nil && elm.Next != nil {
		elm.Next.Prev = nil
	}
	if l.backItem == i {
		l.backItem = l.backItem.Prev
	}
	if l.frontItem == i {
		l.frontItem = l.frontItem.Next
	}
	l.allItem = append(l.allItem[:index], l.allItem[index+1:]...)
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
