/*
 * @lc app=leetcode.cn id=58 lang=java
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
class Solution {
    public int lengthOfLastWord(String s) {
        boolean start = false;
        int res = 0;
        for (int i = s.length() - 1; i >= 0; i--) {
            char word = s.charAt(i);

            if (word == ' ' && start) {
                break;
            }

            if (word != ' ') {
                start = true;
                res ++;
            }
        }

        return res;
    }
    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.print(
            solution.lengthOfLastWord("Hello World") == 5
        );

        System.out.print(
            solution.lengthOfLastWord("   fly me   to   the moon  ") == 4
        );

        System.out.print(
            solution.lengthOfLastWord("luffy is still joyboy") == 6
        );
    }
}
// @lc code=end

