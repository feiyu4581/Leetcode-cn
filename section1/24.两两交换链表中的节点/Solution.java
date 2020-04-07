/*
 * @lc app=leetcode.cn id=24 lang=java
 *
 * [24] 两两交换链表中的节点
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
    public ListNode swapPairs(ListNode head) {
        ListNode dummy = new ListNode(0), current = dummy, prefix = head, suffix = head;
        dummy.next = head;

        while (prefix != null && prefix.next != null) {
            suffix = prefix.next;
            prefix.next = suffix.next;
            suffix.next = prefix;
            current.next = suffix;

            current = prefix;
            prefix = prefix.next;
        }

        return dummy.next;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.swapPairs(ListNode.createNodes(new int[]{1, 2, 3, 4}))
        );
    }
}
// @lc code=end

