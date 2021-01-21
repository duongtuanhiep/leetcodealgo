package dp.Q5;

/*
 * Question 5: https://leetcode.com/problems/longest-palindromic-substring/
 * 
 * 
 * **/

class Solution {
	public String longestPalindrome(String s) {
		String res = "";
		for (int i = 0; i < s.length(); i++) { // Iterating the start pos of str
			for (int j = i; j < s.length(); j++) { // iterating the end pos of str
				boolean isPal = true;
				int h = j;
				for (int k = i; k <= h; k++, h--) { // Checks for palindromic
					if (s.charAt(k) != s.charAt(h)) {
						isPal = false;
						break;
					}
				}
				if (j - i + 1 > res.length() && isPal) {
					res = s.substring(i, j + 1);
				}
			}
		}
		return res;
	}
}

/*
 * solution for real 
 class Solution {
    public String longestPalindrome(String s) {

        if (s == null || "".equals(s)) {
            return s;
        }
        
        int len = s.length();

        String ans = "";
        int max = 0;

        boolean[][] dp = new boolean[len][len];

        for (int j = 0; j < len; j++) {
            
            for (int i = 0; i <= j; i++) {
                
                boolean judge = s.charAt(i) == s.charAt(j);
                
                dp[i][j] = j - i > 2 ? dp[i + 1][j - 1] && judge : judge;

                if (dp[i][j] && j - i + 1 > max) {
                    max = j - i + 1;
                    ans = s.substring(i, j + 1);
                }
            }
        }
        return ans;
    }
}
 * */
