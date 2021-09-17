import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/*
 * @lc app=leetcode.cn id=67 lang=java
 *
 * [67] 二进制求和
 */

// @lc code=start
class Solution {
    public String addBinary(String a, String b) {
        int carry = 0;
        int leftSize = a.length();
        int rightSize = b.length();
        int maxSize = Math.max(leftSize, rightSize);
        List<Integer> nums = new ArrayList<>(maxSize + 1);
        for (int i = 0; i < maxSize; i++) {
            int left = i < leftSize? a.charAt(leftSize - i - 1) - '0' : 0;
            int right = i < rightSize? b.charAt(rightSize - i - 1) - '0' : 0;

            int value = left + right + carry;
            if (value > 1) {
                value = value - 2;
                carry = 1;
            } else {
                carry = 0;
            }

            nums.add(value);
        }

        if (carry > 0) {
            nums.add(carry);
        }

        char[] res = new char[nums.size()];
        for (int i = 0; i < nums.size(); i++) {
            res[i] = (char)(nums.get(nums.size() - i - 1) + '0');
        }

        return String.valueOf(res);
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.addBinary("11", "1").equals("100"));
        System.out.println(solution.addBinary("1010", "1011").equals("10101"));
    }
}
// @lc code=end

