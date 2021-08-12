/*
 * @lc app=leetcode.cn id=64 lang=java
 *
 * [64] 最小路径和
 */

// @lc code=start
class Solution {
    public int minPathSum(int[][] grid) {
        int[][] res = new int[grid.length][grid[0].length];
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (i == 0 && j == 0) {
                    res[i][j] = grid[0][0];
                    continue;
                }

                int top = i - 1 >= 0? res[i - 1][j] : Integer.MAX_VALUE;
                int left = j - 1 >= 0? res[i][j - 1] : Integer.MAX_VALUE;

                res[i][j] = grid[i][j] + Math.min(top, left);
            }
        }

        return res[res.length - 1][res[0].length - 1];
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.minPathSum(new int[][]{
            new int[]{1, 3, 1},
            new int[]{1, 5, 1},
            new int[]{4, 2, 1}
        }) == 7);

        System.out.println(solution.minPathSum(new int[][]{
            new int[]{1, 2, 3},
            new int[]{4, 5, 6}
        }) == 12);
    }
}
// @lc code=end

