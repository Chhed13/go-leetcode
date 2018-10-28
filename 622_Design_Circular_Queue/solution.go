package _622_Design_Circular_Queue

/*https://leetcode.com/problems/design-circular-queue/

Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.

Your implementation should support following operations:

MyCircularQueue(k): Constructor, set the size of the queue to be k.
Front: Get the front item from the queue. If the queue is empty, return -1.
Rear: Get the last item from the queue. If the queue is empty, return -1.
enQueue(value): Insert an element into the circular queue. Return true if the operation is successful.
deQueue(): Delete an element from the circular queue. Return true if the operation is successful.
isEmpty(): Checks whether the circular queue is empty or not.
isFull(): Checks whether the circular queue is full or not.


Example:

MyCircularQueue circularQueue = new MycircularQueue(3); // set the size to be 3
circularQueue.enQueue(1);  // return true
circularQueue.enQueue(2);  // return true
circularQueue.enQueue(3);  // return true
circularQueue.enQueue(4);  // return false, the queue is full
circularQueue.Rear();  // return 3
circularQueue.isFull();  // return true
circularQueue.deQueue();  // return true
circularQueue.enQueue(4);  // return true
circularQueue.Rear();  // return 4

Note:

All values will be in the range of [0, 1000].
The number of operations will be in the range of [1, 1000].
Please do not use the built-in Queue library.
*/

type MyCircularQueue struct {
    array []int
    head  int
    tail  int
}

func (this *MyCircularQueue) next(i int) int {
    if i >= len(this.array)-1 || i == -1 {
        return 0
    }
    return i + 1
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
    a := make([]int, k)
    for i := range a {
        a[i] = -1
    }

    return MyCircularQueue{a, -1, -1,}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }

    this.tail = this.next(this.tail)
    this.array[this.tail] = value
    if this.head == -1 {
        this.head = this.next(this.head)
    }
    return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }

    this.array[this.head] = -1
    if this.head == this.tail {
        this.head, this.tail = -1, -1
    } else {
        this.head = this.next(this.head)
    }

    return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
    if this.head < 0 {
        return -1
    }
    return this.array[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
    if this.tail < 0 {
        return -1
    }
    return this.array[this.tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
    return this.head < 0 && this.tail < 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
    return this.next(this.tail) == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
