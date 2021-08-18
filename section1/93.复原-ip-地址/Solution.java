/*
 * @lc app=leetcode.cn id=93 lang=java
 *
 * [93] 复原 IP 地址
 */

import java.util.ArrayList;
import java.util.List;

// @lc code=start
class Solution {
    char[] nums;

    public List<String> restoreIpAddresses(String s) {
        List<String> res = new ArrayList<>();

        this.nums = s.toCharArray();
        restoreIpAddressesRecursive(res, new char[s.length() + 3], 0, 0);

        return res;
    }

    private void restoreIpAddressesRecursive(List<String> res, char[] current, int start, int length) {
        if (length == 4 && start == nums.length) {
            res.add(String.valueOf(current));
        } else if ((4 - length) * 3 >= (nums.length - start) && length < 4 && start < nums.length) {
            if (length >= 1) {
                current[start + length - 1] = '.';
            }
            if (nums[start] == '0') {
                current[start + length] = '0';
                restoreIpAddressesRecursive(res, current, start + 1, length + 1);
            } else {
                current[start + length] = nums[start];
                restoreIpAddressesRecursive(res, current, start + 1, length + 1);

                if (start < nums.length - 1) {
                    current[start + length + 1] = nums[start + 1];
                    restoreIpAddressesRecursive(res, current, start + 2, length + 1);
                }

                if (start < nums.length - 2) {
                    if (nums[start] < '2' || (nums[start] == '2' && (nums[start + 1] < '5' || (nums[start + 1] == '5' && nums[start + 2] <= '5')))) {
                        current[start + length + 2] = nums[start + 2];
                        restoreIpAddressesRecursive(res, current, start + 3, length + 1);
                    }
                }
            }
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        // ["255.255.11.135","255.255.111.35"]
        System.out.println(solution.restoreIpAddresses("25525511135"));
        // ["0.0.0.0"]
        System.out.println(solution.restoreIpAddresses("0000"));
        // ["1.1.1.1"]
        System.out.println(solution.restoreIpAddresses("1111"));
        // ["0.10.0.10","0.100.1.0"]
        System.out.println(solution.restoreIpAddresses("010010"));
        // ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
        System.out.println(solution.restoreIpAddresses("101023"));
        // ["17.216.25.41","17.216.254.1","172.16.25.41","172.16.254.1","172.162.5.41","172.162.54.1"]
        System.out.println(solution.restoreIpAddresses("172162541"));
    }
}
// @lc code=end

