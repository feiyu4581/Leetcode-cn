import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=48 lang=java
 *
 * [48] 旋转图像
 */

// @lc code=start
class Solution {
    public void rotate(int[][] matrix) {
        int length = matrix.length / 2;
        int index = 0;
        while (index < length) {
            for (int i = index; i < matrix.length - index - 1; i++) {
                int temp = matrix[index][i];

                matrix[index][i] = matrix[matrix.length - i - 1][index];
                matrix[matrix.length - i - 1][index] = matrix[matrix.length - index - 1][matrix.length - i - 1];
                matrix[matrix.length - index - 1][matrix.length - i - 1] = matrix[i][matrix.length - index - 1];
                matrix[i][matrix.length - index - 1] = temp;

            }

            index++;
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[][] nums = new int[][]{
            new int[]{1, 2, 3},
            new int[]{4, 5, 6},
            new int[]{7, 8, 9}
        };
        solution.rotate(nums);
        for (int i = 0; i < nums.length; i++) {
            System.out.println(Arrays.toString(nums[i]));
        }

        int[][] nums2 = new int[][]{
            new int[]{5, 1, 9, 11},
            new int[]{2, 4, 8, 10},
            new int[]{13, 3, 6, 7},
            new int[]{15, 14, 12, 16}
        };
        solution.rotate(nums2);
        for (int i = 0; i < nums2.length; i++) {
            System.out.println(Arrays.toString(nums2[i]));
        }
    }
}
// @lc code=end

