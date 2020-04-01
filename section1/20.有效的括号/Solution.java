import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

/*
 * @lc app=leetcode.cn id=20 lang=java
 *
 * [20] 有效的括号
 */

// @lc code=start
class Solution {
    private HashMap<Character, Character> symbolMaps = new HashMap<>();

    public Solution() {
        symbolMaps.put(')', '(');
        symbolMaps.put(']', '[');
        symbolMaps.put('}', '{');
    }

    public boolean isValid(String s) {
        List<Character> stack = new ArrayList<>();
        for (int i = 0; i < s.length(); i++) {
            char current = s.charAt(i);
            if (current == ')' || current == ']' || current == '}') {
                if (stack.size() > 0 && stack.get(stack.size() - 1) == symbolMaps.get(current)) {
                    stack.remove(stack.size() - 1);
                } else {
                    return false;
                }
            } else {
                stack.add(current);
            }
        }

        return stack.size() == 0;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.isValid("()") == true
        );

        System.out.println(
            solution.isValid("()[]{}") == true
        );

        System.out.println(
            solution.isValid("(]") == false
        );

        System.out.println(
            solution.isValid("([)]") == false
        );

        System.out.println(
            solution.isValid("{[]}") == true
        );
    }
}
// @lc code=end

