/*
 * @lc app=leetcode.cn id=10 lang=java
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
class Solution {
    private boolean isSubMatch(String s, int i, String p, int j) {
        if (i == s.length() && j == p.length()) {
            return true;
        }

        if (j >= p.length()) {
            return false;
        }

        boolean existsStar = j + 1 < p.length() && p.charAt(j + 1) == '*';
        if (existsStar && isSubMatch(s, i, p, j + 2)) {
            return true;
        }

        if (i < s.length()) {
            char current = p.charAt(j);
            boolean match = current == '.' || current == s.charAt(i);

            if (existsStar && match) {
                return isSubMatch(s, i + 1, p, j);
            }

            return match && isSubMatch(s, i + 1, p, j + 1);
        }
        return false;
    }

    public boolean isMatch(String s, String p) {
        return isSubMatch(s, 0, p, 0);
    }

    public static void main(String[] args) {
        Solution solution = new Solution();

        System.out.println(
            solution.isMatch("aa", "a") == false
        );

        System.out.println(
            solution.isMatch("aa", "a*") == true
        );

        System.out.println(
            solution.isMatch("ab", ".*") == true
        );

        System.out.println(
            solution.isMatch("aab", "c*a*b") == true
        );

        System.out.println(
            solution.isMatch("mississippi", "mis*is*p*.") == false
        );
    }
}
// @lc code=end

