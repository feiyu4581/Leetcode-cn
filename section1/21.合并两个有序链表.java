/*
 * @lc app=leetcode.cn id=21 lang=java
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }

    public static ListNode createNodes(int[] nums) {
        if (nums.length == 0) {
            return null;
        }

        ListNode res = new ListNode(nums[0]);
        ListNode current = res;
        for (int i = 1; i < nums.length; i++) {
            current.next = new ListNode(nums[i]);
            current = current.next;
        }

        return res;
    }

    public String toString() {
        String res = "";
        ListNode current = this;
        while (current != null) {
            res += current.val;
            current = current.next;
        }

        return res;
    }

}

class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(0), current = dummy;
        while (l1 != null || l2 != null) {
            int left = l1 != null? l1.val : Integer.MAX_VALUE;
            int right = l2 != null? l2.val : Integer.MAX_VALUE;

            if (left <= right) {
                current.next = new ListNode(left);
                l1 = l1.next;
            } else {
                current.next = new ListNode(right);
                l2 = l2.next;
            }

            current = current.next;
        }
        return dummy.next;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.mergeTwoLists(
                ListNode.createNodes(new int[]{1, 2, 4}),
                ListNode.createNodes(new int[]{1, 3, 4}))
        );
    }
}
// @lc code=end

