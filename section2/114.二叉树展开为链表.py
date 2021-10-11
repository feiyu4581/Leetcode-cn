#
# @lc app=leetcode.cn id=114 lang=python3
#
# [114] 二叉树展开为链表
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
    def fatten_recursive(self, root):
        current = root
        if root:
            right_node = root.right
            if root.left:
                left_root, left_tail = self.fatten_recursive(root.left)
                current.left, current.right = None, left_root
                current = left_tail

            if right_node:
                right_root, right_tail = self.fatten_recursive(right_node)
                current.left, current.right = None, right_root
                current = right_tail

        return root, current

    def flatten(self, root: TreeNode) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        self.fatten_recursive(root)

    def test(self):
        # [1,null,2,null,3,null,4,null,5,null,6]
        root = TreeNode.generate([1, 2, 5, 3, 4, None, 6])
        print(self.flatten(root))
        print(root.show())
        print(self.flatten(TreeNode.generate([])))
        print(self.flatten(TreeNode.generate([0])))


Solution().test()

# @lc code=end
