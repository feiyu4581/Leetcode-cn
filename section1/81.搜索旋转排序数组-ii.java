/*
 * @lc app=leetcode.cn id=81 lang=java
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
class Solution {
    public boolean search(int[] nums, int target) {
        int left = 0, right = nums.length - 1;

        while (left < right && nums[left] == nums[right]) {
            right--;
        }

        while (left <= right) {
            int mid = left + ((right - left) >> 1);
            if (nums[mid] == target || nums[left] == target || nums[right] == target) {
                return true;
            } else if (nums[mid] > target) {
                if (nums[left] <= nums[right] || nums[mid] < nums[left] || target >= nums[left]) {
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            } else {
                if (nums[left] <= nums[right] || nums[mid] > nums[left] || target < nums[left]) {
                    left = mid + 1;
                } else {
                    right = mid - 1;
                }
            }
        }
        return false;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.search(new int[]{2,5,6,0,0,1,2}, 0));
        System.out.println(solution.search(new int[]{3, 1, 1}, 3));
        System.out.println(!solution.search(new int[]{2,5,6,0,0,1,2}, 3));
        System.out.println(solution.search(new int[]{3, 1, 2, 2, 2}, 1));
    }
}
// @lc code=end

