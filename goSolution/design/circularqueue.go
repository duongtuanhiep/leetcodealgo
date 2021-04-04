package design
/*
Question 622: https://leetcode.com/problems/design-circular-queue/submissions/


Idea: 
- Utilising Slice, keeping initial size. 
**/


type MyCircularQueue struct {
    Queue []int
	Size int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{Queue: make([]int, 0),Size: k }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.Queue) < this.Size { 
		this.Queue = append(this.Queue, value)
		return true
	}
	return false
}


func (this *MyCircularQueue) DeQueue() bool {
	if len(this.Queue) == 0 {
		return false 
	}
    this.Queue = this.Queue[1:]
	return true
}


func (this *MyCircularQueue) Front() int {
	if len(this.Queue) > 0 {
		return this.Queue[0]
	}
	return -1
}


func (this *MyCircularQueue) Rear() int {
    if len(this.Queue) > 0 {
		return this.Queue[len(this.Queue)-1]
	}
	return -1
}


func (this *MyCircularQueue) IsEmpty() bool {
    return len(this.Queue) == 0 
}


func (this *MyCircularQueue) IsFull() bool {
    return len(this.Queue) == this.Size
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