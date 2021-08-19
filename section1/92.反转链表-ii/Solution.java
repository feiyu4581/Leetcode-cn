import javax.sound.midi.Soundbank;

/*
 * @lc app=leetcode.cn id=92 lang=java
 *
 * [92] 反转链表 II
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
    public ListNode reverseBetween(ListNode head, int left, int right) {
        ListNode start = head;
        int currentIndex = 1;

        ListNode startNode = null, lastNode = null, currentNode = head;
        while (currentNode != null) {
            ListNode nextNode = currentNode.next;

            if (currentIndex < left) {
                start = currentNode;
            } else if (currentIndex <= right) {
                currentNode.next = lastNode;
                lastNode = currentNode;
            }

            if (currentIndex == left) {
                startNode = currentNode;
            }

            if (currentIndex == right) {
                startNode.next = nextNode;

                if (startNode == head) {
                    head = currentNode;
                } else {
                    start.next = currentNode;
                }
                break;
            }

            currentNode = nextNode;
            currentIndex++;
        }

        return head;
    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        // [1, 4, 3, 2, 5]
        System.out.println(solution.reverseBetween(ListNode.createNodes(new int[]{1,2,3,4,5}), 2, 4));
        // [5]
        System.out.println(solution.reverseBetween(ListNode.createNodes(new int[]{5}), 1, 1));
        System.out.println(solution.reverseBetween(ListNode.createNodes(new int[]{3, 5}), 1, 1));
    }
}
// @lc code=end

