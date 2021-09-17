import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

/*
 * @lc app=leetcode.cn id=15 lang=java
 *
 * [15] 三数之和
 */

// @lc code=start
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);

        List<List<Integer>> res = new ArrayList<>();
        HashMap<List<Integer>, Boolean> cache = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            HashMap<Integer, Integer> numMaps = new HashMap<>();
            int target = -nums[i];
            for (int j = i + 1; j < nums.length; j++) {
                if (numMaps.containsKey(target - nums[j]) && numMaps.get(target - nums[j]) > i) {
                    List<Integer> targetNum = new ArrayList<>();
                    targetNum.add(nums[i]);
                    targetNum.add(target - nums[j]);
                    targetNum.add(nums[j]);

                    if (!cache.containsKey(targetNum)) {
                        res.add(targetNum);
                    }

                    cache.put(targetNum, true);
                } else {
                    numMaps.put(nums[j], j);
                }
            }
        }

        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.threeSum(new int[]{-1, 0, 1, 2, -1, -4})
        );
    }
}
// @lc code=end

