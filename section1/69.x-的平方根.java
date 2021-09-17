/*
 * @lc app=leetcode.cn id=69 lang=java
 *
 * [69] x 的平方根
 */

// @lc code=start
class Solution {
    private static int minValue = 46340;

    public int mySqrt(int x) {
        int left = 0;
        int right = x;

        while (left <= right) {
            int mid = left + ((right - left) >> 2);
            if (mid > minValue) {
                right = mid - 1;
                continue;
            }

            int value = mid * mid;
            if (value == x) {
                return mid;
            } else if (value > x) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        return right;
    }
    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.mySqrt(4) == 2);
        System.out.println(solution.mySqrt(8) == 2);
        System.out.println(solution.mySqrt(9) == 3);
        System.out.println(solution.mySqrt(80) == 8);
        System.out.println(solution.mySqrt(2147395599) == 46339);
    }
}
// @lc code=end

