class Solution:
    def convert(self, s: str, numRows: int) -> str:
        length = len(s)
        if length <= numRows or numRows < 2:
            return s
        
        ret = ""
        for r in range(numRows):
            index = r
            offset0 = 2 * (numRows - r - 1)
            offset1 = 2 * r

            ret += s[index]
            while index < length:
                if offset0 > 0:
                    index += offset0
                    if index < length:
                        ret += s[index]
                if offset1 > 0:
                    index += offset1
                    if index < length:
                        ret += s[index]
        return ret