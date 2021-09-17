import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;

/*
 * @lc app=leetcode.cn id=37 lang=java
 *
 * [37] 解数独
 */

// @lc code=start
class Solution {
    public boolean recursiveSolveSodoku(
        char[][] board,
        List<HashSet<Character>> rowNums,
        List<HashSet<Character>> colNums,
        List<HashSet<Character>> subNums,
        int row,
        int col) {

        if (row > 8) {
            return true;
        }

        if (board[row][col] != '.') {
            return recursiveSolveSodoku(board, rowNums, colNums, subNums, col == 8? row + 1 : row, col == 8? 0 : col + 1);
        }

        int subIndex = (row / 3) * 3 + (col / 3);
        boolean res = false;
        for (int i = 1; i <= 9; i++) {
            char num = (char)(i + '0');
            if (!rowNums.get(row).contains(num) && !colNums.get(col).contains(num) && !subNums.get(subIndex).contains(num)) {
                board[row][col] = num;
                rowNums.get(row).add(num);
                colNums.get(col).add(num);
                subNums.get(subIndex).add(num);
                res = recursiveSolveSodoku(board, rowNums, colNums, subNums, col == 8? row + 1 : row, col == 8? 0 : col + 1);
                if (res) {
                    return res;
                }

                board[row][col] = '.';
                rowNums.get(row).remove(num);
                colNums.get(col).remove(num);
                subNums.get(subIndex).remove(num);
            }
        }

        return res;
    }

    public void solveSudoku(char[][] board) {
        List<HashSet<Character>> rowNums = new ArrayList<>();
        List<HashSet<Character>> colNums = new ArrayList<>();
        List<HashSet<Character>> subNums = new ArrayList<>();

        for (int i = 0; i < 9; i++) {
            rowNums.add(new HashSet<>());
            colNums.add(new HashSet<>());
            subNums.add(new HashSet<>());
        }

        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                if (board[i][j] != '.') {
                    int subIndex = (i / 3) * 3 + (j / 3);
                    char current = board[i][j];
                    rowNums.get(i).add(current);
                    colNums.get(j).add(current);
                    subNums.get(subIndex).add(current);
                }
            }
        }

        recursiveSolveSodoku(board, rowNums, colNums, subNums, 0, 0);

        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                if (board[i][j] == '.') {
                    int subIndex = (i / 3) * 3 + (j / 3);
                    char findNum = '0';
                    for (int num = 1; num <= 9; num++) {
                        char charNum = (char)(num + '0');
                        if (!rowNums.get(i).contains(charNum) && !rowNums.get(j).contains(charNum) && !rowNums.get(subIndex).contains(charNum)) {
                            findNum = charNum;
                            break;
                        }
                    }

                    if (findNum > '0') {
                        board[i][j] = findNum;
                        rowNums.get(i).add(findNum);
                        colNums.get(j).add(findNum);
                        subNums.get(subIndex).add(findNum);
                    } else {
                        return;
                    }
                }
            }
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        char[][] board = new char[][]{
            new char[]{'5','3','.','.','7','.','.','.','.'},
            new char[]{'6','.','.','1','9','5','.','.','.'},
            new char[]{'.','9','8','.','.','.','.','6','.'},
            new char[]{'8','.','.','.','6','.','.','.','3'},
            new char[]{'4','.','.','8','.','3','.','.','1'},
            new char[]{'7','.','.','.','2','.','.','.','6'},
            new char[]{'.','6','.','.','.','.','2','8','.'},
            new char[]{'.','.','.','4','1','9','.','.','5'},
            new char[]{'.','.','.','.','8','.','.','7','9'}
        };

        solution.solveSudoku(board);
        for (int i = 0; i < board.length; i++) {
            System.out.println(Arrays.toString(board[i]));
        }
    }
}
// @lc code=end

