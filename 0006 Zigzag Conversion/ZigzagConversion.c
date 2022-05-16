char * convert(char * s, int numRows){
    int len = strlen(s);
    if (len <= numRows || numRows < 2) return s;
    
    char *ret = (char *) calloc(len + 1, sizeof(char));
    
    int i = 0;
    for (int r = 0; r < numRows; r++) {
        int index = r;
        int offset0 = 2 * (numRows - r - 1);
        int offset1 = 2 * r;

        ret[i++] = s[index];
        for (; index < len;) {
            if (offset0 > 0) {
                index += offset0;
                if (index < len) {
                    ret[i++] = s[index];
                }
            }
            if (offset1 > 0) {
                index += offset1;
                if (index < len) {
                    ret[i++] = s[index];
                }
            }
        }
    }
    return ret;
}