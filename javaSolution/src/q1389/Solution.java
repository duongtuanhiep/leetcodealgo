package q1389;

import java.util.*;
import java.lang.*;
import java.io.*;

/*
 * https://leetcode.com/problems/create-target-array-in-the-given-order/
 * 
 * 
 * */

class Solution {
	public int[] createTargetArray(int[] nums, int[] index) {
		int[] res = new int[nums.length];
		for (int i = 0; i < index.length; i++) {
			System.arraycopy(res, index[i], res, index[i] + 1, res.length - index[i] - 1);
			res[index[i]] = nums[i];
		}
		return res;
	}
}