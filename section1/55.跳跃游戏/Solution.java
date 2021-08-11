/*
 * @lc app=leetcode.cn id=55 lang=java
 *
 * [55] 跳跃游戏
 */

// @lc code=start
class Solution {
    public boolean canJump(int[] nums) {
        int index = 0;
        while (index + nums[index] < nums.length - 1) {
            if (nums[index] == 0) {
                return false;
            }
            int maxVals = 0, maxIndex = 0;
            for (int i = 1; i <= nums[index]; i++) {
                if (index + i + nums[index + i] >= maxVals) {
                    maxIndex = index + i;
                    maxVals = maxIndex + nums[maxIndex];
                }
            }

            index = maxIndex;
        }

        return true;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.print(
            solution.canJump(new int[]{2, 3, 1, 1, 4}) == true
        );

        System.out.print(
            solution.canJump(new int[]{3, 2, 1, 0, 4}) == false
        );
    }
}
// @lc code=end

