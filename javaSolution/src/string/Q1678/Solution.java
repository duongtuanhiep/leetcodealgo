package string.Q1678;

/* 
 * Question 1678: https://leetcode.com/problems/goal-parser-interpretation/
 * 
 * String parser
 * **/

class Solution {
	public String interpret(String command) {
		String res = "";

		for (int i = 0; i < command.length(); i++) {
			if (command.charAt(i) == 'G') {
				res += "G";
			} else if (i + 1 < command.length()
					&& (Character.toString(command.charAt(i)) + Character.toString(command.charAt(i + 1))).equals("()")) {
				i++;
				res += "o";
			} else if (i + 1 < command.length()
					&& (Character.toString(command.charAt(i)) + Character.toString(command.charAt(i + 1))).equals("(a")) {
				i += 3;
				res += "al";
			}
			System.out.println(command.charAt(i));
			}

		return res;
	}
}