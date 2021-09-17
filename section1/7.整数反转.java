import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=7 lang=java
 *
 * [7] 整数反转
 */

// @lc code=start
class Solution {
    private final int maxRangeNum = 214748364; // 这个数字乘以10后的结果为 (2 << 31) / 10;
    private final int maxNum = 7;  // maxRangeNum * 10 + maxNum == x << 31 - 1
    private final int maxPositionNum = 8; // maxRangeNum * 10 + maxNum = x << 31;
    public int reverse(int x) {
        int flag = x > 0? 1 : -1;
        x = Math.abs(x);
        ArrayList<Integer> nums = new ArrayList<>();

        while (x > 0) {
            nums.add(x % 10);
            x = x / 10;
        }

        int returnNum = 0;
        for (int num : nums) {
            if (returnNum != 0 || num != 0) {
                if (returnNum == maxRangeNum) {
                    if ((flag > 0 && num > maxNum) || (flag < 0 && num > maxPositionNum)) {
                        return 0;
                    }
                }
                
                if (returnNum > maxRangeNum) {
                    return 0;
                }

                returnNum = returnNum * 10 + num;
            }
        }

        return returnNum * flag;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.reverse(1534236469) == 0
        );
        System.out.println(
            solution.reverse(123) == 321
        );

        System.out.println(
            solution.reverse(-123) == -321
        );

        System.out.println(
            solution.reverse(120) == 21
        );
    }
}
// @lc code=end

