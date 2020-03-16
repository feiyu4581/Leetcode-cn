/*
 * @lc app=leetcode.cn id=5 lang=java
 *
 * [5] 最长回文子串
 */

// @lc code=start
class Solution {
    public String longestPalindrome(String s) {
        int maxLength = -1;
        String maxString = "";
        for (int i = 0; i < s.length(); i++)  {
            int subMaxLength = getSubMaxLength(s, i, i);
            if (maxLength < subMaxLength) {
                maxLength = subMaxLength;
                maxString = s.substring(i - subMaxLength, i + subMaxLength + 1);
            }

            if (i + 1 < s.length() && s.charAt(i) == s.charAt(i + 1)) {
                subMaxLength = getSubMaxLength(s, i, i + 1);
                if (maxLength <= subMaxLength && i + subMaxLength + 2 <= s.length()) {
                    maxLength = subMaxLength;
                    maxString = s.substring(i - subMaxLength, i + subMaxLength + 2);
                }
            }
        }

        return maxString;
    }

    private int getSubMaxLength(String s, int begain, int end) {
        boolean isPalidrome = true;
        int maxLength = 0;
        for (int i = 0; i < s.length(); i++) {
            if (begain - i >= 0 && end + i < s.length()) {
                isPalidrome = s.charAt(begain - i) == s.charAt(end + i);
                if (isPalidrome && i > maxLength) {
                    maxLength = i;
                }
            } else {
                break;
            }

            if (!isPalidrome) {
                break;
            }
        }

        return maxLength;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.longestPalindrome("babad")
        );

        System.out.println(
            solution.longestPalindrome("cbbd")
        );

        System.out.println(
            solution.longestPalindrome("ccc")
        );

        System.out.println(
            solution.longestPalindrome("abcda")
        );

        System.out.println(
            solution.longestPalindrome("tattarrattat")
        );
    }
}
// @lc code=end

