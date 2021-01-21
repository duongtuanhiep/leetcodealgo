package array.q1679;

import java.util.*;

/**
 * Question 1679: https://leetcode.com/problems/max-number-of-k-sum-pairs/
 * 
 * Idea: Create a hashmap to count the reccurence of each 
 */

class Solution {
	public int maxOperations(int[] nums, int k) {
		int res = 0;
		HashMap<Integer, Integer> occur = new HashMap<Integer, Integer>();
		for (int numero : nums)
			occur.put(numero, occur.getOrDefault(numero, 0) + 1);

		for (int numero : nums) {
			int target = k - numero;
			if (target == numero && occur.get(numero) > 1) {
				occur.put(numero, occur.get(numero) - 2);
				res++;
			} else if (target != numero && occur.getOrDefault(numero, 0) > 0 && occur.getOrDefault(target, 0) > 0) {
				occur.put(numero, occur.get(numero) - 1);
				occur.put(target, occur.get(target) - 1);
				res++;
			}
		}
		return res;
	}
}