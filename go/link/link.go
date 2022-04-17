package main

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

func NewNode(val interface{}) *Node {
	return &Node{val, nil}
}

type LinkList struct {
	Head *Node
}

func NewLinkList() *LinkList {
	return &LinkList{}
}

func (l *LinkList) InsertBefore(n *Node, val interface{}) bool {
	if n == nil {
		return false
	}
	newNode := NewNode(val)
	if l.Head == n {
		newNode.Next = l.Head
		l.Head = newNode
		return true
	}
	pre := l.Head
	for {
		if pre.Next == n {
			break
		}
		pre = pre.Next
	}
	if pre == nil {
		return false
	}
	newNode.Next = n
	pre.Next = newNode
	return true
}

func (l *LinkList) InsertAfter(n *Node, val interface{}) bool {
	if n == nil {
		return false
	}
	cur := l.Head
	for cur != n {
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	if cur != n {
		return false
	}
	newNode := NewNode(val)
	newNode.Next = cur.Next
	cur.Next = newNode
	return true
}

func (l *LinkList) InsertToHead(val interface{}) bool {
	if l.IsEmpty() {
		l.Head = NewNode(val)
		return true
	}
	return l.InsertBefore(l.Head, val)
}

func (l *LinkList) Insert(val interface{}) bool {
	if l.IsEmpty() {
		l.Head = NewNode(val)
		return true
	}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	return l.InsertAfter(cur, val)
}

func (l *LinkList) IsEmpty() bool {
	return l.Head == nil
}

func (l *LinkList) RemoveNode(val interface{}) bool {
	if l.Head == nil {
		return false
	}
	if l.Head.Val == val {
		l.Head = l.Head.Next
		return true
	}
	pre := l.Head
	for {
		if pre.Next == nil {
			break
		}
		if pre.Next.Val == val {
			break
		}
		pre = pre.Next
	}
	if pre.Next == nil {
		return false
	}
	pre.Next = pre.Next.Next
	return true
}

func (l *LinkList) Show() {
	cur := l.Head
	index := 0
	for {
		index++
		if cur == nil {
			break
		}
		fmt.Println(fmt.Sprintf("%d:%v", index, cur.Val))
		cur = cur.Next
	}
}

func (l *LinkList)Reversal() {
	if l.Head == nil || l.Head.Next == nil{
		return
	}
	pre := l.Head
	cur := l.Head.Next
	pre.Next = nil
	for {
		if cur == nil{
			break
		}
		fmt.Println(cur.Val)
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	fmt.Println(pre.Next.Val)
	l.Head = pre
}

func (l *LinkList)GetNodeByIndex(index int) *Node {
	if l.Head == nil{
		return nil
	}
	cur := l.Head
	i := 0
	for cur != nil{
		if index == i{
			return cur
		}
		i = i+1
		cur = cur.Next
	}
	return nil
}

func (l *LinkList)GetLastNode() *Node {
	if l.Head == nil{
		return nil
	}
	cur := l.Head
	for cur.Next != nil{
		cur = cur.Next
	}
	return cur
}


func (l *LinkList)IsPalindrome() bool  {
	if l.Head == nil || l.Head.Next == nil{
		return false
	}
	slow := l.Head.Next
	quick := l.Head
	for slow != nil && quick != nil &&  quick.Next != nil{
		slow  = slow.Next
		quick = quick.Next.Next
	}
	tmpList := make([]interface{}, 0)

	for slow != nil{
		tmpList = append(tmpList, slow.Val)
		slow = slow.Next
	}

	cur := l.Head
	res := true
	for  i := len(tmpList)-1; i >= 0; i--{
		v := tmpList[i]
		if v != cur.Val{
			res = false
			break
		}
		cur = cur.Next
	} 
	return res
}


func (l *LinkList)IsRing() bool  {
	if l.Head == nil || l.Head.Next == nil{
		return false
	}
	if l.Head == l.Head.Next {
		return true
	}
	slow := l.Head
	quick := l.Head.Next.Next
	for slow != nil && quick != nil &&  quick.Next != nil{
		if slow == quick{
			return true
		}
		slow  = slow.Next
		quick = quick.Next.Next
	}
	return false

}


func main() {
	link := NewLinkList()
	link.Insert("1")
	link.Insert("4")
	link.Insert("2")
	link.Insert("3")
	link.Insert("5")
	lastNode := link.GetLastNode()
	node := link.GetNodeByIndex(1)
	if node != nil{
		lastNode.Next = node
	}
	res := link.IsRing()
	fmt.Println(res)
	// res := link.IsPalindrome()
	// fmt.Println(fmt.Sprintf("IsPalindrome :%v", res))
}
