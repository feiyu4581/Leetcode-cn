import java.util.*;

/*
 * @lc app=leetcode.cn id=78 lang=java
 *
 * [78] 子集
 */


// @lc code=start
class Solution {
    Set<List<Integer>> records;
    int[] nums;

    public List<List<Integer>> subsets(int[] nums) {
        this.records = new HashSet<>();
        this.nums = nums;

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
            subsetsRecursive(current, start + 1, length);
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.subsets(new int[]{1, 2, 3}));  // [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
        System.out.println(solution.subsets(new int[]{0}));  // [[],[0]]
    }
}
// @lc code=end

