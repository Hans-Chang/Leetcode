/*
 * 剑指 Offer 09. 用两个栈实现队列
 */

package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	data *list.List
}

func (this *Stack) Push(value interface{})  {
	if this.data == nil {
		this.data = new(list.List)
	}
	this.data.PushBack(value)
}

func (this *Stack) Pop() interface{} {
	result := this.data.Front().Value
	this.data.Remove(this.data.Front())
	return result
}

func (this *Stack) Top() interface{} {
	result := this.data.Front().Value
	return result
}

func (this *Stack) Len() int {
	if this.data == nil {
		return 0
	}
	return this.data.Len()
}

type CQueue struct {
	s1 Stack
	s2 Stack
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
	this.s1.Push(value)
}


func (this *CQueue) DeleteHead() int {
	if this.s1.Len() > 0 {
		return this.s1.Pop().(int)
	}
	if this.s2.Len() <= 0 {
		return -1
	}
	for this.s2.Len() > 0 {
		this.s1.Push(this.s2.Pop())
	}
	return this.s1.Pop().(int)
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main()  {
	obj := Constructor()
	obj.AppendTail(2)
	obj.AppendTail(3)
	fmt.Println(obj.DeleteHead())
	obj.AppendTail(5)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}
