/*
 * @lc app=leetcode.cn id=4 lang=java
 *
 * [4] 寻找两个有序数组的中位数
 */

// @lc code=start
class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        /*
        nums1[0..m]
        nums2[0..n]

        i = 0~m
        j = 0~n

        i + j = (m + n) / 2

        if m + n % 2 == 0:
            results = (max(left) + min(right)) / 2
        else:
            results = min(right)

        i = 0 ~ m
        j = (m + n) / 2 - j

        A[i - 1] <= B[j]
        A[i] >= B[j - 1]

        if m > n => j may be position

        */

        if (nums1.length > nums2.length) {
            int[] temp = nums1;
            nums1 = nums2;
            nums2 = temp;
        }

        int left = 0, right = nums1.length, middle = (nums1.length + nums2.length) / 2;
        while (left <= right) {
            int i = (left + right) / 2;
            int j = middle - i;

            if (i > 0 && nums1[i - 1] > nums2[j]) {
                right = i - 1;
            } else if (i < nums1.length && nums1[i] < nums2[j - 1]) {
                left = i + 1;
            } else {
                int rightMin = 0;
                if (i == nums1.length) {
                    rightMin = nums2[j];
                } else if (j == nums2.length) {
                    rightMin = nums1[i];
                } else {
                    rightMin = Math.min(nums1[i], nums2[j]);
                }

                if ((nums1.length + nums2.length) % 2 != 0) {
                    return rightMin;
                }

                int leftMax = 0;
                if (i == 0) {
                    leftMax = nums2[j - 1];
                } else if (j == 0) {
                    leftMax = nums1[i - 1];
                } else {
                    leftMax = Math.max(nums1[i - 1], nums2[j - 1]);
                }

                return (leftMax + rightMin) / 2.0;
            }
        }
        return 0;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();

        System.out.println(solution.findMedianSortedArrays(
            new int[]{1, 3},
            new int[]{2}
        ));

        System.out.println(solution.findMedianSortedArrays(
            new int[]{1, 2},
            new int[]{3, 4}
        ));
    }
}
// @lc code=end

