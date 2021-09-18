#
# @lc app=leetcode.cn id=103 lang=python3
#
# [103] 二叉树的锯齿形层序遍历
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


class Solution:
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        stacks = [root]
        rows = []
        layer = 1
        while stacks:
            row, new_stacks = [], []
            for node in stacks:
                if node:
                    row.append(node.val)
                    new_stacks.append(node.left)
                    new_stacks.append(node.right)

            stacks = new_stacks
            if layer % 2 == 0:
                row = list(reversed(row))

            if row:
                rows.append(row)

            layer += 1

        return rows

    def test(self):
        # [
        #     [3],
        #     [20,9],
        #     [15,7]
        # ]
        print(self.zigzagLevelOrder(TreeNode.generate([3, 9, 20, None, None, 15, 7])))


# Solution().test()
# @lc code=end

