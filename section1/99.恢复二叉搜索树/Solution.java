/*
 * @lc app=leetcode.cn id=99 lang=java
 *
 * [99] 恢复二叉搜索树
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
        return "" + val;
    }
}

class Solution {
    private TreeNode InorderRecursive(List<TreeNode> errorNodes, TreeNode root, TreeNode lastNode) {
        if (root != null && errorNodes.size() < 3) {
            lastNode = InorderRecursive(errorNodes, root.left, lastNode);
            if (lastNode != null && lastNode.val > root.val) {
                if (errorNodes.isEmpty()) {
                    errorNodes.add(lastNode);
                    errorNodes.add(root);
                } else {
                    errorNodes.add(root);
                }
            }
            lastNode = InorderRecursive(errorNodes, root.right, root);
        }

        return lastNode;
    }

    public void recoverTree(TreeNode root) {
        List<TreeNode> errorNodes = new ArrayList<>();
        InorderRecursive(errorNodes, root, null);

        if (errorNodes.size() >= 2) {
            swapTreeNode(errorNodes.get(0), errorNodes.get(errorNodes.size() - 1));
        }
    }

    private void swapTreeNode(TreeNode treeNode, TreeNode treeNode1) {
        int temp = treeNode.val;
        treeNode.val = treeNode1.val;
        treeNode1.val = temp;
    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        solution.recoverTree(new TreeNode(1, new TreeNode(3, null, new TreeNode(2)), null));
        solution.recoverTree(new TreeNode(3, new TreeNode(1), new TreeNode(4, new TreeNode(2), null)));
    }
}
// @lc code=end

