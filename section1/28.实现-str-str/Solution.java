/*
 * @lc app=leetcode.cn id=28 lang=java
 *
 * [28] 实现 strStr()
 */

// @lc code=start
class Solution {
    public int strStr(String haystack, String needle) {
        int needleLength = needle.length();
        if (needleLength == 0) {
            return 0;
        }

        for (int i = 0; i <= haystack.length() - needleLength; i++) {
            if (haystack.substring(i, i + needleLength).equals(needle)) {
                return i;
            }
        }

        return -1;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.strStr("hello", "ll") == 2
        );

        System.out.println(
            solution.strStr("aaaaa", "bba") == -1
        );

        System.out.println(
            solution.strStr("a", "a") == 0
        );
    }
}
// @lc code=end

