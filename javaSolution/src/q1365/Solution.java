package q1365;

import java.util.*;
import java.lang.*;
import java.io.*;

/* https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

Idea: sort, gets left most elem of each value, recreate arr

class Solution {
	public int[] smallerNumbersThanCurrent(int[] nums) {
		int[] res = nums.clone();
		HashMap<Integer,Integer> pos = new HashMap<Integer,Integer>();
		int i;
		Arrays.sort(nums);
		System.out.println(Arrays.toString(res));
		for (i=0; i <nums.length;i++) {
			if (!pos.containsKey(nums[i])) pos.put(nums[i],i);
		}
		for (i=0;i<res.length;i++) {
			res[i] = pos.get(res[i]);
		}
		
		return res;
	}
}
 **/

/*
 * Since we know the possible range of the number, we can bucket count the reccurence
 * of elements that are greater other previous elem
 * */

class Solution {
	public int[] smallerNumbersThanCurrent(int[] nums) {
		int[] bucketCount = new int[101]; // range 0->100
		for (int i = 0; i < nums.length; i++) {
			bucketCount[nums[i]] += 1;
		}

		for (int i = 1; i < bucketCount.length; i++) {
			bucketCount[i] += bucketCount[i - 1];
		}

		for (int k = 0; k < nums.length; k++) {
			int pos = nums[k];
			nums[k] = pos == 0 ? 0 : bucketCount[pos - 1];
		}

		return nums;
	}
}
