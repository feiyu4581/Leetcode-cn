#
# @lc app=leetcode.cn id=104 lang=python3
#
# [104] 二叉树的最大深度
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
    def maxDepth(self, root: TreeNode) -> int:
        if not root:
            return 0

        return max(self.maxDepth(root.left), self.maxDepth(root.right)) + 1

    def test(self):
        print(self.maxDepth(TreeNode.generate([3, 9, 20, None, None, 15, 7])) == 3)

# Solution().test()

# @lc code=end

