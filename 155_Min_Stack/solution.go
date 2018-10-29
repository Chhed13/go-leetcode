package _155_Min_Stack

import "math"

/*https://leetcode.com/problems/min-stack/
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
Example:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
*/

type MinStack struct {
    stack []int
    minElement int

}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack {
        []int{}, math.MaxInt32,
    }
}


func (this *MinStack) Push(x int)  {
    if x < this.minElement {
        this.minElement = x
    }
    this.stack = append(this.stack,x)
}


func (this *MinStack) Pop()  {
    if len(this.stack) == 0 {
        return
    }
    x := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    if x == this.minElement {
        this.minElement = math.MaxInt32
        for _,e := range this.stack {
            if e < this.minElement {
                this.minElement = e
            }
        }
    }
}


func (this *MinStack) Top() int {
    if len(this.stack) == 0 {
        return -1
    }
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.minElement
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
