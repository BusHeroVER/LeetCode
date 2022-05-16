class Solution {
    public String convert(String s, int numRows) {
        int len = s.length();
        if (len <= numRows || numRows < 2) return s;

        char[] ret = new char[len];
        int i = 0;
        for (int r = 0; r < numRows; r++) {
            int index = r;
            int offset0 = 2 * (numRows - r - 1);
            int offset1 = 2 * r;

            ret[i++] = s.charAt(index);
            for (; index < len;) {
                if (offset0 > 0) {
                    index += offset0;
                    if (index < len) {
                        ret[i++] = s.charAt(index);
                    }
                }
                if (offset1 > 0) {
                    index += offset1;
                    if (index < len) {
                        ret[i++] = s.charAt(index);
                    }
                }
            }
        }
        return String.valueOf(ret);
    }
}