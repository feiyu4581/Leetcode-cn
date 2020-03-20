import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=8 lang=java
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
class Solution {
    private final int maxRangeNum = 214748364; // 这个数字乘以10后的结果为 (2 << 31) / 10;
    private final int maxNum = 7;  // maxRangeNum * 10 + maxNum == x << 31 - 1
    private final int maxPositionNum = 8; // maxRangeNum * 10 + maxNum = x << 31;

    public int myAtoi(String str) {
        boolean header = true;
        ArrayList<Integer> nums = new ArrayList<>();
        int flag = 1;

        for (int i = 0; i < str.length(); i++) {
            char current = str.charAt(i);
            if (header & current == ' ') {
                continue;
            } else if (header & current == '+') {
                flag = 1;
                header = false;
            } else if (header & current == '-') {
                header = false;
                flag = -1;
            } else if (Character.isDigit(current)) {
                nums.add(Integer.parseInt(String.valueOf(current)));
                header = false;
            } else {
                break;
            }
        }

        int returnNum = 0;
        for (int num : nums) {
            if (returnNum != 0 | num != 0) {
                if (flag > 0 & ((maxRangeNum == returnNum & num > maxNum) | returnNum > maxRangeNum)) {
                    return maxRangeNum * 10 + maxNum;
                } else if (flag < 0 & ((maxRangeNum == returnNum & num > maxPositionNum) | returnNum > maxRangeNum)) {
                    return maxRangeNum * -10 - maxPositionNum;
                }

                returnNum = returnNum * 10 + num;
            }
        }

        return returnNum * flag;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.myAtoi("42") == 42
        );

        System.out.println(
            solution.myAtoi("    -42") == -42
        );

        System.out.println(
            solution.myAtoi("4193 with words") == 4193
        );

        System.out.println(
            solution.myAtoi("words and 987") == 0
        );

        System.out.println(
            solution.myAtoi("-91283472332") == -2147483648
        );

        System.out.println(
            solution.myAtoi("0-1") == 0
        );
    }
}
// @lc code=end

