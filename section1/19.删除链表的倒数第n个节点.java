import java.nio.channels.Pipe;

/*
 * @lc app=leetcode.cn id=19 lang=java
 *
 * [19] 删除链表的倒数第N个节点
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
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode dummy = new ListNode(0);
        dummy.next = head;
        ListNode current = dummy, pionner = dummy;
        for (int i = 0; i < n; i++) {
            pionner = pionner.next;
        }

        while (pionner != null && pionner.next != null) {
            current = current.next;
            pionner = pionner.next;
        }

        current.next = current.next.next;
        return dummy.next;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.removeNthFromEnd(ListNode.createNodes(new int[]{1, 2, 3, 4, 5}), 2)
        );

        System.out.println(
            solution.removeNthFromEnd(ListNode.createNodes(new int[]{1}), 1)
        );
    }
}
// @lc code=end

