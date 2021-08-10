import java.util.ArrayList;
import java.util.List;


/*
 * @lc app=leetcode.cn id=54 lang=java
 *
 */

// @lc code=start
class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        int numSize = matrix.length * matrix[0].length;
        List<Integer> res = new ArrayList<>(numSize);

        int step = 0, row = 0, col = 0;
        while (numSize >= 1) {
            for (row=step; col < matrix[0].length - step && numSize >= 1; col++) {
                res.add(matrix[row][col]);
                numSize--;
            }

            row++;

            for (col=matrix[0].length - step - 1; row < matrix.length - step && numSize >= 1 ; row++) {
                res.add(matrix[row][col]);
                numSize--;
            }

            col--;

            for (row=matrix.length - step - 1; col >= step && numSize >= 1; col--) {
                res.add(matrix[row][col]);
                numSize--;
            }

            row--;

            step += 1;
            for (col=step - 1; row >= step && numSize >= 1; row--) {
                res.add(matrix[row][col]);
                numSize--;
            }

            col++;
        }

        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.spiralOrder(new int[][]{
                new int[]{1, 2, 3},
                new int[]{4, 5, 6},
                new int[]{7, 8, 9}
            })
        );

        System.out.println(
            solution.spiralOrder(new int[][]{
                new int[]{1, 2, 3, 4},
                new int[]{5, 6, 7, 8},
                new int[]{9, 10, 11, 12}
            })
        );
    }
}
// @lc code=end

