package queue

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	q := NewQueue()
	q.Push([]string{"1","2"})
	//url := reflect.ValueOf(q.Pop())
	fmt.Println(q.Pop().([]string))
}
