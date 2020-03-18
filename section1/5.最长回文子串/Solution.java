/*
 * @lc app=leetcode.cn id=5 lang=java
 *
 * [5] 最长回文子串
 */

// @lc code=start
class Solution {
    public String longestPalindromeWithRange(String s) {
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

    public String longestPalindrome(String s) {
        int[][] palidromeMaps = new int[s.length()][s.length()];

        String maxPalidrome = "";
        int maxLength = 0;

        for (int i = 0; i < s.length(); i++) {
            palidromeMaps[i][i] = 1;
            if (maxLength < 1) {
                maxLength = 1;
                maxPalidrome = s.substring(i, i + 1);
            }
        }

        for (int i = 1; i < s.length(); i++) {
            for (int j = 0; j < s.length(); j++) {
                if (j + i < s.length() && s.charAt(j) == s.charAt(j + i)) {
                    if (i > 2 && palidromeMaps[j + 1][j + i - 1] > 0) {
                        palidromeMaps[j][j + i] = palidromeMaps[j + 1][j + i - 1] + 2;
                    } else if (i <= 2) {
                        palidromeMaps[j][j + i] = i + 1;
                    }

                    if (palidromeMaps[j][j + i] > maxLength) {
                        maxLength = palidromeMaps[j][j + i];
                        maxPalidrome = s.substring(j, j + i + 1);
                    }
                }
            }
        }
        return maxPalidrome;
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

