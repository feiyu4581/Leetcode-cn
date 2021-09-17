import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

/*
 * @lc app=leetcode.cn id=36 lang=java
 *
 * [36] 有效的数独
 */

// @lc code=start
class Solution {
    public boolean isValidSudoku(char[][] board) {
        List<HashSet<Character>> rowNums = new ArrayList<>();
        List<HashSet<Character>> columnNums = new ArrayList<>();
        List<HashSet<Character>> subNums = new ArrayList<>();

        for (int i = 0; i < 9; i++) {
            rowNums.add(new HashSet<>());
            columnNums.add(new HashSet<>());
            subNums.add(new HashSet<>());
        }

        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                if (board[i][j] != '.') {
                    int subIndex = (i / 3) * 3 + (j / 3);
                    char current = board[i][j];
                    if (rowNums.get(i).contains(current) || columnNums.get(j).contains(current) || subNums.get(subIndex).contains(current)) {
                        return false;
                    }

                    rowNums.get(i).add(current);
                    columnNums.get(j).add(current);
                    subNums.get(subIndex).add(current);
                }
            }
        }
        return true;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.isValidSudoku(new char[][]{
                new char[]{'5','3','.','.','7','.','.','.','.'},
                new char[]{'6','.','.','1','9','5','.','.','.'},
                new char[]{'.','9','8','.','.','.','.','6','.'},
                new char[]{'8','.','.','.','6','.','.','.','3'},
                new char[]{'4','.','.','8','.','3','.','.','1'},
                new char[]{'7','.','.','.','2','.','.','.','6'},
                new char[]{'.','6','.','.','.','.','2','8','.'},
                new char[]{'.','.','.','4','1','9','.','.','5'},
                new char[]{'.','.','.','.','8','.','.','7','9'}
            }) == true
        );

        System.out.println(
            solution.isValidSudoku(new char[][]{
                new char[]{'8','3','.','.','7','.','.','.','.'},
                new char[]{'6','.','.','1','9','5','.','.','.'},
                new char[]{'.','9','8','.','.','.','.','6','.'},
                new char[]{'8','.','.','.','6','.','.','.','3'},
                new char[]{'4','.','.','8','.','3','.','.','1'},
                new char[]{'7','.','.','.','2','.','.','.','6'},
                new char[]{'.','6','.','.','.','.','2','8','.'},
                new char[]{'.','.','.','4','1','9','.','.','5'},
                new char[]{'.','.','.','.','8','.','.','7','9'}
            }) == false
        );
    }
}
// @lc code=end

