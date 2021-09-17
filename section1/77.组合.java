import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/*
 * @lc app=leetcode.cn id=77 lang=java
 *
 * [77] 组合
 */

// @lc code=start
class Solution {
    List<List<Integer>> records;
    int upper;
    int target;

    private void combileRecursive(int[] current, int start, int length) {
        if (length == target) {
            List<Integer> record = new ArrayList<>();
            for (int i : current) {
                record.add(i);
            }
            records.add(record);
        } else {
            current[length] = start;
            if (start <= this.upper) {
                combileRecursive(current, start + 1, length + 1);
                combileRecursive(current, start + 1, length);
            }
        }
    }

    public List<List<Integer>> combine(int n, int k) {
        this.records = new ArrayList<>();
        this.upper = n;
        this.target = k;

        combileRecursive(new int[k], 1, 0);
        return this.records;
    }
    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.combine(4, 2));
        // [
        //     [2,4],
        //     [3,4],
        //     [2,3],
        //     [1,2],
        //     [1,3],
        //     [1,4],
        // ]
        System.out.println(solution.combine(1, 1)); // [[1]]
    }
}
// @lc code=end

