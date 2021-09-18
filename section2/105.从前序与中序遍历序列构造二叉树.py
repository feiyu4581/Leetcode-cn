#
# @lc app=leetcode.cn id=105 lang=python3
#
# [105] 从前序与中序遍历序列构造二叉树
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
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        if len(preorder) == 0:
            return None

        node = TreeNode(preorder[0])
        index = 0
        while index < len(inorder):
            if inorder[index] == preorder[0]:
                break

            index += 1

        node.left = self.buildTree(preorder[1:index + 1], inorder[0:index])
        node.right = self.buildTree(preorder[index + 1:], inorder[index + 1:])
        return node

    def test(self):
        # [3,9,20,null,null,15,7]
        print(self.buildTree([3,9,20,15,7], [9,3,15,20,7]).show())

        # [-1]
        print(self.buildTree([-1], [-1]).show())


Solution().test()
# @lc code=end

