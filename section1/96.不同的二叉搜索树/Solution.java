/*
 * @lc app=leetcode.cn id=96 lang=java
 *
 * [96] 不同的二叉搜索树
 */

import java.util.HashMap;
import java.util.Map;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode() {}
    TreeNode(int val) { this.val = val; }
    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }

    @Override
    public String toString() {
        return "[" + val + "," + left + "," + right + "]";
    }
}

// @lc code=start
class Solution {
    Map<String, Integer> caches;
    private int numTreesRecursive(int left, int right) {
        String cacheStr = left + "/" + right;
        if (!caches.containsKey(cacheStr)) {
            int nums = 0;
            for (int i = left; i <= right; i++) {
                nums += numTreesRecursive(left ,i - 1) * numTreesRecursive(i + 1, right);
            }
            caches.put(cacheStr, nums == 0? 1 : nums);
        }

        return caches.get(cacheStr);
    }

    public int numTrees(int n) {
        this.caches = new HashMap<>();
        return numTreesRecursive(1, n);
    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        System.out.println(solution.numTrees(3) == 5);
        System.out.println(solution.numTrees(1) == 1);
        System.out.println(solution.numTrees(18) == 477638700);
    }
}
// @lc code=end

