import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;

/*
 * @lc app=leetcode.cn id=47 lang=java
 *
 * [47] 全排列 II
 */

// @lc code=start
class Solution {
    public void recursivePermuteUnique(int[] nums, HashSet<Integer> prefixSet, List<Integer> prefix, HashSet<List<Integer>> res) {
        if (prefix.size() == nums.length) {
            List<Integer> current = new ArrayList<>(prefix);
            res.add(current);
        } else {
            for (int i = 0; i < nums.length; i++) {
                if(!prefixSet.contains(i)) {
                    prefixSet.add(i);
                    prefix.add(nums[i]);
                    recursivePermuteUnique(nums, prefixSet, prefix, res);
                    prefixSet.remove(i);
                    prefix.remove(prefix.size() - 1);
                }
            }
        }
    }
    public List<List<Integer>> permuteUnique(int[] nums) {
        HashSet<List<Integer>> res = new HashSet<>();
        recursivePermuteUnique(nums, new HashSet<>(), new ArrayList<>(), res);

        return new ArrayList<>(res);
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.permuteUnique(new int[]{1, 1, 2})
        );
    }
}
// @lc code=end


// @lc code=end

