/*
 * @lc app=leetcode.cn id=33 lang=java
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
class Solution {
    public int search(int[] nums, int target) {
        int left = 0, right = nums.length - 1;
        while (left <= right) {
            int current = (left + right) / 2;
            if (nums[current] == target) {
                return current;
            }

            if (nums[current] >= nums[left] && nums[current] <= nums[right]) {
                if(nums[current] > target) {
                    right = current - 1;
                } else {
                    left = current + 1;
                }
            } else if (nums[current] >= nums[left]) {
                if (nums[current] > target) {
                    if (target >= nums[left]) {
                        right = current - 1;
                    } else {
                        left = current + 1;
                    }
                } else {
                    left = current + 1;
                }
            } else {
                if (nums[current] > target) {
                    right = current - 1;
                } else if (target <= nums[right]) {
                    left = current + 1;
                } else {
                    right = current - 1;
                }
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.search(new int[]{4,5,6,7,0,1,2}, 0)
        );

        System.out.println(
            solution.search(new int[]{4,5,6,7,0,1,2}, 3)
        );

        System.out.println(
            solution.search(new int[]{3, 1}, 1)
        );
    }
}
// @lc code=end

