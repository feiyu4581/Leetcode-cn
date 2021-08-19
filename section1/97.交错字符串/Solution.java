/*
 * @lc app=leetcode.cn id=97 lang=java
 *
 * [97] 交错字符串
 */

// @lc code=start
class Solution {
    public boolean isInterleave(String s1, String s2, String s3) {
        if (s1.length() + s2.length() != s3.length()) {
            return false;
        }

        boolean[][] res = new boolean[s1.length() + 1][s2.length() + 1];
        res[0][0] = true;
        for (int i = 1; i <= s1.length(); i++) {
            res[i][0] = res[i - 1][0] && s1.charAt(i - 1) == s3.charAt(i - 1);
        }

        for (int j = 1; j <= s2.length(); j++) {
            res[0][j] = res[0][j - 1] && s2.charAt(j - 1) == s3.charAt(j - 1);
        }

        for (int i = 1; i <= s1.length(); i++) {
            for (int j = 1; j <= s2.length(); j++) {
                if (s3.charAt(i + j - 1) == s1.charAt(i - 1) && res[i - 1][j]) {
                    res[i][j] = true;
                } else if (s3.charAt(i + j - 1) == s2.charAt(j - 1) && res[i][j - 1]) {
                    res[i][j] = true;
                } else {
                    res[i][j] = false;
                }
            }
        }

        return res[s1.length()][s2.length()];
    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        System.out.println(solution.isInterleave("aabcc", "dbbca", "aadbbcbcac"));
        System.out.println(!solution.isInterleave("aabcc", "dbbca", "aadbbbaccc"));
        System.out.println(solution.isInterleave("", "", ""));
        System.out.println(solution.isInterleave("", "b", "b"));
    }
}
// @lc code=end

