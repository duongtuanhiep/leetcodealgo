package design

/*

Question 1352: https://leetcode.com/problems/product-of-the-last-k-numbers/

**/

type ProductOfNumbers struct {
	sums []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.sums = []int{}
		return
	}

	if len(this.sums) == 0 {
		this.sums = append(this.sums, num)
	} else {
		this.sums = append(this.sums, this.sums[len(this.sums)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	// case when we truncate arr due to 0 input
	if k > len(this.sums) {
		return 0
	}

	if k == len(this.sums) {
		return this.sums[len(this.sums)-1]
	}
	return this.sums[len(this.sums)-1] / this.sums[len(this.sums)-1-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
