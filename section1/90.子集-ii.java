import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Set;

import java.util.HashSet;

/*
 * @lc app=leetcode.cn id=90 lang=java
 *
 * [90] 子集 II
 */

// @lc code=start
class Solution {
    Set<List<Integer>> records;
    int[] nums;

    public List<List<Integer>> subsetsWithDup(int[] nums) {
        this.records = new HashSet<>();
        this.nums = nums;

        Arrays.sort(this.nums);
        subsetsRecursive(new int[nums.length], 0, 0);

        return new ArrayList<>(this.records);
    }

    private void subsetsRecursive(int[] current, int start, int length) {
        List<Integer> record = new ArrayList<>();
        for (int i = 0; i < length; i++) {
            record.add(current[i]);
        }

        records.add(record);

        if (length >= nums.length) {
            return;
        }

        if (start < nums.length) {
            current[length] = nums[start];
            subsetsRecursive(current, start + 1, length + 1);
            while (start < nums.length - 1 && nums[start] == nums[start + 1]) {
                start++;
            }

            subsetsRecursive(current, start + 1, length);
        }
    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        // [[],[1],[1,1],[1,1,2],[1,1,2,2],[1,2],[1,2,2],[2],[2,2]]
        System.out.println(solution.subsetsWithDup(new int[]{1,1,2,2}));
        // [0,1,3,2]
        System.out.println(solution.subsetsWithDup(new int[]{1, 2, 2}));
        System.out.println(solution.subsetsWithDup(new int[]{0}));
    }
}
// @lc code=end

