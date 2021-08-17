/*
 * @lc app=leetcode.cn id=85 lang=java
 *
 * [85] 最大矩形
 */

// @lc code=start
class Solution {
    public int maximalRectangle(char[][] matrix) {

    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        System.out.println(solution.maximalRectangle(new char[][]{
            new char[]{'1','0','1','0','0'},
            new char[]{'1','0','1','1','1'},
            new char[]{'1','1','1','1','1'},
            new char[]{'1','0','0','1','0'},
        }) == 6);

        System.out.println(solution.maximalRectangle(new char[][]{new char[]{'0'}}) == 0);
        System.out.println(solution.maximalRectangle(new char[][]{new char[]{'1'}}) == 1);
        System.out.println(solution.maximalRectangle(new char[][]{new char[]{'0', '0'}}) == 0);
    }
}
// @lc code=end

