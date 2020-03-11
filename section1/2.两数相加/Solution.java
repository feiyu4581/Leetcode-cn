/*
 * @lc app=leetcode.cn id=2 lang=java
 *
 * [2] 两数相加
 */

// @lc code=start
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
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        if (l1 == null || l2 == null) {
            return null;
        }

        ListNode dummy = new ListNode(0);
        ListNode current = dummy;
        int carry = 0;
        while (l1 != null || l2 != null) {
            int leftVal = l1 == null ? 0 : l1.val;
            int rightVal = l2 == null ? 0 : l2.val;

            int amount = leftVal + rightVal + carry;
            current.next = new ListNode(amount % 10);
            carry = amount / 10;

            current = current.next;
            l1 = l1 == null ? null : l1.next;
            l2 = l2 == null ? null : l2.next;
        }

        if (carry > 0) {
            current.next = new ListNode(carry);
        }

        return dummy.next;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();

        System.out.println(
            solution.addTwoNumbers(
                ListNode.createNodes(new int[]{2, 4, 3}),
                ListNode.createNodes(new int[]{5, 6, 4})
            )
        );
    }
}
// @lc code=end

