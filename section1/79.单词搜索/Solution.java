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
    Map<String, Boolean> cache;

    private boolean existRecursive(HashSet<String> histories, int row, int column, int index) {
        String position = row + "/" + column;
        List<String> cacheKey = new ArrayList<>(histories);
        cacheKey.add(position);
//        cacheKey.sort(Comparator.naturalOrder());
        String currentCacheStr = cacheKey.toString();

        if (!cache.containsKey(currentCacheStr)) {
            boolean cacheRes = false;
            if (index >= word.length()) {
                cacheRes = true;
            } else if (row < 0 || column < 0 || row >= board.length || column >= board[0].length) {
                return false;
            } else if (board[row][column] == word.charAt(index)) {
                if (histories.contains(position)) {
                    return false;
                }

                histories.add(position);
                cacheRes = existRecursive(histories, row - 1, column, index + 1) ||
                        existRecursive(histories, row, column - 1, index + 1) ||
                        existRecursive(histories, row, column + 1, index + 1) ||
                        existRecursive(histories, row + 1, column, index + 1);

                histories.remove(position);
            }

            cache.put(currentCacheStr, cacheRes);

            if (!cacheRes) {
                if (column < board[0].length - 1) {
                    column++;
                } else {
                    row++;
                    column = 0;
                }
               return existRecursive(new LinkedHashSet<>(), row, column, 0);
            }
        }

        return cache.get(currentCacheStr);
    }

    public boolean exist(char[][] board, String word) {
        this.board = board;
        this.word = word;
        this.cache = new HashMap<>();
        return existRecursive(new LinkedHashSet<>(), 0, 0, 0);
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
//        System.out.println(solution.exist(new char[][]{
//                new char[]{'A','B','C','E'},
//                new char[]{'S','F','C','S'},
//                new char[]{'A','D','E','E'}
//        }, "ABCCED"));
//
//        System.out.println(solution.exist(new char[][]{
//                new char[]{'A','B','C','E'},
//                new char[]{'S','F','C','S'},
//                new char[]{'A','D','E','E'}
//        }, "SEE"));
//
//        System.out.println(!solution.exist(new char[][]{
//                new char[]{'A','B','C','E'},
//                new char[]{'S','F','C','S'},
//                new char[]{'A','D','E','E'}
//        }, "ABCB"));
//
//        System.out.println(!solution.exist(new char[][]{
//                new char[]{'a','b'},
//                new char[]{'c','d'}
//        }, "abcd"));
//
//        System.out.println(solution.exist(new char[][]{
//                new char[]{'C','A', 'A'},
//                new char[]{'A','A', 'A'},
//                new char[]{'B','C', 'D'}
//        }, "AAB"));
//
//        System.out.println(solution.exist(new char[][]{
//                new char[]{'a'}
//        }, "a"));

        System.out.println(solution.exist(new char[][]{
                new char[]{'A','B','C','E'},
                new char[]{'S','F','E','S'},
                new char[]{'A','D','E','E'}
        }, "ABCESEEEFS"));
    }
}
// @lc code=end

