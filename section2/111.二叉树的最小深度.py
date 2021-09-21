#
# @lc app=leetcode.cn id=111 lang=python3
#
# [111] 二叉树的最小深度
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
    def minDepth(self, root: TreeNode) -> int:
        left_depth, right_depth = 0, 0
        if not root:
            return 0

        if root.left:
            left_depth = self.minDepth(root.left)

        if root.right:
            right_depth = self.minDepth(root.right)

        if left_depth and right_depth:
            return min(left_depth, right_depth) + 1

        return max(left_depth, right_depth) + 1

    def test(self):
        print(self.minDepth(TreeNode.generate([2, None, 3, None, 4, None, 5, None, 6])) == 5)
        print(self.minDepth(TreeNode.generate([3, 9, 20, None, None, 15, 7])) == 2)


Solution().test()
# @lc code=end
