#
# @lc app=leetcode.cn id=110 lang=python3
#
# [110] 平衡二叉树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

    def __str__(self):
        return f'node(val={self.val})'

    @staticmethod
    def generate(nums: list):
        if not nums:
            return None
        root = TreeNode(nums[0])
        nodes = [root]

        index = 1
        while index < len(nums):
            new_nodes = []
            for node in nodes:
                if node:
                    node.left = None if index >= len(nums) or nums[index] is None else TreeNode(nums[index])
                    node.right = None if index + 1 >= len(nums) or nums[index + 1] is None else TreeNode(nums[index + 1])

                new_nodes.append(node.left if node else None)
                new_nodes.append(node.right if node else None)
                index += 2

            nodes = new_nodes

        return root

    def show(self):
        stacks = [self]
        nodes = []
        while stacks:
            row, new_stacks = [], []
            for node in stacks:
                nodes.append(node.val if node else None)
                if node:
                    new_stacks.append(node.left)
                    new_stacks.append(node.right)

            stacks = new_stacks

        return nodes


class Solution:
    def computeHeight(self, root):
        if not root:
            return True, 0

        if not root.left and not root.right:
            return True, 1

        left_ok, left_height = self.computeHeight(root.left)
        right_ok, right_height = self.computeHeight(root.right)

        return left_ok and right_ok and abs(left_height - right_height) <= 1, max(left_height, right_height) + 1

    def isBalanced(self, root: TreeNode) -> bool:
        is_ok, _ = self.computeHeight(root)

        return is_ok

    def test(self):
        print(self.isBalanced(TreeNode.generate([])))
        print(not self.isBalanced(TreeNode.generate([1, 2, 2, 3, 3, None, None, 4, 4])))
        print(self.isBalanced(TreeNode.generate([3, 9, 20, None, None, 15, 7])))


Solution().test()
# @lc code=end
