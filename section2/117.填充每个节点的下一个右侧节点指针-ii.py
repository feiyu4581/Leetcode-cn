#
# @lc app=leetcode.cn id=117 lang=python3
#
# [117] 填充每个节点的下一个右侧节点指针 II
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""


class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next

    @staticmethod
    def generate(nums: list):
        if not nums:
            return None
        root = Node(nums[0])
        nodes = [root]

        index = 1
        while index < len(nums):
            new_nodes = []
            for node in nodes:
                if node:
                    node.left = None if index >= len(nums) or nums[index] is None else Node(nums[index])
                    node.right = None if index + 1 >= len(nums) or nums[index + 1] is None else Node(
                        nums[index + 1])
                    index += 2

                new_nodes.append(node.left if node else None)
                new_nodes.append(node.right if node else None)

            nodes = new_nodes

        return root

    def __str__(self):
        return f'[Node]: val={self.val}'


class Solution:
    def connect_recursive(self, root, next_root):
        if root:
            root.next = next_root
            if next_root:
                next_right = next_root.left or next_root.right
                next_node = next_root.next
                while not next_right and next_node:
                    next_right = next_node.left or next_node.right
                    next_node = next_node.next

                self.connect_recursive(root.right, next_right)
                self.connect_recursive(root.left, root.right or next_right)
            else:
                self.connect_recursive(root.right, None)
                self.connect_recursive(root.left, root.right)

    def connect(self, root: 'Node') -> 'Node':
        self.connect_recursive(root, None)

        return root

    def test(self):
        # [1, 2, 3, 4, 5, None, 6, 7, None, None, None, None, 8]
        root = Node.generate([2, 1, 3, 0, 7, 9, 1, 2, None, 1, 0, None, None, 8, 8, None, None, None, None, 7])
        self.connect(root)

Solution().test()
# @lc code=end
