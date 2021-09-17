import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=59 lang=java
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
class Solution {
    public int[][] generateMatrix(int n) {
        int[][] res = new int[n][n];

        int step = 0, row = 0, col = 0, numSize = n * n, index = 1;
        while (index <= numSize) {
            for (row=step; col < n - step && index <= numSize; col++) {
                res[row][col] = index++;
            }

            row++;

            for (col=n - step - 1; row < n - step && index <= numSize; row++) {
                res[row][col] = index++;
            }

            col--;

            for (row=n - step - 1; col >= step && index <= numSize; col--) {
                res[row][col] = index++;
            }

            row--;

            step += 1;
            for (col=step - 1; row >= step && index <= numSize; row--) {
                res[row][col] = index++;
            }

            col++;
        }

        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[][] res = solution.generateMatrix(3);
        for (int[] num : res) {
            System.out.println(Arrays.toString(num));
        }

        System.out.println(Arrays.toString(solution.generateMatrix(1)[0]));
    }
}
// @lc code=end

