/*
 * @lc app=leetcode.cn id=27 lang=java
 *
 * [27] 移除元素
 */

// @lc code=start
class Solution {
    public int removeElement(int[] nums, int val) {
        int current = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != val) {
                nums[current++] = nums[i];
            }
        }

        return current;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.removeElement(new int[]{3, 2, 2, 3}, 3) == 2
        );

        System.out.println(
            solution.removeElement(new int[]{0,1,2,2,3,0,4,2}, 2) == 5
        );
    }
}
// @lc code=end

