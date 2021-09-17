/*
 * @lc app=leetcode.cn id=35 lang=java
 *
 * [35] 搜索插入位置
 */

// @lc code=start
class Solution {
    public int searchInsert(int[] nums, int target) {
        int left = 0, right = nums.length - 1;
        while (left < right) {
            int current = (left + right) / 2;
            if (nums[current] == target) {
                return current;
            } else if (nums[current] > target) {
                right = current - 1;
            } else {
                left = current + 1;
            }
        }

        return nums[left] >= target? left : left + 1;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.searchInsert(new int[]{1, 3, 5, 6}, 5) == 2
        );

        System.out.println(
            solution.searchInsert(new int[]{1, 3, 5, 6}, 2) == 1
        );

        System.out.println(
            solution.searchInsert(new int[]{1, 3, 5, 6}, 7) == 4
        );

        System.out.println(
            solution.searchInsert(new int[]{1, 3, 5, 6}, 0) == 0
        );

        System.out.println(
            solution.searchInsert(new int[]{1}, 1) == 0
        );
    }
}
// @lc code=end

