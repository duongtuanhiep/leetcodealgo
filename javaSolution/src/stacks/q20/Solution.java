package stacks.q20;

import java.io.*;
import java.util.*;

/*
 * Question 20: https://leetcode.com/problems/valid-parentheses/
 * 
 * 
 * **/

class Solution {
	public boolean isValid(String s) {
		Stack<Character> lastOpen = new Stack<Character>();
		for (int i = 0; i < s.length(); i++) {
			if (lastOpen.empty()) {
				if (s.charAt(i) != '(' && s.charAt(i) != '{' && s.charAt(i) != '[') {
					return false;
				}
				lastOpen.push(s.charAt(i));
			} else if (s.charAt(i) == '(' || s.charAt(i) == '{' || s.charAt(i) == '[') {
					lastOpen.push(s.charAt(i));
			} else {
				Character last = lastOpen.pop();
				if (last == '(' && s.charAt(i) != ')') {
					return false;
				} else if (last == '[' && s.charAt(i) != ']') {
					return false;
				} else if (last == '{' && s.charAt(i) != '}') {
					return false;
				}
			}
		}
		return lastOpen.empty();
	}
}