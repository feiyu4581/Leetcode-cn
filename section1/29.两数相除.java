/*
 * @lc app=leetcode.cn id=29 lang=java
 *
 * [29] 两数相除
 */

// @lc code=start
class Solution {
    private final int maxNum = Integer.MAX_VALUE >> 2;
    public int divide(int dividend, int divisor) {
        if (divisor == Integer.MIN_VALUE) {
            return dividend == Integer.MIN_VALUE? 1 : 0;
        }

        if (dividend == Integer.MIN_VALUE && divisor == -1) {
            return Integer.MAX_VALUE;
        }

        int resNum = 0;
        int flag = dividend > 0? 1 : -1;
        flag = divisor > 0? flag : -flag;

        dividend = dividend > 0 ? -dividend : dividend;
        divisor = Math.abs(divisor);
        while (dividend <= -divisor) {
            int position = 0;
            while (true) {
                if (divisor << position > maxNum || -(divisor << position) < dividend) {
                    break;
                }

                position++;
            }

            if (position == 0) {
                resNum += 1;
                dividend += divisor;
            } else {
                resNum += 1 << (position - 1);
                dividend += divisor << (position - 1);
            }
        }

        return flag > 0? resNum : -resNum;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.divide(10, 3) == 3
        );

        System.out.println(
            solution.divide(7, -3) == -2
        );

        System.out.println(
            solution.divide(Integer.MIN_VALUE, -1) == Integer.MAX_VALUE
        );

        System.out.println(
            solution.divide(Integer.MIN_VALUE, 1) == Integer.MIN_VALUE
        );

        System.out.println(
            solution.divide(1026117192, -874002063) == -1
        );

        System.out.println(
            solution.divide(-1046638222, 564425312)
        );
    }
}
// @lc code=end



