/*
 * @lc app=leetcode.cn id=87 lang=java
 *
 * [87] 扰乱字符串
 */

import java.util.HashMap;
import java.util.Map;

// @lc code=start
class Solution {
    char[] char1;
    char[] char2;
    Map<String, Boolean> caches;

    public boolean isScramble(String s1, String s2) {
        this.char1 = s1.toCharArray();
        this.char2 = s2.toCharArray();
        this.caches = new HashMap<>();

        return isScrambleRecursive(0, s1.length(), 0, s2.length());
    }

    private boolean isScrambleRecursive(int left1, int right1, int left2, int right2) {
        String cacheStr = left1 + "/" + right1 + "/" + left2 + "/" + right2;
        if (!caches.containsKey(cacheStr)) {
            boolean res = false;
            if (right2 - left2 != right1 - left1 || !compareWords(left1, right1, left2, right2)) {
                res = false;
            } else if (right1 - left1 == 1) {
                res = true;
            } else {
                for (int i = left1 + 1; i < right1; i++) {
                    if (isScrambleRecursive(left1, i, left2, i + left2 - left1) && isScrambleRecursive(i, right1, i + right2 - right1, right2)) {
                        res = true;
                        break;
                    }

                    if (isScrambleRecursive(left1, i, right2 - i + left1, right2) && isScrambleRecursive(i, right1, left2, right1 - i + left2)) {
                        res = true;
                        break;
                    }
                }
            }

            caches.put(cacheStr, res);
        }

        return caches.get(cacheStr);
    }

    private boolean compareWords(int left1, int right1, int left2, int right2) {
        Map<Character, Integer> leftWords = new HashMap<>();
        Map<Character, Integer> rightWords = new HashMap<>();

        for (int i = left1; i < right1; i++) {
            leftWords.putIfAbsent(char1[i], 0);
            leftWords.put(char1[i], leftWords.get(char1[i]) + 1);
        }

        for (int j = left2; j < right2; j++) {
            rightWords.putIfAbsent(char2[j], 0);
            rightWords.put(char2[j], rightWords.get(char2[j]) + 1);

            if (rightWords.get(char2[j]) > leftWords.getOrDefault(char2[j], 0)) {
                return false;
            }
        }

        return true;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.isScramble("great", "rgeat"));
        System.out.println(!solution.isScramble("abcde", "caebd"));
        System.out.println(solution.isScramble("a", "a"));
        System.out.println(solution.isScramble("abc", "cab"));
        System.out.println(!solution.isScramble("dbdac", "abcdd"));
    }
}
// @lc code=end

