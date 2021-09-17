import java.util.Arrays;
import java.util.HashMap;

/*
 * @lc app=leetcode.cn id=1 lang=java
 *
 * [1] 两数之和
 */

// @lc code=start
class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> numIndexMaps = new HashMap<Integer, Integer>();
        for (int index = 0; index < nums.length; index++) {
            if (numIndexMaps.containsKey(target - nums[index])) {
                return new int[]{numIndexMaps.get(target - nums[index]), index};
            }

            numIndexMaps.put(nums[index], index);
        }

        return null;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            Arrays.toString(solution.twoSum(new int[] {2, 7, 11, 15}, 9))
        );
    }
}
// @lc code=end

