import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=12 lang=java
 *
 * [12] 整数转罗马数字
 */

// @lc code=start

class Roman<A, B> {
    public A symbol;
    public B num;

    public Roman(A symbol, B num) {
        this.symbol = symbol;
        this.num = num;
    }
}

class Solution {
    private ArrayList<Roman<String, Integer>> romanMaps = new ArrayList<>();

    public Solution() {
        romanMaps.add(new Roman<String, Integer>("M", 1000));
        romanMaps.add(new Roman<String, Integer>("CM", 900));
        romanMaps.add(new Roman<String, Integer>("D", 500));
        romanMaps.add(new Roman<String, Integer>("CD", 400));
        romanMaps.add(new Roman<String, Integer>("C", 100));
        romanMaps.add(new Roman<String, Integer>("XC", 90));
        romanMaps.add(new Roman<String, Integer>("L", 50));
        romanMaps.add(new Roman<String, Integer>("XL", 40));
        romanMaps.add(new Roman<String, Integer>("X", 10));
        romanMaps.add(new Roman<String, Integer>("IX", 9));
        romanMaps.add(new Roman<String, Integer>("V", 5));
        romanMaps.add(new Roman<String, Integer>("IV", 4));
        romanMaps.add(new Roman<String, Integer>("I", 1));
    }

    public String intToRoman(int num) {
        StringBuilder builder = new StringBuilder();
        int index = 0;
        while (num > 0 && index < romanMaps.size()) {
            if (romanMaps.get(index).num <= num) {
                num -= romanMaps.get(index).num;
                builder.append(romanMaps.get(index).symbol);
            } else {
                index += 1;
            }
        }

        return builder.toString();
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(
            solution.intToRoman(3).equals("III")
        );
        System.out.println(
            solution.intToRoman(4).equals("IV")
        );
        System.out.println(
            solution.intToRoman(9).equals("IX")
        );
        System.out.println(
            solution.intToRoman(58).equals("LVIII")
        );
        System.out.println(
            solution.intToRoman(1994).equals("MCMXCIV")
        );
    }
}
// @lc code=end

