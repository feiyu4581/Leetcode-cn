import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/*
 * @lc app=leetcode.cn id=56 lang=java
 *
 * [56] 合并区间
 */

// @lc code=start
class Solution {
    public int[][] merge(int[][] intervals) {
        Arrays.sort(intervals, (left, right) -> left[0] - right[0]);

        List<int[]> res = new ArrayList<>();
        int[] leftVal = null;
        for (int[] interval : intervals) {
            if (leftVal == null) {
                leftVal = interval;
            } else if (leftVal[1] >= interval[0]) {
                leftVal[1] = Math.max(leftVal[1], interval[1]);
            } else {
                res.add(leftVal);
                leftVal = interval;
            }
        }

        if (leftVal != null) {
            res.add(leftVal);
        }

        return res.toArray(new int[][]{});
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.print(
            solution.merge(new int[][]{
                new int[]{1, 3},
                new int[]{2, 6},
                new int[]{8, 10},
                new int[]{15,18}
            }) // [[1,6],[8,10],[15,18]]
        );

        System.out.print(
            solution.merge(new int[][]{
                new int[]{1, 4},
                new int[]{4, 5}
            }) // [[1,5]]
        );
    }
}
// @lc code=end
