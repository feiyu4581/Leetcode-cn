import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=43 lang=java
 *
 * [43] 字符串相乘
 */

// @lc code=start
class Solution {
    public String multiply(String num1, String num2) {
        List<String> nums = new ArrayList<>();
        for (int i = num2.length() - 1; i >= 0; i--) {
            StringBuilder num = new StringBuilder();
            for (int offset = num2.length() - 1; offset > i; offset--) {
                num.append('0');
            }
            int current = (int)num2.charAt(i) - '0', carry = 0;
            for (int j = num1.length() - 1; j >= 0; j--) {
                int multiplicand = (int)num1.charAt(j) - '0';
                int result = current * multiplicand + carry;
                num.append(result % 10);
                carry = result / 10;
            }

            if (carry > 0) {
                num.append((char)(carry + '0'));
            }

            nums.add(num.toString());
        }

        StringBuilder resNums = new StringBuilder();
        boolean isEnd = false;
        int carry = 0, index = 0;
        while (!isEnd) {
            int num = carry;
            for (int i = 0; i < nums.size(); i++) {
                if (index >= nums.get(i).length()) {
                    isEnd = true;
                } else {
                    num += (int)(nums.get(i).charAt(index)) - '0';
                    isEnd = false;
                }
            }

            if (!isEnd) {
                resNums.append(num % 10);
                carry = num / 10;
                index += 1;
            }
        }
        if (carry > 0) {
            resNums.append(carry);
        }

        resNums.reverse();
        String res = resNums.toString();
        int findZero = -1;
        for (int i = 0; i < res.length(); i++) {
            if (res.charAt(i) != '0') {
                break;
            }

            findZero = i;
        }

        if (findZero >= 0) {
            res = res.substring(findZero + 1, res.length());
        }

        if (res.length() == 0) {
            return "0";
        }

        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.multiply("2", "3").equals("6")
        );

        System.out.println(
            solution.multiply("123", "456").equals("56088")
        );

        System.out.println(
            solution.multiply("140", "721").equals("100940")
        );
    }
}
// @lc code=end

