import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=32 lang=java
 *
 * [32] 最长有效括号
 */

// @lc code=start
class Solution {
    public int longestValidParentheses(String s) {
        List<Integer> stack = new ArrayList<>();
        int[] leftIndexs = new int[s.length()];
        int[] RightIndexs = new int[s.length()];
        for (int i = 0; i < s.length(); i++) {
            leftIndexs[i] = 0;
            RightIndexs[i] = -1;
        }

        int maxLength = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') {
                stack.add(i);
            } else if (stack.size() > 0) {
                int leftIndex = stack.get(stack.size() - 1);
                stack.remove(stack.size() - 1);

                int size = i - leftIndex + 1;
                leftIndexs[leftIndex] = size;
                RightIndexs[i] = leftIndex;

                int maxSize = size;
                while (leftIndex >= 1) {
                    if (RightIndexs[leftIndex - 1] >= 0) {
                        leftIndex = RightIndexs[leftIndex - 1];
                        leftIndexs[leftIndex] += size;
                        maxSize = leftIndexs[leftIndex];
                    } else {
                        break;
                    }
                }

                if (maxSize > maxLength) {
                    maxLength = maxSize;
                }
            }
        }
        return maxLength;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.longestValidParentheses("(()") == 2
        );

        System.out.println(
            solution.longestValidParentheses(")()())") == 4
        );

        System.out.println(
            solution.longestValidParentheses("(()()(())((") == 8
        );
    }
}
// @lc code=end

