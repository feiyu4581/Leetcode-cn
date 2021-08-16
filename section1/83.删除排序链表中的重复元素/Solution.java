/*
 * @lc app=leetcode.cn id=83 lang=java
 *
 * [83] 删除排序链表中的重复元素
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
    public ListNode deleteDuplicates(ListNode head) {
        ListNode dummy = new ListNode(0);
        ListNode start = head, current = dummy;

        ListNode lastNode = null;
        while (start != null) {
            if (lastNode == null || start.val != lastNode.val) {
                current.next = start;
                current = start;
            }

            lastNode = start;
            start = start.next;
        }

        current.next = null;
        return dummy.next;
    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        // 1, 2
        System.out.println(solution.deleteDuplicates(ListNode.createNodes(new int[]{1, 1, 2})));
        // 1, 2, 3
        System.out.println(solution.deleteDuplicates(ListNode.createNodes(new int[]{1,1,2,3,3})));
    }
}
// @lc code=end

