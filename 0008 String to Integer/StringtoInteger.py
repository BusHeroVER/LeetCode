class Solution:
    def myAtoi(self, s: str) -> int:
        ret = 0
        isReading = False
        isPostive = True
        
        for _, c in enumerate(s):
            if c == ' ':
                if isReading:
                    return ret
            elif c == '+':
                if isReading:
                    return ret
                else:
                    isReading = True
            elif c == '-':
                if isReading:
                    return ret
                else:
                    isReading = True
                    isPostive = False
            elif '0' <= c and c <= '9':
                isReading = True

                ret *= 10
                if isPostive:
                    ret += ord(c) - 48
                else:
                    ret -= ord(c) - 48

                if ret < -2147483648:
                    return -2147483648
                elif 2147483647 < ret:
                    return 2147483647
            else:
                return ret
        return ret