import java.util.List;
import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=57 lang=java
 *
 * [57] 插入区间
 */

// @lc code=start
class Solution {
    public int[][] insert(int[][] intervals, int[] newInterval) {
        List<int[]> res = new ArrayList<>();
        if (intervals.length == 0) {
            res.add(newInterval);
            return res.toArray(new int[][]{});
        }

        int[] leftVal = intervals[0][0] > newInterval[0]? newInterval : null;
        for (int[] interval : intervals) {
            if (leftVal != null && newInterval[0] >= leftVal[0] && newInterval[0] <= interval[0]) {
                if (leftVal[1] < newInterval[0]) {
                    res.add(leftVal);
                    leftVal = newInterval;
                } else {
                    leftVal[1] = Math.max(leftVal[1], newInterval[1]);
                }
            }

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

        if (newInterval[0] >= res.get(res.size() - 1)[0]) {
            if (res.get(res.size() - 1)[1] >= newInterval[0]) {
                res.get(res.size() - 1)[1] = Math.max(newInterval[1], res.get(res.size() - 1)[1]);
            } else {
                res.add(newInterval);
            }
        }

        return res.toArray(new int[][]{});
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.print(
            solution.insert(new int[][]{
                new int[]{1, 3},
                new int[]{6, 9}
            }, new int[]{2, 5})
        );  // [[1,5],[6,9]]

        System.out.print(
            solution.insert(new int[][]{
                new int[]{1, 2},
                new int[]{3, 5},
                new int[]{6, 7},
                new int[]{8, 10},
                new int[]{12, 16},
            }, new int[]{4, 8})
        );  // [[1,2],[3,10],[12,16]]

        System.out.print(
            solution.insert(new int[][]{
            }, new int[]{5, 7})
        );  // [[5,7]]

        System.out.print(
            solution.insert(new int[][]{
                new int[]{1, 5},
            }, new int[]{2, 3})
        );  // [[1,s]]

        System.out.print(
            solution.insert(new int[][]{
                new int[]{1, 5}
            }, new int[]{2, 7})
        );  // [[1,7]]
    }
}
// @lc code=end

