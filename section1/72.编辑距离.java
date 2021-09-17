import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=72 lang=java
 *
 * [72] 编辑距离
 */

// @lc code=start
class Solution {
    public int minDistance(String word1, String word2) {
        if (word1.length() == 0 || word2.length() == 0) {
            return Math.max(word1.length(), word2.length());
        }

        int[][] distances = new int[word1.length()][word2.length()];
        
        int last, top, bottom;
        for (int i = 0; i < word1.length(); i++) {
            for (int j = 0; j < word2.length(); j++) {
                top = i > 0? distances[i - 1][j] + 1 : Integer.MAX_VALUE;
                bottom = j > 0? distances[i][j - 1] + 1 : Integer.MAX_VALUE;

                if (i > 0 && j > 0) {
                    last = distances[i - 1][j - 1];
                } else {
                    last = Math.max(i, j);
                }
                if (word1.charAt(i) != word2.charAt(j)) {
                    last++;
                }

                distances[i][j] = Math.min(Math.min(top, bottom), last);
            }
        }

        return distances[word1.length() - 1][word2.length() - 1];
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.minDistance("horse", "ros") == 3);
        System.out.println(solution.minDistance("intention", "execution") == 5);
        System.out.println(solution.minDistance("sea", "eat") == 2);
    }
}
// @lc code=end

