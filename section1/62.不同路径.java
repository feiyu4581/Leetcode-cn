/*
 * @lc app=leetcode.cn id=62 lang=java
 *
 * [62] 不同路径
 */

// @lc code=start
class Solution {
    public int uniquePaths(int m, int n) {
        int[][] res = new int[m][n];

        // res[0][0] = 1;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int top = i - 1 >= 0? res[i - 1][j] : 0;
                int left = j - 1 >= 0? res[i][j - 1] : 0;
                
                res[i][j] = Math.max(top + left, 1);
            }
        }

        return res[m - 1][n - 1];
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.uniquePaths(3, 7) == 28);
        System.out.println(solution.uniquePaths(3, 2) == 3);
        System.out.println(solution.uniquePaths(7, 3) == 28);
        System.out.println(solution.uniquePaths(3, 3) == 6);
    }
}
// @lc code=end

