/*
 * @lc app=leetcode.cn id=94 lang=java
 *
 * [94] 二叉树的中序遍历
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
}


class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        inorderTraversalRecursive(root, res);

        return res;
    }

    private void inorderTraversalRecursive(TreeNode root, List<Integer> res) {
        if (root != null) {
            inorderTraversalRecursive(root.left, res);
            res.add(root.val);
            inorderTraversalRecursive(root.right, res);
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.inorderTraversal(new TreeNode(1, null, new TreeNode(2, new TreeNode(3), null))));
        System.out.println(solution.inorderTraversal(null));
        System.out.println(solution.inorderTraversal(new TreeNode(1)));
        System.out.println(solution.inorderTraversal(new TreeNode(1, new TreeNode(2), null)));
        System.out.println(solution.inorderTraversal(new TreeNode(1, null, new TreeNode(2))));
    }
}
// @lc code=end

