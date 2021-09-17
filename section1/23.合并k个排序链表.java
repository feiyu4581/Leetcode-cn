/*
 * @lc app=leetcode.cn id=23 lang=java
 *
 * [23] 合并K个排序链表
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


class ListNodeStack {
    public int length;
    public ListNode[] stack;

    private int parent(int position) {
        return position / 2;
    }

    private int left(int position) {
        return 2 * position;
    }

    private int right(int position) {
        return 2 * position + 1;
    }

    private void lowHeapify(int position) {
        if (position > 1) {
            ListNode parentNode = stack[parent(position)], currentNode = stack[position];
            if (parentNode.val > currentNode.val) {
                stack[parent(position)] = currentNode;
                stack[position] = parentNode;
                lowHeapify(parent(position));
            }
        }
    }

    private void swap(int left, int right) {
        ListNode temp = stack[left];
        stack[left] = stack[right];
        stack[right] = temp;
    }

    private void highHeapify(int position) {
        if (position < length) {
            boolean changeLeft = false, changeRight = false;
            if (left(position) < length && stack[left(position)].val < stack[position].val) {
                changeLeft = true;
            }

            if (right(position) < length && stack[right(position)].val < stack[position].val) {
                if (changeLeft && stack[right(position)].val > stack[left(position)].val) {
                    changeLeft = true;
                } else {
                    changeRight = true;
                }
            }

            if (changeLeft) {
                swap(position, left(position));
                highHeapify(left(position));
            }

            if (changeRight) {
                swap(position, right(position));
                highHeapify(right(position));
            }
        }
    }

    public ListNodeStack(int size) {
        this.length = 1;
        this.stack = new ListNode[size + 1];
    }

    public void push(ListNode node) {
        if (node != null) {
            stack[length] = node;
            lowHeapify(length++);
        }
    }

    public ListNode pop() {
        if (length > 1) {
            ListNode res = stack[1];
            stack[1] = stack[--length];
            highHeapify(1);
            return res;
        }

        return null;
    }

    public boolean isEmpty() {
        return length == 1;
    }
}


class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0) {
            return null;
        } else if (lists.length == 1) {
            return lists[0];
        }

        ListNode dummy = new ListNode(0), current = dummy;
        ListNodeStack stack = new ListNodeStack(lists.length);
        for (int i = 0; i < lists.length; i++) {
            stack.push(lists[i]);
        }

        while (!stack.isEmpty()) {
            ListNode nextNode = stack.pop();
            current.next = new ListNode(nextNode.val);
            current = current.next;

            stack.push(nextNode.next);
        }

        return dummy.next;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.mergeKLists(new ListNode[]{
                ListNode.createNodes(new int[]{1, 4, 5}),
                ListNode.createNodes(new int[]{1, 3, 4}),
                ListNode.createNodes(new int[]{2, 6}),
            })
        );

        System.out.println(
            solution.mergeKLists(new ListNode[]{
                ListNode.createNodes(new int[]{-8,2,4}),
                ListNode.createNodes(new int[]{-9,-9,-9,-9,-8,-5,-3,-2,1}),
                ListNode.createNodes(new int[]{-9,-7,-5,-3,-3,-1,0,3,4}),
                ListNode.createNodes(new int[]{-9,-7,-6,-4,-2,-1,3,4}),
                ListNode.createNodes(new int[]{-10,-3,0}),
                ListNode.createNodes(new int[]{-9,0,4}),
                ListNode.createNodes(new int[]{-8,-8})
            })
        );
    }
}
// @lc code=end

