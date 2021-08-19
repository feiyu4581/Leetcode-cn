/*
 * @lc app=leetcode.cn id=98 lang=java
 *
 * [98] 验证二叉搜索树
 */

// @lc code=start
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
    private boolean checkValidBST(TreeNode root, Integer lowwer, Integer upper) {
        if (root == null) {
            return true;
        }

        if (lowwer != null && root.val <= lowwer) {
            return false;
        }

        if (upper != null && root.val >= upper) {
            return false;
        }

        if (root.left != null && root.left.val >= root.val) {
            return false;
        }

        if (root.right != null && root.right.val <= root.val) {
            return false;
        }

        return checkValidBST(root.left, lowwer, root.val) && checkValidBST(root.right, root.val, upper);
    }

    public boolean isValidBST(TreeNode root) {
        return checkValidBST(root, null, null);
    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        System.out.println(solution.isValidBST(new TreeNode(2147483647)));

        System.out.println(solution.isValidBST(
                new TreeNode(2, new TreeNode(1), new TreeNode(3))
        ));

        System.out.println(!solution.isValidBST(
                new TreeNode(5, new TreeNode(1), new TreeNode(4, new TreeNode(3), new TreeNode(6)))
        ));

        System.out.println(!solution.isValidBST(
                new TreeNode(5, new TreeNode(4), new TreeNode(6, new TreeNode(3), new TreeNode(7)))
        ));
    }
}
// @lc code=end

