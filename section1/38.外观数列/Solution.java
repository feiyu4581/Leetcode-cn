/*
 * @lc app=leetcode.cn id=38 lang=java
 *
 * [38] 报数
 */

// @lc code=start
class Solution {
    public String countAndSay(int n) {
        String nums = "1";
        for (int i = 1; i < n; i++) {
            StringBuilder nextNums = new StringBuilder();
            char num = '0';
            int size = 0;

            for (int j = 0; j < nums.length(); j++) {
                if (size == 0) {
                    num = nums.charAt(j);
                    size = 1;
                } else if (nums.charAt(j) == num) {
                    size += 1;
                } else {
                    nextNums.append((char)(size + '0'));
                    nextNums.append(num);

                    num = nums.charAt(j);
                    size = 1;
                }
            }

            if (size > 0) {
                nextNums.append((char)(size + '0'));
                nextNums.append(num);
            }

            nums = nextNums.toString();
        }

        return nums;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.countAndSay(1).equals("1")
        );

        System.out.println(
            solution.countAndSay(4).equals("1211")
        );
    }
}
// @lc code=end

