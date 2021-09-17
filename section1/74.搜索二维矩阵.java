/*
 * @lc app=leetcode.cn id=74 lang=java
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        int leftRow = 0, leftColumn = 0;
        int rightRow = matrix.length - 1, rightColumn = matrix[0].length - 1;

        while (leftRow <= rightRow && leftColumn <= rightColumn) {
            int midRow = leftRow + ((rightRow - leftRow) >> 1);
            int midColumn = leftColumn + ((rightColumn - leftColumn) >> 1);

            int midValue = matrix[midRow][midColumn];
            if (midValue == target) {
                return true;
            } else if (midValue > target) {
                if (midColumn > 0) {
                    rightColumn = midColumn - 1;
                } else {
                    rightRow = midRow - 1;
                    rightColumn = matrix[0].length - 1;
                }
            } else {
                if (midColumn < matrix[0].length - 1) {
                    leftColumn = midColumn + 1;
                } else {
                    leftRow = leftRow + 1;
                    leftColumn = 0;
                }
            }
        }

        return false;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
                solution.searchMatrix(new int[][]{
                        new int[]{1, 3, 4, 7},
                        new int[]{10, 11, 16, 20},
                        new int[]{23, 30, 34, 60}
                }, 3)
        );

        System.out.println(
                !solution.searchMatrix(new int[][]{
                        new int[]{1, 3, 4, 7},
                        new int[]{10, 11, 16, 20},
                        new int[]{23, 30, 34, 60}
                }, 13)
        );
    }
}
// @lc code=end

