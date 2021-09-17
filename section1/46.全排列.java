import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;

/*
 * @lc app=leetcode.cn id=46 lang=java
 *
 * [46] 全排列
 */

// @lc code=start
class Solution {
    public void recursivePermute(int[] nums, HashSet<Integer> prefixSet, List<Integer> prefix, List<List<Integer>> res) {
        if (prefix.size() == nums.length) {
            List<Integer> current = new ArrayList<>(prefix);
            res.add(current);
        } else {
            for (int i = 0; i < nums.length; i++) {
                if(!prefixSet.contains(nums[i])) {
                    prefixSet.add(nums[i]);
                    prefix.add(nums[i]);
                    recursivePermute(nums, prefixSet, prefix, res);
                    prefixSet.remove(nums[i]);
                    prefix.remove(prefix.size() - 1);
                }
            }
        }
    }
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        recursivePermute(nums, new HashSet<>(), new ArrayList<>(), res);

        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.permute(new int[]{1, 2, 3})
        );
    }
}
// @lc code=end

