import java.text.DecimalFormat;
import java.util.HashMap;

/*
 * @lc app=leetcode.cn id=50 lang=java
 *
 * [50] Pow(x, n)
 */

// @lc code=start
class Solution {
    public double myPow(double x, int n) {
        if (n == 0) {
            return 1;
        }

        if (n == Integer.MIN_VALUE && x > 1) {
            return 0;
        }

        if (x == 0) {
            return 0;
        }

        if (n < 0) {
            x = 1 / x;
            n = -n;
        }
        System.out.println(n);

        double res = 1L;
        int current = 1, left = n;
        HashMap<Integer, Double> caches = new HashMap<>();
        caches.put(current, x);
        while (left > 0) {
            if (current > left) {
                current = current / 2;
            } else if (current < left) {
                if (current * 2 < left) {
                    caches.put(current * 2, caches.get(current) * caches.get(current));
                    current *= 2;
                }
            }

            if (current <= left) {
                res *= caches.get(current);
                left -= current;
            }

        }

        DecimalFormat dl = new DecimalFormat("0.00000");
        String strRes = dl.format(res);
        return Double.parseDouble(strRes);
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        // System.out.println(
        //     solution.myPow(2.00000, 10) == 1024.00000
        // );

        // System.out.println(
        //     solution.myPow(2.10000, 3) == 9.26100
        // );

        // System.out.println(
        //     solution.myPow(2.00000, -2) == 0.25000
        // );

        System.out.println(
            solution.myPow(2.00000, -2147483648)
        );
    }
}
// @lc code=end

