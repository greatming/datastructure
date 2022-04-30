package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX_LEVEL = 15
)

type Level struct{
	Next *Node
}
type Node struct{
	Val int
	Levels []*Level
}

type SkipList struct {
	Head *Node
	Level int
	Length int
}


func NewNode(val int, level int) *Node {
	n := Node{
		Val: val,
	}
	n.Levels = make([]*Level, level)
	return &n
}


func NewSkipList() *SkipList {
	skip := SkipList{
		Level: 1,
		Length: 0,
	}
	node := NewNode(0, MAX_LEVEL)
	skip.Head = node
	fmt.Println(node)
	
	return &skip
}

func (s *SkipList)GetRandLevel() int  {
	p := 0.5
	level := 1
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Float64() 
	if rand.Float64() < p && level < MAX_LEVEL{
		level ++
	}
	return level
}
func (s *SkipList) Add(val int) bool {
	update := make([]*Node, s.Level)
	tmp := s.Head
	for i := s.Level-1; i >= 0; i--{
		for tmp.Levels[i].Next != nil && tmp.Levels[i].Next.Val < val{
			tmp = tmp.Levels[i].Next
		}
		//查询已经存在val 返回false
		if tmp.Levels[i].Next.Val == val{
			return false
		}
		update[i] = tmp
	}

	level := s.GetRandLevel()
	node := NewNode(val, level)

	for i := level -1; i >= 0; i--{
		//如果当前层级的next为空， 说明超过了目前最高层级，需要在head 进行关联
		if update[i].Levels[i].Next == nil{
			s.Head.Levels[i].Next = node
			continue
		}
		node.Levels[i].Next = update[i].Levels[i].Next
		update[i].Levels[i].Next = node
	}
	if level > s.Level{
		s.Level = level
	}
	s.Length++
	return true
}

func (s *SkipList)Del(val int) bool  {
	list := make([]*Node, s.Level)
	tmp := s.Head
	for i := s.Level-1; i > 0; i--{
		for tmp.Levels[i].Next != nil && tmp.Levels[i].Next.Val < val{
			tmp = tmp.Levels[i].Next
		}
		list[i] = tmp
	}

	if list[0].Levels[0].Next.Val != val{
		return false
	}
	node := list[0].Levels[0].Next

	for i := s.Level -1; i >= 0; i--{
		list[i].Levels[i].Next = node.Levels[i].Next
		node.Levels[i].Next = nil
	}
	return true
}

func (s *SkipList) Search(val int) (*Node, bool) {
	tmp := s.Head
	var node *Node
	for i := s.Level-1; i >= 0; i--{
		for tmp.Levels[i].Next != nil && tmp.Levels[i].Next.Val < val{
			tmp = tmp.Levels[i].Next
		}
		if tmp.Levels[i].Next.Val == val{
			node = tmp.Levels[i].Next
			break
		}
	}
	if node == nil{
		return nil, false
	}
	return node, true
}

func (s *SkipList) Show()  {
	tmp := s.Head
	temp := "  -->  "
	retList := make([]string, s.Level)
	for i := s.Level -1; i >= s.Level; i--{
		line := fmt.Sprintf("%d级索引:   header%s", i, temp)
		for tmp.Levels[i].Next != nil{
			val := tmp.Levels[i].Next.Val
			line = fmt.Sprintf("%s, %d%s", line, val, temp)
		}
		retList[i] = line
	}

	for _, line := range retList{
		fmt.Println(line)
	}
}


func main()  {
	ll := NewSkipList()
	ll.Add(3)
	ll.Show()
}

