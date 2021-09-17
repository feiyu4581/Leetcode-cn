import java.util.ArrayList;
import java.util.HashMap;

/*
 * @lc app=leetcode.cn id=13 lang=java
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
class Solution {
    private HashMap<String, Integer> romanMaps = new HashMap<>();

    public Solution() {
        romanMaps.put("M", 1000);
        romanMaps.put("CM", 900);
        romanMaps.put("D", 500);
        romanMaps.put("CD", 400);
        romanMaps.put("C", 100);
        romanMaps.put("XC", 90);
        romanMaps.put("L", 50);
        romanMaps.put("XL", 40);
        romanMaps.put("X", 10);
        romanMaps.put("IX", 9);
        romanMaps.put("V", 5);
        romanMaps.put("IV", 4);
        romanMaps.put("I", 1);
    }

    public int romanToInt(String s) {
        int resNum = 0, index = 0;
        while (index < s.length()) {
            if (index + 1 < s.length() && romanMaps.containsKey(s.substring(index, index + 2))) {
                resNum += romanMaps.get(s.substring(index, index + 2));
                index += 2;
            } else {
                resNum += romanMaps.get(s.substring(index, index + 1));
                index += 1;
            }
        }

        return resNum;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.romanToInt("III") == 3
        );

        System.out.println(
            solution.romanToInt("IV") == 4
        );

        System.out.println(
            solution.romanToInt("IX") == 9
        );

        System.out.println(
            solution.romanToInt("LVIII") == 58
        );

        System.out.println(
            solution.romanToInt("MCMXCIV") == 1994
        );
    }
}
// @lc code=end

