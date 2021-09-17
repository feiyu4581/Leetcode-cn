import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=22 lang=java
 *
 * [22] 括号生成
 */

// @lc code=start
class Solution {
    private void generateSubParenthesis(int target, int left, StringBuilder prefix, List<String> res) {
        if (target == 0) {
            StringBuilder suffix = new StringBuilder();
            for (int i = 0; i < left; i++) {
                suffix.append(')');
            }
            res.add(prefix.toString() + suffix.toString());
        } else {
            if (left > 0) {
                prefix.append(')');
                generateSubParenthesis(target, left - 1, prefix, res);
                prefix.deleteCharAt(prefix.length() - 1);
            }
            prefix.append('(');
            generateSubParenthesis(target - 1, left + 1, prefix, res);
            prefix.deleteCharAt(prefix.length() - 1);
        }
    }

    public List<String> generateParenthesis(int n) {
        List<String> res = new ArrayList<>();

        generateSubParenthesis(n, 0, new StringBuilder(), res);
        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.generateParenthesis(3)
        );
    }
}
// @lc code=end

