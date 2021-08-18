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
}
// @lc code=end

