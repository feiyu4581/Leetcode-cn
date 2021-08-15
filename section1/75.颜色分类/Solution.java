import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=75 lang=java
 *
 * [75] 颜色分类
 */

// @lc code=start
class Solution {
    private void swap(int[] nums, int left, int right) {
        int temp = nums[left];
        nums[left] = nums[right];
        nums[right] = temp;
    }

    public void sortColors(int[] nums) {
        int zeroNums = 0, oneNums = 0;
        for (int i : nums) {
            if (i == 0) {
                zeroNums++;
            } else if (i == 1) {
                oneNums++;
            }
        }

        int zeroStart = 0, oneStart = zeroNums, twoStart = zeroNums + oneNums;
        int zeroEnd = 0, oneEnd = oneStart, twoEnd = twoStart;

        int index = 0;
        while (index < nums.length) {
            if (nums[index] == 0) {
                if (index >= zeroStart && index < zeroEnd) {
                    index++;
                } else {
                    swap(nums, index, zeroEnd++);
                }
            } else if (nums[index] == 1) {
                if (index >= oneStart && index < oneEnd) {
                    index++;
                } else {
                    swap(nums, index, oneEnd++);
                }
            } else {
                if (index >= twoStart && index < twoEnd) {
                    index++;
                } else {
                    swap(nums, index, twoEnd++);
                }
            }
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] nums = new int[]{2,0,2,1,1,0};
        solution.sortColors(nums);
        System.out.println(Arrays.toString(nums)); // [0,0,1,1,2,2]

        int[] nums2 = new int[]{2,0,1};
        solution.sortColors(nums2);
        System.out.println(Arrays.toString(nums2)); // [0, 1, 2]

        int[] nums3 = new int[]{0};
        solution.sortColors(nums3);
        System.out.println(Arrays.toString(nums3)); // [0]

        int[] nums4 = new int[]{1};
        solution.sortColors(nums4);
        System.out.println(Arrays.toString(nums4)); // [1]
    }
}
// @lc code=end

