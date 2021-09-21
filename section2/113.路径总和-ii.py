#
# @lc app=leetcode.cn id=113 lang=python3
#
# [113] 路径总和 II
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from typing import Optional, List


# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
#
#     def __str__(self):
#         return f'node(val={self.val})'
#
#     @staticmethod
#     def generate(nums: list):
#         if not nums:
#             return None
#         root = TreeNode(nums[0])
#         nodes = [root]
#
#         index = 1
#         while index < len(nums):
#             new_nodes = []
#             for node in nodes:
#                 if node:
#                     node.left = None if index >= len(nums) or nums[index] is None else TreeNode(nums[index])
#                     node.right = None if index + 1 >= len(nums) or nums[index + 1] is None else TreeNode(
#                         nums[index + 1])
#                     index += 2
#
#                 new_nodes.append(node.left if node else None)
#                 new_nodes.append(node.right if node else None)
#
#             nodes = new_nodes
#
#         return root
#
#     def show(self):
#         stacks = [self]
#         nodes = []
#         while stacks:
#             row, new_stacks = [], []
#             for node in stacks:
#                 nodes.append(node.val if node else None)
#                 if node:
#                     new_stacks.append(node.left)
#                     new_stacks.append(node.right)
#
#             stacks = new_stacks
#
#         return nodes
#

class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> List[List[int]]:
        records = []
        self.compute_path_sub(root, targetSum, records, [])
        return records

    def compute_path_sub(self, root, targetSum, records, current):
        if not root:
            return

        current.append(root.val)
        if root and not root.left and not root.right and root.val == targetSum:
            records.append(list(current))

        self.compute_path_sub(root.left, targetSum - root.val, records, current)
        self.compute_path_sub(root.right, targetSum - root.val, records, current)

        current.pop()

    def test(self):
        # []
        print(not self.pathSum(TreeNode.generate([1, 2]), 0))
        # []
        print(not self.pathSum(TreeNode.generate([1, 2, 3]), 5))
        # [[5, 4, 11, 2], [5, 8, 4, 5]]
        print(self.pathSum(TreeNode.generate([5, 4, 8, 11, None, 13, 4, 7, 2, None, None, 5, 1]), 22))


# Solution().test()
# @lc code=end

