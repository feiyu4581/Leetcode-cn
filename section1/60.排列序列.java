/*
 * @lc app=leetcode.cn id=60 lang=java
 *
 * [60] 排列序列
 */

// @lc code=start
class Solution {
    public String getPermutation(int n, int k) {
        char[] start = new char[n];
        for (int i = 1; i <= n; i++) {
            start[i - 1] = (char) ('0' + i);
        }

        getPermutationRecursive(start, k);
        return new String(start);
    }

    private void swap(char[] arrangement, int left, int right) {
        char temp = arrangement[left];
        arrangement[left] = arrangement[right];
        arrangement[right] = temp;
    }

    private void reversed(char[] arrangement, int start) {
        int right = arrangement.length - 1;
        while (start < right) {
            swap(arrangement, start++, right--);
        }
    }

    public void getPermutationRecursive(char[] arrangement, int k) {
        if (k <= 1) {
            return;
        }

        int index = 1, nums = 1;
        while (nums * (index + 1) < k) {
            index += 1;
            nums = nums * index;
        }

        reversed(arrangement, arrangement.length - index);

        if (k > nums) {
            int i = arrangement.length - 1;
            while (arrangement[i] < arrangement[arrangement.length - index - 1]) {
                i--;
            }

            swap(arrangement, i, arrangement.length - index - 1);
            reversed(arrangement, arrangement.length - index);
            getPermutationRecursive(arrangement, k - nums);
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.getPermutation(3, 3).equals("213"));
        System.out.println(solution.getPermutation(4, 9).equals("2314"));
        System.out.println(solution.getPermutation(3, 1).equals("123"));
        System.out.println(solution.getPermutation(3, 5).equals("312"));
    }
}
// @lc code=end

