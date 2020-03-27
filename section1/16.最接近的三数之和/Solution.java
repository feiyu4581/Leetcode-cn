import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=16 lang=java
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
class Solution {
    public int threeSumClosest(int[] nums, int target) {
        int resNum = Integer.MAX_VALUE;
        Arrays.sort(nums);
        for(int i = 0; i < nums.length; i++) {
            int left = i + 1, right = nums.length - 1;
            while (left < right) {
                int diff = target - nums[left] - nums[right] - nums[i];
                if (diff == 0) {
                    return target;
                } else if (diff > 0) {
                    left += 1;
                } else {
                    right -= 1;
                }

                if (Math.abs(diff) < Math.abs(resNum)) {
                    resNum = diff;
                }
            }
        }

        return target - resNum;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.threeSumClosest(new int[]{-1, 2, 1, -4}, 1) == 2
        );
    }
}
// @lc code=end

