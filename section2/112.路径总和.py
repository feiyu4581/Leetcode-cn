#
# @lc app=leetcode.cn id=112 lang=python3
#
# [112] 路径总和
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from typing import Optional


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
                    node.right = None if index + 1 >= len(nums) or nums[index + 1] is None else TreeNode(
                        nums[index + 1])
                    index += 2

                new_nodes.append(node.left if node else None)
                new_nodes.append(node.right if node else None)

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
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        if not root:
            return False

        if root and not root.left and not root.right and root.val == targetSum:
            return True

        return self.hasPathSum(root.left, targetSum - root.val) or self.hasPathSum(root.right, targetSum - root.val)

    def test(self):
        print(not self.hasPathSum(TreeNode.generate([1, 2]), 0))
        print(not self.hasPathSum(TreeNode.generate([1, 2, 3]), 5))
        print(self.hasPathSum(TreeNode.generate([5, 4, 8, 11, None, 13, 4, 7, 2, None, None, None, 1]), 22))


# Solution().test()
# @lc code=end
