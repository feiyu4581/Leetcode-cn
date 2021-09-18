#
# @lc app=leetcode.cn id=107 lang=python3
#
# [107] 二叉树的层序遍历 II
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from typing import List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

    def __str__(self):
        return f'node(val={self.val})'

    @staticmethod
    def generate(nums: list):
        root = TreeNode(nums[0])
        nodes = [root]

        index = 1
        while index < len(nums):
            new_nodes = []
            for node in nodes:
                if node:
                    node.left = None if nums[index] is None else TreeNode(nums[index])
                    node.right = None if nums[index + 1] is None else TreeNode(nums[index + 1])

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
    def level_traverse(self, nodes: List[TreeNode], rows: List[List[int]]):
        if not nodes:
            return

        row, new_nodes = [], []
        for node in nodes:
            if node:
                row.append(node.val)
                new_nodes.extend([node.left, node.right])

        self.level_traverse(new_nodes, rows)
        if row:
            rows.append(row)

    def levelOrderBottom(self, root: TreeNode) -> List[List[int]]:
        rows = []
        self.level_traverse([root], rows)

        return rows

    def test(self):
        # [[15, 7], [9, 20], [3]]
        print(self.levelOrderBottom(TreeNode.generate([3, 9, 20, None, None, 15, 7])))

Solution().test()
# @lc code=end

