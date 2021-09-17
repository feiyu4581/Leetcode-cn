import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=89 lang=java
 *
 * [89] 格雷编码
 */

// @lc code=start
class Solution {

    public List<Integer> grayCode(int n) {
        List<Integer> res = new ArrayList<>();
        if (n == 0) {
            res.add(n);
            return res;
        }

        int currentNum = 0;
        List<Integer> childs = grayCode(n - 1);
        for (int num : childs) {
            res.add(currentNum + num);
        }

        currentNum = (int) Math.pow(2, n - 1);
        for (int i = childs.size() - 1; i >= 0; i--) {
            res.add(currentNum + childs.get(i));
        }

        return res;
    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        // [0,1,3,2]
        System.out.println(solution.grayCode(2));
        // [0]
        System.out.println(solution.grayCode(0));
        System.out.println(solution.grayCode(3));
    }
}
