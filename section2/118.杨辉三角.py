#
# @lc app=leetcode.cn id=118 lang=python3
#
# [118] 杨辉三角
#

# @lc code=start
from typing import List


class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        res = []
        for index in range(1, numRows + 1):
            row = [1] * index
            for column in range(1, index - 1):
                row[column] = res[-1][column - 1] + res[-1][column]

            res.append(row)

        return res

    def test(self):
        print(self.generate(5) == [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]])
        print(self.generate(1) == [[1]])


Solution().test()

# @lc code=end
