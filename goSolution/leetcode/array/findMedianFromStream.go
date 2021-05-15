package array

import "sort"

/*
Question 295: https://leetcode.com/problems/find-median-from-data-stream/

**/

type MedianFinder struct {
	Arr []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	var arr []int
	return MedianFinder{arr}
}

func (this *MedianFinder) AddNum(num int) {
	this.Arr = append(this.Arr, num)
	sort.Ints(this.Arr)
	// BS leftmost location.
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.Arr)%2 == 1 {
		return float64(this.Arr[len(this.Arr)/2])
	}
	return float64(this.Arr[len(this.Arr)/2-1]+this.Arr[len(this.Arr)/2]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
