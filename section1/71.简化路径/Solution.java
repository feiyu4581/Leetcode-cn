import java.io.File;
import java.util.ArrayList;
import java.util.List;


/*
 * @lc app=leetcode.cn id=71 lang=java
 *
 * [71] 简化路径
 */

// @lc code=start
class Solution {

    private static Integer POINT = 1;
    private static Integer FILE = 2;
    private static Integer OTHER = 3;

    public String simplifyPath(String path) {
        List<String> files = new ArrayList<>();
        StringBuilder file = new StringBuilder();
        int index = 0;
        int lastFlag = OTHER;

        char c;
        while (index <= path.length()) {
            if (index == path.length()) {
                c = '/';
            } else {
                c = path.charAt(index);
            }

            int currentFlag = OTHER;
            if (c == '.') {
                currentFlag = POINT;
                if (lastFlag == FILE) {
                    currentFlag = FILE;
                }
            } else if (c != '/') {
                currentFlag = FILE;
                if (lastFlag == POINT) {
                    lastFlag = FILE;
                }
            }

            if (currentFlag != lastFlag) {
                if (lastFlag == POINT) {
                    if (file.length() == 2 && !files.isEmpty()) {
                        files.remove(files.size() - 1);
                    } else if (file.length() > 2) {
                        files.add(file.toString());
                    }
                } else if (lastFlag == FILE) {
                    files.add(file.toString());
                }
                file.delete(0, file.length());
                lastFlag = currentFlag;
            }

            if (currentFlag != OTHER) {
                file.append(c);
            }

            index++;
        }

        return "/" + String.join("/", files);
    }
    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.simplifyPath("/home/").equals("/home"));
        System.out.println(solution.simplifyPath("/../").equals("/"));
        System.out.println(solution.simplifyPath("/home//foo/").equals("/home/foo"));
        System.out.println(solution.simplifyPath("/a/./b/../../c/").equals("/c"));
        System.out.println(solution.simplifyPath("/..hidden").equals("/..hidden"));
        System.out.println(solution.simplifyPath("/hello../world").equals("/hello../world"));
    }
}
// @lc code=end

