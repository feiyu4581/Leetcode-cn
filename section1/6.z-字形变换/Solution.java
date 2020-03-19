import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=6 lang=java
 *
 * [6] Z 字形变换
 */

// @lc code=start
class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1) {
            return s;
        }
        ArrayList<ArrayList<Character>> maps = new ArrayList<ArrayList<Character>>();
        for (int i = 0; i < numRows; i++) {
            maps.add(new ArrayList<>());
        }

        int row = 0, flag = 1;
        for (int i = 0; i < s.length(); i++) {
            maps.get(row).add(s.charAt(i));

            if (row + flag >= numRows) {
                flag = -1;
            } else if (row + flag < 0) {
                flag = 1;
            }
            row += flag;
        }

        StringBuilder builder = new StringBuilder();
        for (ArrayList<Character> rowMaps : maps) {
            for (Character c : rowMaps) {
                builder.append(c);
            }
        }
        return builder.toString();
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.convert("LEETCODEISHIRING", 3).equals("LCIRETOESIIGEDHN")
        );

        System.out.println(
            solution.convert("LEETCODEISHIRING", 4).equals("LDREOEIIECIHNTSG")
        );

        System.out.println(
            solution.convert("AB", 1)
        );
    }
}
// @lc code=end

