package main

import "fmt"


type Queue struct {
	list []interface{}
	tail int
	head int
	size int
}


func NewQueue(size int) *Queue {
	size = size + 1
	return &Queue{
		list: make([]interface{}, size),
		size :size,
	}
}

func (q *Queue) Push(val interface{}) bool {
	if q.IsFull(){
		return false
	}
	q.list[q.tail] = val
	if q.tail + 1 == q.size{
		q.tail = 0
	}else{
		q.tail ++
	}
	return true
}


func (q *Queue) Pop() (interface{}, bool) {
	if q.IsEmpty(){
		return nil, false
	}
	val := q.list[q.head]
	if q.head + 1 == q.size{
		q.head = 0
	}else{
		q.head ++
	}	
	return val, true


}

func (q *Queue) IsFull() bool {
	return (q.tail + 1) % q.size == q.head
}

func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}
func (q *Queue)Show()  {
	if q.IsEmpty(){
		fmt.Println("empty")
		return 
	}
	index := q.head
	for{
		val := q.list[index]
		ss := fmt.Sprintf("index:%d, val:%v", index, val)
		fmt.Println(ss)
		if index + 1 == q.size{
			index = 0
		}else{
			index ++
		}	
		if index == q.tail{
			break
		}
	}
}
func (q *Queue)GetTail() int  {
	return q.tail
}

func (q *Queue)GetHead() int  {
	return q.head
}

func (q *Queue)GetHeadAndTail()  {
	fmt.Println(q.head, q.tail)
}


func main()  {
	que := NewQueue(3)	
	que.Push(1)
	que.Push(2)
	que.Push(3)
	que.Push(4)
	que.Show()
	fmt.Println(que.Pop())
	fmt.Println(que.Pop())
	fmt.Println(que.Pop())
	que.GetHeadAndTail()
	que.Push(4)
	que.Push(5)
	que.GetHeadAndTail()
	que.Show()
}