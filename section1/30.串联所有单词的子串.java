import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

/*
 * @lc app=leetcode.cn id=30 lang=java
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
class Solution {
    public List<Integer> findSubstring(String s, String[] words) {
        List<Integer> resNums = new ArrayList<>();
        if (words.length == 0) {
            return resNums;
        }

        int wordLength = words[0].length();

        HashMap<String, Integer> strMaps = new HashMap<>();
        for (String word : words) {
            strMaps.putIfAbsent(word, 0);
            strMaps.put(word, strMaps.get(word) + 1);
        }

        for (int i = 0; i <= s.length() - wordLength * words.length; i++) {
            HashMap<String, Integer> currentMaps = new HashMap<>();
            int step = 0;
            for (step = 0; step < words.length; step++) {
                String current = s.substring(i + step * wordLength, i + step * wordLength + wordLength);
                currentMaps.putIfAbsent(current, 0);

                if (!strMaps.containsKey(current) || currentMaps.get(current) >= strMaps.get(current)) {
                    break;
                }

                currentMaps.put(current, currentMaps.get(current) + 1);
            }

            if (step == words.length) {
                resNums.add(i);
            }
        }

        return resNums;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.findSubstring("barfoothefoobarman", new String[]{"foo", "bar"})
        );

        System.out.println(
            solution.findSubstring("wordgoodgoodgoodbestword", new String[]{"word","good","best","word"})
        );
    }
}
// @lc code=end

