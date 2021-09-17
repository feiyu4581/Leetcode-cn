import java.util.*;

/*
 * @lc app=leetcode.cn id=79 lang=java
 *
 * [79] 单词搜索
 */

// @lc code=start
class Solution {
    char[][] board;
    String word;
    Set<String> histories;

    private boolean existRecursive(int row, int column, int index) {
        if (index >= word.length()) {
            return true;
        } else if (row < 0 || column < 0 || row >= board.length || column >= board[0].length) {
            return false;
        } else if (board[row][column] == word.charAt(index)) {
            String position = row + "/" + column;
            if (histories.contains(position)) {
                return false;
            }

            histories.add(position);
            boolean res = existRecursive(row - 1, column, index + 1) ||
                    existRecursive(row, column - 1, index + 1) ||
                    existRecursive(row, column + 1, index + 1) ||
                    existRecursive(row + 1, column, index + 1);

            histories.remove(position);
            return res;
        } else {
            return false;
        }
    }

    public boolean exist(char[][] board, String word) {
        this.board = board;
        this.word = word;
        this.histories = new HashSet<>();
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if (existRecursive(i, j, 0)) {
                    return true;
                }
            }
        }
        return false;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.exist(new char[][]{
                new char[]{'A','B','C','E'},
                new char[]{'S','F','C','S'},
                new char[]{'A','D','E','E'}
        }, "ABCCED"));

        System.out.println(solution.exist(new char[][]{
                new char[]{'A','B','C','E'},
                new char[]{'S','F','C','S'},
                new char[]{'A','D','E','E'}
        }, "SEE"));

        System.out.println(!solution.exist(new char[][]{
                new char[]{'A','B','C','E'},
                new char[]{'S','F','C','S'},
                new char[]{'A','D','E','E'}
        }, "ABCB"));

        System.out.println(!solution.exist(new char[][]{
                new char[]{'a','b'},
                new char[]{'c','d'}
        }, "abcd"));

        System.out.println(solution.exist(new char[][]{
                new char[]{'C','A', 'A'},
                new char[]{'A','A', 'A'},
                new char[]{'B','C', 'D'}
        }, "AAB"));

        System.out.println(solution.exist(new char[][]{
                new char[]{'a'}
        }, "a"));

        System.out.println(solution.exist(new char[][]{
                new char[]{'A','B','C','E'},
                new char[]{'S','F','E','S'},
                new char[]{'A','D','E','E'}
        }, "ABCESEEEFS"));
    }
}
// @lc code=end

