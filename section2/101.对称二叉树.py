#
# @lc app=leetcode.cn id=101 lang=python3
#
# [101] 对称二叉树
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
    def is_same_by_recursive(self, left: TreeNode, right: TreeNode) -> bool:
        if not left and not right:
            return True

        if left and right and left.val == right.val:
            # 对于每个对称节点判断他俩左右节点的值
            return self.is_same_by_recursive(left.left, right.right) and self.is_same_by_recursive(left.right, right.left)

        return False

    def is_same(self, left: TreeNode, right: TreeNode) -> bool:
        """
        使用栈的方式来模拟递归算法
        """
        stacks = [(left, right)]
        while stacks:
            left_node, right_node = stacks.pop()
            if left_node and right_node and left_node.val == right_node.val:
                stacks.append((left_node.left, right_node.right))
                stacks.append((left_node.right, right_node.left))
            elif left_node or right_node:
                return False

        return True

    def isSymmetric(self, root: TreeNode) -> bool:
        return self.is_same(root.left, root.right)

    def test(self):
        print(self.isSymmetric(TreeNode.generate([1, 2, 2, 3, 4, 4, 3])))
        print(not self.isSymmetric(TreeNode.generate([1, 2, 2, None, 3, None, 3])))


Solution().test()

# @lc code=end

