import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=66 lang=java
 *
 * [66] 加一
 */

// @lc code=start
class Solution {
    public int[] plusOne(int[] digits) {
        int carry = 1;
        for (int i = digits.length - 1; i >= 0; i--) {
            if (digits[i] + carry > 9) {
                digits[i] = 0;
                carry = 1;
            } else {
                digits[i] += carry;
                carry = 0;
            }

            if (carry == 0) {
                break;
            }
        }

        if (carry > 0) {
            int[] res = new int[digits.length + 1];
            res[0] = carry;
            for (int index = 0; index < digits.length; index++) {
                res[index + 1] = digits[index];
            }

            return res;
        }

        return digits;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();

        System.out.println(Arrays.equals(
            solution.plusOne(new int[]{1, 2, 3}),
            new int[]{1, 2, 4}
        ));

        System.out.println(Arrays.equals(
            solution.plusOne(new int[]{4, 3, 2, 1}),
            new int[]{4, 3, 2, 2}
        ));

        System.out.println(Arrays.equals(
            solution.plusOne(new int[]{0}),
            new int[]{1}
        ));

        System.out.println(Arrays.equals(
            solution.plusOne(new int[]{9, 9, 9, 9}),
            new int[]{1, 0, 0, 0, 0}
        ));
    }
}
// @lc code=end

