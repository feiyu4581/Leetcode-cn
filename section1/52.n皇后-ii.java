import java.util.HashSet;
import java.util.Set;

/*
 * @lc app=leetcode.cn id=52 lang=java
 *
 * [52] N皇后 II
 */

// @lc code=start
class Solution {
    private int count = 0;
    private Set<Integer> cols;
    private Set<Integer> gradients;
    private Set<Integer> positionGradients;

    private void initSolution(int count) {
        this.count = count;
        this.cols = new HashSet<>();
        this.gradients = new HashSet<>();
        this.positionGradients = new HashSet<>();
    }

    private boolean testCell(int row, int col) {
        if (cols.contains(col)) {
            return false;
        }

        if (gradients.contains(row - col)) {
            return false;
        }

        if (positionGradients.contains(-row - col)) {
            return false;
        }

        return true;
    }

    private void addCell(int row, int col) {
        cols.add(col);
        gradients.add(row - col);
        positionGradients.add(-row - col);
    }

    private void removeCell(int row, int col) {
        cols.remove(col);
        gradients.remove(row - col);
        positionGradients.remove(-row - col);
    }

    private int totalNQueensRecursive(int row) {
        if (row >= this.count) {
            return 1;
        }

        int res = 0;
        for (int col = 0; col < this.count; col++) {
            if (testCell(row, col)) {
                addCell(row, col);
                res += totalNQueensRecursive(row + 1);
                removeCell(row, col);
            }
        }

        return res;
    }

    public int totalNQueens(int n) {
        initSolution(n);
        return totalNQueensRecursive(0);
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.totalNQueens(4) == 2
        );

        System.out.println(
            solution.totalNQueens(1) == 1
        );
    }
}
// @lc code=end

