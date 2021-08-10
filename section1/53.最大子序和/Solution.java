/*
 * @lc app=leetcode.cn id=53 lang=java
 *
 * [53] 最大子序和
 */

// @lc code=start
class Solution {
    public int maxSubArray(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }

        int maxVals = nums[0];
        int leftVal = maxVals;

        for (int i = 1; i < nums.length; i++) {
            leftVal = Math.max(nums[i], nums[i] + leftVal);
            if (leftVal > maxVals) {
                maxVals = leftVal;
            }
        }

        return maxVals;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.maxSubArray(new int[]{-2,1,-3,4,-1,2,1,-5,4}) == 6
        );

        System.out.println(
            solution.maxSubArray(new int[]{1}) == 1
        );

        System.out.println(
            solution.maxSubArray(new int[]{0}) == 0
        );

        System.out.println(
            solution.maxSubArray(new int[]{-10000}) == -10000
        );
    }
}
// @lc code=end

