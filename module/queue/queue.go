package queue

import (
	"container/list"
	"fmt"
	"sync"
)

type Queue struct {
	l    sync.Mutex
	data *list.List
}

func NewQueue() *Queue {
	q := new(Queue)
	q.data = list.New()
	return q
}

func (q *Queue) Push(v interface{}) *list.Element {
	q.l.Lock()
	defer q.l.Unlock()
	return q.data.PushFront(v)
}

func (q *Queue) PushBack(v interface{}) *list.Element {
	q.l.Lock()
	defer q.l.Unlock()
	return q.data.PushBack(v)
}

func (q *Queue) Pop() interface{} {
	q.l.Lock()
	defer q.l.Unlock()
	iter := q.data.Back()
	if nil == iter {
		return nil
	}
	v := iter.Value
	q.data.Remove(iter)
	return v
}

//返回pop列表和实际长度
func (q *Queue) Pops(num int) ([]interface{}, int) {
	vals := make([]interface{}, num)
	i := 0
	q.l.Lock()
	defer q.l.Unlock()
	for {
		if i >= num {
			break
		}
		iter := q.data.Back()
		if iter == nil {
			return vals, i
		}
		q.data.Remove(iter)
		vals[i] = iter.Value
		i++
	}
	if i < num {
		return vals[0:i], i
	}
	return vals, i
}

func (q *Queue) Remove(v *list.Element) (interface{}) {
	q.l.Lock()
	defer q.l.Unlock()
	return q.data.Remove(v)
}

func (q *Queue) Len() int {
	return q.data.Len()
}

func (q *Queue) Dump() {
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		fmt.Println("item:", iter.Value)
	}
}