#
# @lc app=leetcode.cn id=116 lang=python3
#
# [116] 填充每个节点的下一个右侧节点指针
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


class Solution:
    def connect_recursive(self, root, next_root):
        if root:
            root.next = next_root
            self.connect_recursive(root.left, root.right)
            self.connect_recursive(root.right, next_root.left if next_root else None)

    def connect(self, root: 'Node') -> 'Node':
        self.connect_recursive(root, None)

        return root

# @lc code=end

