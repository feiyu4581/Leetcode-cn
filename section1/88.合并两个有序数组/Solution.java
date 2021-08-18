/*
 * @lc app=leetcode.cn id=88 lang=java
 *
 * [88] 合并两个有序数组
 */

import java.util.Arrays;

// @lc code=start
class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int current = m + n - 1;
        while (m > 0 || n > 0) {
            if (m > 0 && n > 0) {
                if (nums1[m - 1] >= nums2[n - 1]) {
                    nums1[current] = nums1[--m];
                } else {
                    nums1[current] = nums2[--n];
                }
            } else if (n > 0) {
                nums1[current] = nums2[--n];
            } else {
                break;
            }

            current--;
        }

    }
    public static void main(String[] ags) {
        Solution solution = new Solution();
        solution.merge(new int[]{1,2,3,0,0,0}, 3, new int[]{2, 5, 6}, 3);
        solution.merge(new int[]{1}, 1, new int[]{}, 0);
        solution.merge(new int[]{0}, 0, new int[]{1}, 1);
    }
}
// @lc code=end

