/*
 * @lc app=leetcode.cn id=80 lang=java
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
class Solution {
    public int removeDuplicates(int[] nums) {
        int counts = 1;
        int current = 0;
        for (int i = 0; i < nums.length; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) {
                counts++;
            } else {
                counts = 1;
            }

            if (counts > 2) {
                continue;
            }

            nums[current++] = nums[i];
        }

        return current;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
                solution.removeDuplicates(new int[]{1,1,1,2,2,3}) == 5
        );

        System.out.println(
                solution.removeDuplicates(new int[]{0,0,1,1,1,1,2,3,3}) == 7
        );
    }
}
// @lc code=end

