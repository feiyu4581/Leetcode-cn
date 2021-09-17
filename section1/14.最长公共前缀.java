/*
 * @lc app=leetcode.cn id=14 lang=java
 *
 * [14] 最长公共前缀
 */

// @lc code=start
class Solution {
    public String longestCommonPrefix(String[] strs) {
        if (strs.length == 0) {
            return "";
        };

        StringBuilder builder = new StringBuilder();
        int index = 0;
        while (true) {
            boolean need_break = false;
            char prefix = ' ';
            for (int i = 0; i < strs.length; i++) {
                if (index >= strs[i].length()) {
                    need_break = true;
                    break;
                }

                if (i == 0) {
                    prefix = strs[i].charAt(index);
                } else if (prefix != strs[i].charAt(index)) {
                    need_break = true;
                    break;
                }
            }

            if (need_break) {
                break;
            }

            builder.append(prefix);
            index += 1;
        }

        return builder.toString();
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.longestCommonPrefix(new String[]{"flower","flow","flight"}).equals("fl")
        );

        System.out.println(
            solution.longestCommonPrefix(new String[]{"dog","racecar","car"}).equals("")
        );
    }
}
// @lc code=end

