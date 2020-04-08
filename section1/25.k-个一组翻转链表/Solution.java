import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=25 lang=java
 *
 * [25] K 个一组翻转链表
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
    private ListNode reverseGroup(ListNode head, ListNode tail) {
        if (head == tail) {
            return head;
        }

        ListNode nextNode = head.next, current = head;
        head.next = tail.next;
        while (nextNode != tail) {
            ListNode temp = nextNode;
            nextNode = nextNode.next;
            temp.next = current;
            current = temp;
        }

        tail.next = current;
        return tail;
    }

    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode dummy = new ListNode(0);
        dummy.next = head;
        ListNode current = dummy, start = head, end = head;
        while (end != null) {
            for (int i = 1; i < k && end != null; i++) {
                end = end.next;
            }

            if (end != null) {
                ListNode temp = reverseGroup(start, end);
                current.next = temp;
                current = start;
                start = start.next;
                end = start;
            }
        }

        return dummy.next;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.reverseKGroup(ListNode.createNodes(new int[]{1, 2, 3, 4, 5}), 2)
        );

        System.out.println(
            solution.reverseKGroup(ListNode.createNodes(new int[]{1, 2, 3, 4, 5}), 3)
        );
    }
}
// @lc code=end

