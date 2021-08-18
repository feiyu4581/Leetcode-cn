
/*
 * @lc app=leetcode.cn id=91 lang=java
 *
 * [91] 解码方法
 */

// @lc code=start
class Solution {
    public int numDecodings(String s) {
        char[] nums = s.toCharArray();

        int previous = 1, previous2 = 0, res = 0;
        if (nums.length == 0 || nums[0] == '0') {
            return 0;
        }

        for (int i = 0; i < nums.length; i++) {
            res = 0;
            if (nums[i] != '0') {
                res += previous;
            }

            if (i > 0 && (nums[i  - 1] == '1' || (nums[i - 1] == '2' && nums[i] <= '6'))) {
                res += previous2;
            }

            previous2 = previous;
            previous = res;
        }


        return res;
    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        System.out.println(solution.numDecodings("12") == 2);
        System.out.println(solution.numDecodings("226") == 3);
        System.out.println(solution.numDecodings("0") == 0);
        System.out.println(solution.numDecodings("06") == 0);
    }
}
// @lc code=end
