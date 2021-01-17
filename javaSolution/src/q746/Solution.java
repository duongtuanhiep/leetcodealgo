package q746;

import java.util.*;
import java.lang.*;
import java.io.*;

/*
 * https://leetcode.com/problems/min-cost-climbing-stairs/
 * 
 * Idea: 
 * - from the end of the array we can find the minimum cost to reach an index. 
 * - Bottom up. 
 * 
 * */

class Solution {
	public int minCostClimbingStairs(int[] cost) {
		if (cost.length < 2) {
			return cost[0] < cost[1] ? cost[1] : cost[0];
		}

		for (int i = cost.length - 3; i >= 0; i--) {
			cost[i] = cost[i + 1] > cost[i + 2] ? cost[i] + cost[i + 2] : cost[i] + cost[i + 1];
		}

		if (cost[1] > cost[0]) {
			return cost[0];
		}
		return cost[1];
	}
}
