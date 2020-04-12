import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=31 lang=java
 *
 * [31] 下一个排列
 */

// @lc code=start
class Solution {
    /*
     * 1, 2, 3, 4, 5
     * 1, 2, 3, 5, 4
     * 1, 2, 4, 3, 5
     * 1, 2, 4, 5, 3
     * 1，2, 5, 3, 4
     * 1, 2, 5, 4, 3
     * 1, 3, 2, 4, 5
     * 1, 3, 2, 5, 4
     * 1, 3, 4, 2, 5
     * 1, 3, 4, 5, 2
     * 1, 3, 5, 2, 4
     * 1, 3, 5, 4, 2
     * 1, 4, 2, 3, 5
     */

    private void swap(int[] nums, int left, int right) {
        int temp = nums[left];
        nums[left] = nums[right];
        nums[right] = temp;
    }
    private void reverse(int[] nums, int start, int end) {
        while (start < end) {
            swap(nums, start++, end--);
        }
    }
    public void nextPermutation(int[] nums) {
        int turning = -1;
        for (int i = nums.length - 1; i > 0; i--) {
            if (nums[i] > nums[i - 1]) {
                turning = i - 1;
                break;
            }
        }

        if (turning >= 0) {
            for (int i = nums.length - 1; i > turning; i--) {
                if (nums[i] > nums[turning]) {
                    swap(nums, i, turning);
                    break;
                }
            }
        }

        reverse(nums, turning + 1, nums.length - 1);
    }

    public static void main(String[] args) {
        Solution solution = new Solution();

        // int[] testNums = new int[]{1, 2, 3};
        // solution.nextPermutation(testNums);
        // System.out.println(Arrays.toString(testNums));

        // int[] testNums2 = new int[]{3, 2, 1};
        // solution.nextPermutation(testNums2);
        // System.out.println(Arrays.toString(testNums2));

        // int[] testNums3 = new int[]{1, 1, 5};
        // solution.nextPermutation(testNums3);
        // System.out.println(Arrays.toString(testNums3));

        int[] testNums4 = new int[]{1, 3, 2};
        solution.nextPermutation(testNums4);
        System.out.println(Arrays.toString(testNums4));
    }
}
// @lc code=end

