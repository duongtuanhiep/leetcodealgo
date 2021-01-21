package dfs.Q1448;

/**
 * Question 1448: https://leetcode.com/problems/count-good-nodes-in-binary-tree/
 * 
 * Definition for a binary tree node.
 */
class TreeNode {
	int val;
	TreeNode left;
	TreeNode right;

	TreeNode() {
	}

	TreeNode(int val) {
		this.val = val;
	}

	TreeNode(int val, TreeNode left, TreeNode right) {
		this.val = val;
		this.left = left;
		this.right = right;
	}
}
/*
 * Idea: Create a recursive function that will do DFS and remember the max value
 * of path.
 */

class Solution {
	public int goodNodes(TreeNode root) {
		return isGood(root, root.val);
	}

	public int isGood(TreeNode cur, int maxVal) {
		if (cur == null) {
			return 0;
		}
		int result = 0;
		if (cur.val >= maxVal) {
			maxVal = cur.val;
			result++;
		}
		if (cur.left != null) {
			result += isGood(cur.left, maxVal);
		}
		if (cur.right != null) {
			result += isGood(cur.right, maxVal);
		}

		return result;
	}
}