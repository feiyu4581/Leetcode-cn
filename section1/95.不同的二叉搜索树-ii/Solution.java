/*
 * @lc app=leetcode.cn id=95 lang=java
 *
 * [95] 不同的二叉搜索树 II
 */

// @lc code=start

import java.util.ArrayList;
import java.util.List;

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */

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

class Solution {
    private List<TreeNode> generateTreeRecursive(int left, int right) {
        List<TreeNode> trees = new ArrayList<>();
        for (int i = left; i <= right; i++) {
            for (TreeNode leftNode : generateTreeRecursive(left, i - 1)) {
                for (TreeNode rightNode : generateTreeRecursive(i + 1, right)) {
                    trees.add(new TreeNode(i, leftNode, rightNode));
                }
            }
        }

        if (trees.size() == 0) {
            trees.add(null);
        }

        return trees;
    }

    public List<TreeNode> generateTrees(int n) {
        return generateTreeRecursive(1, n);
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.generateTrees(3));
        System.out.println(solution.generateTrees(1));
    }
}
// @lc code=end

