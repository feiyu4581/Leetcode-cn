/*
 * @lc app=leetcode.cn id=63 lang=java
 *
 * [63] 不同路径 II
 */

// @lc code=start
class Solution {
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        int[][] res = new int[obstacleGrid.length][obstacleGrid[0].length];
        if (obstacleGrid[0][0] == 1) {
            return 0;
        }

        for (int i = 0; i < obstacleGrid.length; i++) {
            for (int j = 0; j < obstacleGrid[0].length; j++) {
                if (i == 0 && j == 0) {
                    res[i][j] = 1;
                    continue;
                }

                if (obstacleGrid[i][j] == 1) {
                    res[i][j] = 0;
                    continue;
                }

                int top = i - 1 >= 0? res[i - 1][j] : 0;
                int left = j - 1 >= 0? res[i][j - 1] : 0;

                res[i][j] =top + left;
            }
        }

        return res[res.length - 1][res[0].length - 1];
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.uniquePathsWithObstacles(new int[][]{
            new int[]{0, 0, 0},
            new int[]{0, 1, 0},
            new int[]{0, 0, 0}
        }) == 2);

        System.out.println(solution.uniquePathsWithObstacles(new int[][]{
            new int[]{0, 1},
            new int[]{0, 0}
        }) == 1);

        System.out.println(solution.uniquePathsWithObstacles(new int[][]{
            new int[]{0, 1},
            new int[]{1, 0}
        }) == 0);
    }
}
// @lc code=end

