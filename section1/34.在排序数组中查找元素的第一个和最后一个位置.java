import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=34 lang=java
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
class Solution {
    private int findLeft(int[] nums, int left, int right, int target) {
        while (left <= right) {
            int current = (left + right) / 2;
            if (nums[current] == target) {
                if (current == left || nums[current - 1] != target) {
                    return current;
                } else {
                    right = current - 1;
                }
            } else {
                left = current + 1;
            }
        }

        return -1;
    }

    private int findRight(int[] nums, int left, int right, int target) {
        while (left <= right) {
            int current = (left + right) / 2;
            if (nums[current] == target) {
                if (current == right || nums[current + 1] != target) {
                    return current;
                } else {
                    left = current + 1;
                }
            } else {
                right = current - 1;
            }
        }

        return -1;
    }
    public int[] searchRange(int[] nums, int target) {
        int[] resNums = new int[]{-1, -1};
        int left = 0, right = nums.length - 1;
        while (left <= right) {
            int current = (left + right) / 2;
            if (nums[current] == target) {
                resNums[0] = findLeft(nums, left, current, target);
                resNums[1] = findRight(nums, current, right, target);
                return resNums;
            } else if (nums[current] > target) {
                right = current - 1;
            } else {
                left = current + 1;
            }
        }

        return resNums;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            Arrays.toString(solution.searchRange(new int[]{5,7,7,8,8,10}, 8))
        );

        System.out.println(
            Arrays.toString(solution.searchRange(new int[]{5,7,7,8,8,10}, 6))
        );
    }
}
// @lc code=end

