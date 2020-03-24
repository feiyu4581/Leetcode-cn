import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=9 lang=java
 *
 * [9] 回文数
 */

// @lc code=start
class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0) {
            return false;
        }

        ArrayList<Integer> nums = new ArrayList<>();
        while (x > 0) {
            nums.add(x % 10);
            x = x / 10;
        }

        int left = 0, right = nums.size() - 1;
        while (left <= right) {
            if (nums.get(left) != nums.get(right)) {
                return false;
            }

            left += 1;
            right -= 1;
        }

        return true;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.isPalindrome(121) == true
        );

        System.out.println(
            solution.isPalindrome(-121) == false
        );

        System.out.println(
            solution.isPalindrome(10) == false
        );
    }
}
// @lc code=end

