class Solution {
public:
    string convert(string s, int numRows) {
        int len = s.size();
        if (len <= numRows || numRows < 2) {
            return s;
        }

        std::string ret;
        for (int r = 0; r < numRows; r++) {
            int index = r;
            int offset0 = 2 * (numRows - r - 1);
            int offset1 = 2 * r;

            ret.push_back(s[index]);
            for (; index < len;) {
                if (offset0 > 0) {
                    index += offset0;
                    if (index < len) {
                        ret.push_back(s[index]);
                    }
                }
                if (offset1 > 0) {
                    index += offset1;
                    if (index < len) {
                        ret.push_back(s[index]);
                    }
                }
            }
        }
        return ret;
    }
};