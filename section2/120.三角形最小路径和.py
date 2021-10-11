#
# @lc app=leetcode.cn id=120 lang=python3
#
# [120] 三角形最小路径和
#

# @lc code=start
from typing import List


class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        memo = [0] * len(triangle)
        for width, rows in enumerate(triangle):
            current_memo = [0] * len(triangle)
            for index, column in enumerate(rows):
                left_column = memo[index - 1] if index > 0 else memo[index]
                right_column = memo[index] if index < width else memo[index - 1]

                current_memo[index] = min(left_column, right_column) + column

            memo = current_memo

        return min(memo)

    def test(self):
        print(self.minimumTotal([[2], [3, 4], [6, 5, 7], [4, 1, 8, 3]]) == 11)
        print(self.minimumTotal([[-10]]) == -10)


Solution().test()

# @lc code=end
