import java.util.HashMap;

/*
 * @lc app=leetcode.cn id=3 lang=java
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
class Solution {
    public int lengthOfLongestSubstring(String s) {
        if (s.isEmpty()) {
            return 0;
        }

        int res = 0;
        int current = 0;

        HashMap<Character, Integer> charIndexMaps = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            if (charIndexMaps.getOrDefault(s.charAt(i), -1) >= current) {
                current = charIndexMaps.get(s.charAt(i)) + 1;
            }

            res = Math.max(res, i - current + 1);
            charIndexMaps.put(s.charAt(i), i);
        }

        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.lengthOfLongestSubstring("abcabcbb") == 3);
        System.out.println(solution.lengthOfLongestSubstring("bbbb") == 1);
        System.out.println(solution.lengthOfLongestSubstring("pwwkew") == 3);
    }
}
// @lc code=end

