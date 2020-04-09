/*
 * @lc app=leetcode.cn id=26 lang=java
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
class Solution {
    public int removeDuplicates(int[] nums) {
        int current = 0;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] != nums[current]) {
                nums[++current] = nums[i];
            }
        }

        return current + 1;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.removeDuplicates(new int[]{1, 1, 2})
        );

        System.out.println(
            solution.removeDuplicates(new int[]{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
        );
    }
}
// @lc code=end

