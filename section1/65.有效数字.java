
/*
 * @lc app=leetcode.cn id=65 lang=java
 *
 * [65] 有效数字
 */

// @lc code=start
class Solution {
    public boolean isNumber(String s) {
        return checkNumber(s.toCharArray(), 0, true, true) > 0;
    }

    private boolean isDigist(char num) {
        return num >= '0' && num <= '9';
    }

    public int checkNumber(char[] numbers, int index, boolean includeP, boolean includeE) {
        boolean start = false;
        int res = 0;
        for (; index < numbers.length; index++) {
            char num = numbers[index];
            if (isDigist(num)) {
                start = true;
                res ++;
            } else if (num == '+' || num == '-') {
                if (start) {
                    return 0;
                }

                start = true;
            } else if (num == '.') {
                if (!includeP) {
                    return 0;
                }

                includeP = false;
                if (index > 0 && isDigist(numbers[index - 1])) {
                    continue;
                }

                if (index < numbers.length - 1 && isDigist(numbers[index + 1])) {
                    continue;
                }

                return 0;
            } else if (num == 'e' || num == 'E') {
                if (!includeE) {
                    return 0;
                }

                return res > 0 && checkNumber(numbers, index + 1, false, false) >= 1? 1 : 0;
            } else {
                return 0;
            }
        }

        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.isNumber("0"));
        System.out.println(!solution.isNumber("e"));
        System.out.println(!solution.isNumber("."));
        System.out.println(solution.isNumber(".1"));
        System.out.println(!solution.isNumber("e9"));
        System.out.println(!solution.isNumber("G76"));
        System.out.println(solution.isNumber("-1."));
    }
}
// @lc code=end

