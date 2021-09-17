import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=68 lang=java
 *
 * [68] 文本左右对齐
 */

// @lc code=start
class Solution {
    private String fullJustifyLine(String[] words, int left, int right, int maxWidth) {
        char[] blankStrs = new char[maxWidth];
        int index = 0;
        if (right - left == 1) {
            while (index < words[left].length()) {
                blankStrs[index] = words[left].charAt(index);
                index++;
            }

            while (index < maxWidth) {
                blankStrs[index++] = ' ';
            }
        } else {
            int midLength = 0;
            for (int i = left + 1; i < right - 1; i++) {
                midLength += words[i].length();
            }

            int midWidth = maxWidth - words[left].length() - words[right - 1].length();
            int blankLength = midWidth - midLength;
            int average = blankLength;
            if (right - left > 2) {
                average = blankLength / (right - left - 1);
            }

            int extraBlank = blankLength - (average * (right - left  - 1));
            while (left < right) {
                for (int i = 0; i < words[left].length(); i++) {
                    blankStrs[index++] = words[left].charAt(i);
                }

                if (left == right - 1) {
                    break;
                }

                for (int i = 0; i < average; i++) {
                    blankStrs[index++] = ' ';
                }

                if (extraBlank > 0) {
                    blankStrs[index++] = ' ';
                    extraBlank--;
                }

                left++;
            }
        }

        return String.valueOf(blankStrs);
    }

    public List<String> fullJustify(String[] words, int maxWidth) {
        List<String> res = new ArrayList<>();

        int start = 0;
        int currentWidth = 0;
        for (int i = 0; i < words.length; i++) {
            if (currentWidth + words[i].length() > maxWidth) {
                res.add(fullJustifyLine(words, start, i, maxWidth));
                start = i;
                currentWidth = 0;
            }

            currentWidth += words[i].length() + 1;
        }

        char[] blankStrs = new char[maxWidth];
        int index = 0;
        while (start < words.length) {
            for (int i = 0; i < words[start].length(); i++) {
                blankStrs[index++] = words[start].charAt(i);
            }

            if (index < maxWidth) {
                blankStrs[index++] = ' ';
            }

            start++;
        }

        while (index < maxWidth) {
            blankStrs[index++] = ' ';
        }

        res.add(String.valueOf(blankStrs));

        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.fullJustify(
                new String[] {"This", "is", "an", "example", "of", "text", "justification."},
                16
        ));
        // [
        //     "This    is    an",
        //     "example  of text",
        //     "justification.  "
        // ]

        System.out.println(solution.fullJustify(
                new String[] {"What","must","be","acknowledgment","shall","be"},
                16
        ));
//         [
//             "What   must   be",
//             "acknowledgment  ",
//             "shall be        "
//         ]

        System.out.println(solution.fullJustify(
                new String[] {"Science","is","what","we","understand","well","enough","to","explain", "to","a","computer.","Art","is","everything","else","we","do"},
                20
        ));
        // [
        //     "Science  is  what we",
        //     "understand      well",
        //     "enough to explain to",
        //     "a  computer.  Art is",
        //     "everything  else  we",
        //     "do                  "
        // ]
    }
}
// @lc code=end

