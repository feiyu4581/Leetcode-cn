import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

/*
 * @lc app=leetcode.cn id=39 lang=java
 *
 * [39] 组合总和
 */

// @lc code=start
class Solution {

    public void recursiveCombinationSum(List<List<Integer>> resNums, int[] candidates, int target, List<Integer> prefix, int current) {
        if (target == 0) {
            List<Integer> copyNums = Arrays.asList(new Integer[prefix.size()]);
            Collections.copy(copyNums, prefix);
            resNums.add(copyNums);
        } else if (current >= candidates.length || candidates[current] > target) {
            return;
        } else {
            prefix.add(candidates[current]);
            recursiveCombinationSum(resNums, candidates, target - candidates[current], prefix, current);

            prefix.remove(prefix.size() - 1);
            recursiveCombinationSum(resNums, candidates, target, prefix, current + 1);
        }
    }

    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> resNums = new ArrayList<>();
        Arrays.sort(candidates);
        recursiveCombinationSum(resNums, candidates, target, new ArrayList<>(), 0);
        return resNums;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.combinationSum(new int[]{2, 3, 6, 7}, 7)
        );

        System.out.println(
            solution.combinationSum(new int[]{2, 3, 5}, 8)
        );
    }
}
// @lc code=end

