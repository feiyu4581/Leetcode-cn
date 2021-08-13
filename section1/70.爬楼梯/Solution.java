/*
 * @lc app=leetcode.cn id=70 lang=java
 *
 * [70] 爬楼梯
 */

// @lc code=start
class Solution {
    public int climbStairs(int n) {
        if (n <= 2) {
            return n;
        }

        int first = 1;
        int second = 2;

        int res = 0;
        for (int i = 3; i <= n; i++) {
            res = first + second;
            first = second;
            second = res;
        }

        return second;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.climbStairs(2) == 2);
        System.out.println(solution.climbStairs(3) == 3);
        System.out.println(solution.climbStairs(4) == 5);
        System.out.println(solution.climbStairs(5) == 8);
    }
}
// @lc code=end

