/*
 * @lc app=leetcode.cn id=61 lang=java
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 /*
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
    public ListNode rotateRight(ListNode head, int k) {
        if (head == null) {
            return head;
        }

        ListNode startNode = head;
        int i = 0;
        while (i < k) {
            if (startNode.next != null) {
                startNode = startNode.next;
                i++;
            } else {
                k = k % (i + 1);
                i = 0;
                startNode = head;
            }
        }

        ListNode lastNode = head;
        while (startNode.next != null) {
            startNode = startNode.next;
            lastNode = lastNode.next;
        }

        startNode.next = head;
        head = lastNode.next;
        lastNode.next = null;
        return head;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.rotateRight(ListNode.createNodes(new int[]{1, 2, 3, 4, 5}), 2));
        System.out.println(solution.rotateRight(ListNode.createNodes(new int[]{0, 1, 2}), 4));
    }
}
// @lc code=end