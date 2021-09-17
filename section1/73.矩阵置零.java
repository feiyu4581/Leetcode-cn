import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=73 lang=java
 *
 * [73] 矩阵置零
 */

// @lc code=start
class Solution {
    public void setZeroes(int[][] matrix) {
        boolean firstRowZeroe = false;
        boolean firstColumnZerore = false;
        for (int j = 0; j < matrix[0].length; j++) {
            if (matrix[0][j] == 0) {
                firstRowZeroe = true;
            }
        }

        for (int i = 0; i < matrix.length; i++) {
            if (matrix[i][0] == 0) {
                firstColumnZerore = true;
            }
        }
        

        for (int i = 1; i < matrix.length; i++) {
            for (int j = 1; j < matrix[0].length; j++) {
                if (matrix[i][j] == 0) {
                    matrix[0][j] = 0;
                    matrix[i][0] = 0;
                }
            }
        }

        for (int i = 1; i < matrix.length; i++) {
            for (int j = 1; j < matrix[0].length; j++) {
                if (matrix[0][j] == 0 || matrix[i][0] == 0) {
                    matrix[i][j] = 0;
                }
            }
        }

        if (firstRowZeroe) {
            for (int j = 0; j < matrix[0].length; j++) {
                matrix[0][j] = 0;
            }
        }

        if (firstColumnZerore) {
            for (int i = 0; i < matrix.length; i++) {
                matrix[i][0] = 0;
            }
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        // solution.setZeroes(new int[][]{
        //     new int[]{1, 1, 1},
        //     new int[]{1, 0, 1},
        //     new int[]{1, 1, 1}
        // });  // [[1,0,1],[0,0,0],[1,0,1]]

        solution.setZeroes(new int[][]{
            new int[]{0, 1, 2, 0},
            new int[]{3, 4, 5, 2},
            new int[]{1, 3, 1, 5}
        }); // [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

        solution.setZeroes(new int[][]{
            new int[]{0},
            new int[]{1},
        });
    }
}
// @lc code=end

