/*
 * @lc app=leetcode.cn id=11 lang=java
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
class Solution {
    public int maxArea(int[] height) {
        int maxResult = 0;
        int lMin = 0, rMax = height.length - 1;
        while (lMin < rMax) {
            int current = (rMax - lMin) * Math.min(height[rMax], height[lMin]);
            if (current > maxResult) {
                maxResult = current;
            }

            if (height[rMax] < height[lMin]) {
                rMax -= 1;
            } else {
                lMin += 1;
            }
        }
        return maxResult;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.maxArea(new int[]{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49
        );
    }
}
// @lc code=end

