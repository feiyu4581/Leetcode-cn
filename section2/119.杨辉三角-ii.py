#
# @lc app=leetcode.cn id=119 lang=python3
#
# [119] 杨辉三角 II
#

# @lc code=start
from typing import List


class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        prev_row = []
        for index in range(1, rowIndex + 2):
            row = [1] * index
            for column in range(1, index - 1):
                row[column] = prev_row[column - 1] + prev_row[column]

            prev_row = row

        return prev_row

    def test(self):
        print(self.getRow(3) == [1, 3, 3, 1])
        print(self.getRow(0) == [1])
        print(self.getRow(1) == [1, 1])


# Solution().test()

# @lc code=end
