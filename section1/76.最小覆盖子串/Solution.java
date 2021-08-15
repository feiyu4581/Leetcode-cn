import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/*
 * @lc app=leetcode.cn id=76 lang=java
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
class Solution {
    public String minWindow(String s, String t) {
        Map<Character, Integer> numMaps = new HashMap<>();
        Map<Character, Integer> currentMaps = new HashMap<>();

        for (int i = 0; i < t.length(); i++) {
            numMaps.putIfAbsent(t.charAt(i), 0);
            numMaps.put(t.charAt(i), numMaps.get(t.charAt(i)) + 1);

            currentMaps.putIfAbsent(t.charAt(i), 0);
        }

        Map<Character, List<Integer>> numPositions = new HashMap<>();
        int[] currentNums = new int[s.length()];
        for (int i = 0; i < currentNums.length; i++) {
            currentNums[i] = 0;
        }

        int currentPosition = 0, currentStart = s.length() + 1;
        int minLength = s.length() + 1, minStart = s.length() + 1, minEnd = 0;
        for (int i = 0; i < s.length(); i++) {
            char num = s.charAt(i);
            if (numMaps.containsKey(num)) {
                numPositions.putIfAbsent(num, new ArrayList<>());

                List<Integer> positions = numPositions.get(num);
                positions.add(i);

                if (currentMaps.get(num) < numMaps.get(num)) {
                    currentMaps.put(num, currentMaps.get(num) + 1);
                    currentPosition ++;
                }

                currentNums[i] = 1;
                if (positions.size() > numMaps.get(num)) {
                    int replacePosition = positions.get(positions.size() - numMaps.get(num) - 1);
                    currentNums[replacePosition] = 0;

                    if (replacePosition == currentStart) {
                        while (currentNums[currentStart] == 0) {
                            currentStart++;
                        }
                    }
                } else {
                    currentStart = Math.min(currentStart, i);
                }

                if (currentPosition >= t.length() && i - currentStart < minLength) {
                    minLength = i - currentStart;
                    minStart = currentStart;
                    minEnd = i;
                }
            }
        }

        if (minLength <= s.length()) {
            return s.substring(minStart, minEnd + 1);
        }
        return "";
    }

    public static void main(String[] args) {
        Solution solution = new Solution();

        System.out.println(solution.minWindow("ADOBECODEBANC", "ABC").equals("BANC"));
        System.out.println(solution.minWindow("a", "a").equals("a"));
        System.out.println(solution.minWindow("a", "aa").equals(""));
        System.out.println(solution.minWindow("ab", "b").equals("b"));
        System.out.println(solution.minWindow("bba", "ab").equals("ba"));
    }
}
// @lc code=end

