import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=41 lang=java
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
class Solution {
    public int firstMissingPositive(int[] nums) {
        boolean findOne = false;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 1) {
                findOne = true;
                break;
            }
        }

        if (!findOne) {
            return 1;
        }

        if (nums.length == 1) {
            return 2;
        }

        for (int i = 0; i < nums.length; i++) {
            if (nums[i] <= 0) {
                nums[i] = 1;
            }
        }

        for (int i = 0; i < nums.length; i++) {
            if (Math.abs(nums[i]) <= nums.length) {
                nums[Math.abs(nums[i]) - 1] = -Math.abs(nums[Math.abs(nums[i]) - 1]);
            }
        }

        for (int i = 0; i < nums.length; i++) {
            if (nums[i] > 0) {
                return i + 1;
            }
        }

        return nums.length + 1;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.firstMissingPositive(new int[]{1, 2, 0}) == 3
        );

        System.out.println(
            solution.firstMissingPositive(new int[]{3, 4, -1, 1}) == 2
        );

        System.out.println(
            solution.firstMissingPositive(new int[]{7, 8, 9, 11, 12}) == 1
        );

        System.out.println(
            solution.firstMissingPositive(new int[]{2, 1}) == 3
        );
    }
}
// @lc code=end

