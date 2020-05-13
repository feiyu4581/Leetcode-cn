import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

/*
 * @lc app=leetcode.cn id=49 lang=java
 *
 * [49] 字母异位词分组
 */

// @lc code=start
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        HashMap<String, List<String>> strMaps = new HashMap<>();
        for (String str : strs) {
            char[] charStr = str.toCharArray();
            Arrays.sort(charStr);
            String sortStr = new String(charStr);

            if (!strMaps.containsKey(sortStr)) {
                strMaps.put(sortStr, new ArrayList<>());
            }

            strMaps.get(sortStr).add(str);
        }

        return new ArrayList<>(strMaps.values());
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.groupAnagrams(new String[]{"eat", "tea", "tan", "ate", "nat", "bat"})
        );
    }

}
// @lc code=end

