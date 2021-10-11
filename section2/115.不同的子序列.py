#
# @lc app=leetcode.cn id=115 lang=python3
#
# [115] 不同的子序列
#

# @lc code=start
class Solution:
    def num_distinct_recursive(self, s, left, t, right, mems):
        key = f'{left}-{right}'
        if key not in mems:
            if right == len(t):
                num = 1
            elif left >= len(s) or right >= len(t):
                num = 0
            else:
                num = self.num_distinct_recursive(s, left + 1, t, right, mems)
                if s[left] == t[right]:
                    num += self.num_distinct_recursive(s, left + 1, t, right + 1, mems)

            mems[key] = num

        return mems[key]

    def numDistinct(self, s: str, t: str) -> int:
        return self.num_distinct_recursive(s, 0, t, 0, {})

    def test(self):
        print(self.numDistinct('rabbbit', 'rabbit') == 3)
        print(self.numDistinct('babgbag', 'bag') == 5)


# Solution().test()

# @lc code=end
