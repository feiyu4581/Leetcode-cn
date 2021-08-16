/*
 * @lc app=leetcode.cn id=84 lang=java
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
class Solution {
    private int[] heights;

    private int findMinIndex(int left, int right) {
        int minIndex = left, maxIndex = left;
        for (int i = left + 1; i < right; i++) {
            if (heights[i] < heights[minIndex]) {
                minIndex = i;
            } else if (heights[i] > heights[maxIndex]) {
                maxIndex = i;
            }
        }

        return heights[minIndex] == heights[maxIndex]? -1 : minIndex;
    }

    private int rectangleArea(int left, int right) {
        if (left >= right) {
            return 0;
        }

        if (right == left - 1) {
            return heights[left];
        }

        int minIndex = findMinIndex(left, right);
        if (minIndex == -1) {
            return (right - left) * heights[left];
        }

        return Math.max(
                Math.max(rectangleArea(left, minIndex), rectangleArea(minIndex + 1, right)),
                (right - left) * heights[minIndex]
        );
    }

    public int largestRectangleArea(int[] heights) {
        this.heights = heights;
        return rectangleArea(0, heights.length);
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.largestRectangleArea(new int[]{2, 1, 5, 6, 2, 3}) == 10);
        System.out.println(solution.largestRectangleArea(new int[]{2, 4}) == 4);
    }
}
// @lc code=end

