#
# @lc app=leetcode.cn id=109 lang=python3
#
# [109] 有序链表转换二叉搜索树
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
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


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def __str__(self):
        return f'ListNode(val={self.val})'

    @staticmethod
    def generate(nums):
        head = ListNode(nums[0])
        node = head
        for num in nums[1:]:
            node.next = ListNode(num)
            node = node.next

        return head

class Solution:
    def find_mid(self, head, end):
        mid, fast = head, head
        while fast and fast != end and fast.next and fast.next != end:
            mid = mid.next
            fast = fast.next.next

        return mid

    def list_to_bst(self, head, end):
        if not head or head == end:
            return None

        mid = self.find_mid(head, end)
        node = TreeNode(mid.val)
        node.left = self.list_to_bst(head, mid)
        node.right = self.list_to_bst(mid.next, end)
        return node

    def sortedListToBST(self, head: ListNode) -> TreeNode:
        return self.list_to_bst(head, None)

    def test(self):
        print(self.sortedListToBST(ListNode.generate([-10, -3, 0, 5, 9])).show())


Solution().test()
# @lc code=end

