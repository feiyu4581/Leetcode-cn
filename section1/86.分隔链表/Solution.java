
/*
 * @lc app=leetcode.cn id=86 lang=java
 *
 * [86] 分隔链表
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

    @Override
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
    public ListNode partition(ListNode head, int x) {
        if (head == null) {
            return head;
        }

        ListNode dummy = new ListNode(0);
        ListNode lastNode = null, current = head, nextNode = dummy;

        while (current != null) {
            if (current.val >= x) {
                nextNode.next = current;
                nextNode = current;

                if (lastNode != null) {
                    lastNode.next = current.next;
                } else {
                    head = current.next;
                }
            } else {
                lastNode = current;
            }

            current = current.next;
        }
        nextNode.next = null;
        if (lastNode != null) {
            lastNode.next = dummy.next;
        }
        return head == null? dummy.next : head;
    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        // [1,2,2,4,3,5]
        System.out.println(solution.partition(ListNode.createNodes(new int[]{1, 4, 3, 2, 5, 2}), 3));

        // [1,2]
        System.out.println(solution.partition(ListNode.createNodes(new int[]{2, 1}), 2));
        // [1]
        System.out.println(solution.partition(ListNode.createNodes(new int[]{1}), 0));
    }
}
// @lc code=end

