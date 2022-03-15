package main

import "fmt"

type MinStack struct {
	stack []int
	minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}

}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)
	} else {
		if x < this.minStack[len(this.minStack) - 1] {
			this.minStack = append(this.minStack, x)
		} else {
			this.minStack = append(this.minStack, this.minStack[len(this.minStack) - 1])
		}
	}

}


func (this *MinStack) Pop()  {
	this.stack = this.stack[: len(this.minStack) - 1]
	this.minStack = this.minStack[: len(this.minStack) - 1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.minStack) - 1]
}


func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack) - 1]

}


func main()  {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Pop()
	param_3 := obj.Top()
	fmt.Println(param_3)
	param_4 := obj.Min()
	fmt.Println(param_4)
}
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
