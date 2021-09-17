package Solution;/*
 * @lc app=leetcode.cn id=85 lang=java
 *
 * [85] 最大矩形
 */

// @lc code=start

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Edge {
    private final int start;
    private final int end;

    public Edge(int start, int end) {
        this.start = start;
        this.end = end;
    }

    public int getStart() {
        return start;
    }

    public int getEnd() {
        return end;
    }

    public int getWidth() {
        return end - start + 1;
    }

    public boolean isMixed(Edge other) {
        return other.getEnd() >= this.getStart() && other.getStart() <= this.getEnd();
    }

    public Edge mixedEdge(Edge other) {
        return new Edge(Math.max(this.getStart(), other.getStart()),  Math.min(this.getEnd(), other.getEnd()));
    }

    @Override
    public String toString() {
        return start + "/" + end;
    }
}

class Solution {
    List<List<Edge>> edges;
    Map<String, Integer> caches;
    private static final char POINT = '1';

    private int getMaxRectangle(Edge edge, int row, int length) {
        int rectangle = edge.getWidth();
        if (row < edges.size()) {
            for (Edge rowEdge : edges.get(row)) {
                if (rowEdge.getStart() > edge.getEnd()) {
                    break;
                }

                if (edge.isMixed(rowEdge)) {
                    Edge newEdge = edge.mixedEdge(rowEdge);
                    rectangle = Math.max(
                        Math.max(rectangle, newEdge.getWidth() * (length + 1)),
                        getMaxRectangle(newEdge, row + 1, length + 1)
                    );
                }
            }
        }

        return rectangle;
    }

    public int maximalRectangle(char[][] matrix) {
        this.edges = new ArrayList<>();
        this.caches = new HashMap<>();
        for (int i = 0; i < matrix.length; i++) {
            List<Edge> rowEdges = new ArrayList<Edge>();
            int pointNums = 0;
            for (int j = 0; j < matrix[0].length; j++) {
                if (matrix[i][j] == POINT) {
                    pointNums++;
                } else if (pointNums > 0) {
                    rowEdges.add(new Edge(j - pointNums, j - 1));
                    pointNums = 0;
                }
            }
            if (pointNums > 0) {
                rowEdges.add(new Edge(matrix[0].length - pointNums, matrix[0].length - 1));
            }
            this.edges.add(rowEdges);
        }

        int maxInt = 0;
        for (int row = 0; row < matrix.length; row++) {
            for (Edge edge : edges.get(row)) {
                maxInt = Math.max(maxInt, getMaxRectangle(edge, row + 1, 1));
            }
        }

        return maxInt;
    }

    public static void main(String[] ags) {
        Solution solution = new Solution();
        System.out.println(solution.maximalRectangle(new char[][]{
                new char[]{'1','0','1','1','0','1'},
                new char[]{'1','1','1','1','1','1'},
                new char[]{'0','1','1','0','1','1'},
                new char[]{'1','1','1','0','1','0'},
                new char[]{'0','1','1','1','1','1'},
                new char[]{'1','1','0','1','1','1'},
        }) == 8);

        System.out.println(solution.maximalRectangle(new char[][]{
                new char[]{'0', '0', '1', '0'},
                new char[]{'1', '1', '1', '1'},
                new char[]{'1', '1', '1', '1'},
                new char[]{'1', '1', '1', '0'},
                new char[]{'1', '1', '0', '0'},
                new char[]{'1', '1', '1', '1'},
                new char[]{'1', '1', '1', '0'}
        }) == 12);

        System.out.println(solution.maximalRectangle(new char[][]{
                new char[]{'1','0','1','0','0'},
                new char[]{'1','0','1','1','1'},
                new char[]{'1','1','1','1','1'},
                new char[]{'1','0','0','1','0'}
        }) == 6);

        System.out.println(solution.maximalRectangle(new char[][]{
                new char[]{'0','0','0','0','0','0','1'},
                new char[]{'0','0','0','0','1','1','1'},
                new char[]{'1','1','1','1','1','1','1'},
                new char[]{'0','0','0','1','1','1','1'}
        }) == 9);

        System.out.println(solution.maximalRectangle(new char[][]{new char[]{'0'}}) == 0);
        System.out.println(solution.maximalRectangle(new char[][]{new char[]{'1'}}) == 1);
        System.out.println(solution.maximalRectangle(new char[][]{new char[]{'0', '0'}}) == 0);
    }
}
// @lc code=end

