import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

/*
 * @lc app=leetcode.cn id=18 lang=java
 *
 * [18] 四数之和
 */

// @lc code=start
class Solution {
    public List<List<Integer>> fourSum(final int[] nums, final int target) {
        Arrays.sort(nums);
        List<List<Integer>> resNums = new ArrayList<>();
        HashMap<Integer, List<Integer>> numsMaps = new HashMap<>();
        HashMap<List<Integer>, Boolean> cache = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            numsMaps.putIfAbsent(nums[i], new ArrayList<>());
            numsMaps.get(nums[i]).add(i);
        }
        for (int i = 0; i < nums.length; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                for (int x = j + 1; x < nums.length; x++) {
                    int current = target - nums[i] - nums[j] - nums[x];
                    if (numsMaps.containsKey(current)) {
                        for (int index : numsMaps.get(current)) {
                            if (index > x) {
                                List<Integer> res = new ArrayList<>();
                                res.add(nums[i]);
                                res.add(nums[j]);
                                res.add(nums[x]);
                                res.add(current);

                                if (!cache.containsKey(res)) {
                                    resNums.add(res);
                                }

                                cache.put(res, true);
                                break;
                            }
                        }
                    }
                }
            }
        }

        return resNums;
    }

    public static void main(final String[] args) {
        final Solution solution = new Solution();
        System.out.println(
            solution.fourSum(new int[]{1, 0, -1, 0, -2, 2}, 0)
        );
    }
}
// @lc code=end

