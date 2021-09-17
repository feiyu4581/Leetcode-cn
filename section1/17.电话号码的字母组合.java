import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

/*
 * @lc app=leetcode.cn id=17 lang=java
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
class Solution {
    private HashMap<Character, String> letterMaps = new HashMap<>();

    public Solution() {
        letterMaps.put('2', "abc");
        letterMaps.put('3', "def");
        letterMaps.put('4', "ghi");
        letterMaps.put('5', "jkl");
        letterMaps.put('6', "mno");
        letterMaps.put('7', "pqrs");
        letterMaps.put('8', "tuv");
        letterMaps.put('9', "wxyz");
    }

    private void _letterCombinations(List<String> res, StringBuilder prefix, String digits, int index) {
        if (index >= digits.length()) {
            res.add(prefix.toString());
        } else {
            String current = letterMaps.get(digits.charAt(index));
            for (int i = 0; i < current.length(); i++) {
                prefix.append(current.charAt(i));
                _letterCombinations(res, prefix, digits, index + 1);
                prefix.deleteCharAt(prefix.length() - 1);
            }
        }

    }

    public List<String> letterCombinations(String digits) {
        List<String> res = new ArrayList<>();
        if (!digits.isEmpty()) {
            _letterCombinations(res, new StringBuilder(), digits, 0);
        }
        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.letterCombinations("23")
        );
    }
}
// @lc code=end

