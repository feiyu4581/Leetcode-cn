import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=42 lang=java
 *
 * [42] 接雨水
 */

// @lc code=start
class Solution {
    private int recursiveTrap(int[] height, int start) {
        int resNum = 0;
        if (start >= height.length - 1) {
            return resNum;
        }

        int left = start;
        for (int i = start + 1; i < height.length; i++) {
            if (height[i] >= height[left]) {
                int diff = Math.min(height[i], height[left]);
                for (int j = left + 1; j < i; j++) {
                    resNum += diff - height[j];
                }

                left = i;
            }
        }

        resNum += Math.max(recursiveTrap(height, left + 1), recursiveReverseTrap(height, left, height.length - 1));
        return resNum;
    }

    private int recursiveReverseTrap(int[] height, int start, int end) {
        int resNum = 0;
        if (start >= end) {
            return resNum;
        }

        int left = end;
        for (int i = end - 1; i >= start; i--) {
            if (height[i] >= height[left]) {
                int diff = Math.min(height[i], height[left]);
                for (int j = left - 1; j > i; j--) {
                    resNum += diff - height[j];
                }

                left = i;
            }
        }

        resNum += recursiveReverseTrap(height, start, left - 1);
        return resNum;
    }

    public int trap(int[] height) {
        return recursiveTrap(height, 0);
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.trap(new int[]{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}) == 6
        );

        System.out.println(
            solution.trap(new int[]{4, 2, 3}) == 1
        );

        System.out.println(
            solution.trap(new int[]{6,4,2,0,3,2,0,3,1,4,5,3,2,7,5,3,0,1,2,1,3,4,6,8,1,3}) == 83
        );
    }
}
// @lc code=end

