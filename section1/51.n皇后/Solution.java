import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.List;

/*
 * @lc app=leetcode.cn id=51 lang=java
 *
 * [51] N皇后
 */

// @lc code=start
class Solution {
    private List<Integer> rows;
    private List<Integer> cols;
    private List<Integer> gradients;
    private List<Integer> positionGradients;
    private List<Integer> stacks;

    private void initSolution() {
        rows = new ArrayList<>();
        cols = new ArrayList<>();
        gradients = new ArrayList<>();
        positionGradients = new ArrayList<>();
        stacks = new ArrayList<>();
    }

    private int getStackCol() {
        if (stacks.size() > 0) {
            int col = stacks.get(stacks.size() - 1) + 1;
            stacks.remove(stacks.size() - 1);
            rows.remove(rows.size() - 1);
            cols.remove(cols.size() - 1);
            gradients.remove(gradients.size() - 1);
            positionGradients.remove(positionGradients.size() - 1);

            return col;
        }

        return 0;
    }

    private void addStacks(int row, int col) {
        stacks.add(col);
        rows.add(row);
        cols.add(col);
        gradients.add(row - col);
        positionGradients.add(-row - col);
    }

    public List<List<String>> solveNQueens(int n) {
        int col = 0, row = 0;
        List<List<String>> res = new ArrayList<>();
        initSolution();

        while (row >= 0) {
            if (row >= n) {
                List<String> strRes = new ArrayList<>();
                for (Integer colRes : stacks) {
                    char[] charRes = new char[n];
                    for (int i = 0; i < n; i++) {
                        charRes[i] = i == colRes? 'Q' : '.';
                    }
                    strRes.add(String.valueOf(charRes));
                }

                res.add(strRes);

                row -= 1;
                col = getStackCol();
            } else if (col >= n) {
                col = getStackCol();
                row -= 1;
            } else {
                if (testQueue(row, col)) {
                    addStacks(row, col);
                    row += 1;
                    col = 0;
                } else {
                    col += 1;
                }
            }
        }

        return res;
    }

    private boolean testQueue(int row, int col) {
        if (rows.contains(row)) {
            return false;
        }

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

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.print(
            solution.solveNQueens(4)  // 1 - 3 - 0 - 2
        );
    }
}
// @lc code=end

