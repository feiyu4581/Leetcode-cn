/*
 * @lc app=leetcode.cn id=45 lang=java
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
class Solution {
    public int jump(int[] nums) {
        int index = 0;
        int steps = 0;
        while (index + nums[index] < nums.length - 1) {
            int maxRange = 0;
            int range = nums[index] + index;
            for (int i = index + 1; i <= range; i++) {
                if (i + nums[i] > maxRange) {
                    maxRange = i + nums[i];
                    index = i;
                }
            }

            steps++;
        }

        return index < nums.length - 1? steps + 1 : steps;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.jump(new int[]{2, 3, 1, 1, 4}) == 2
        );

        System.out.println(
            solution.jump(new int[]{0})
        );
    }
}
// @lc code=end

