class Solution:
    def newMod(self, a, b):
        res = a % b
        return res if not res else res - b if a < 0 else res
    def max(self):
        return 214748364
    def min(self):
        return -214748364

    def reverse(self, x: int) -> int:
        ret = 0

        while x != 0:
            digit = self.newMod(x, 10)
            
            if ret > self.max() or (ret == self.max() and digit > 7):
                return 0
            if ret < self.min() or (ret == self.min() and digit < -8):
                return 0
            
            ret = ret * 10 + digit
            x = int(x / 10)
        return ret