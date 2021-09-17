import java.util.HashMap;

/*
 * @lc app=leetcode.cn id=44 lang=java
 *
 * [44] 通配符匹配
 */

// @lc code=start
class Solution {
    private boolean recursiveIsMatch(String s, String p, int left, int right, HashMap<String, Boolean> backups) {
        String key = String.format("%s-%s", left, right);
        if (!backups.containsKey(key)) {
            boolean res = false;
            if (left == s.length() && right == p.length()) {
                res = true;
            } else if (right >= p.length()) {
                res = false;
            } else if (left >= s.length()) {
                if (p.charAt(right) == '*') {
                    res = recursiveIsMatch(s, p, left, right + 1, backups);
                } else {
                    res = false;
                }
            } else if (s.charAt(left) == p.charAt(right) || p.charAt(right) == '?') {
                res = recursiveIsMatch(s, p, left + 1, right + 1, backups);
            } else if (p.charAt(right) == '*') {
                res = recursiveIsMatch(s, p, left + 1, right, backups) || recursiveIsMatch(s, p, left, right + 1, backups);
            }
            backups.put(key, res);
        }

        return backups.get(key);
    }

    public boolean isMatch(String s, String p) {
        return recursiveIsMatch(s, p, 0, 0, new HashMap<>());
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.isMatch(
                "aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba",
                "a*******b")
        );

        System.out.println(
            solution.isMatch(
                "abbabbbaabaaabbbbbabbabbabbbabbaaabbbababbabaaabbab",
                "*aabb***aa**a******aa*"
            )
        );

        System.out.println(
            solution.isMatch("aa", "a") == false
        );

        System.out.println(
            solution.isMatch("aa", "*") == true
        );

        System.out.println(
            solution.isMatch("cb", "?a") == false
        );

        System.out.println(
            solution.isMatch("adceb", "*a*b") == true
        );

        System.out.println(
            solution.isMatch("acdcb", "a*c?b") == false
        );
    }
}
// @lc code=end

