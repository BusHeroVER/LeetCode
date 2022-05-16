/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
 var convert = function(s, numRows) {
    var len = s.length;
    if (len <= numRows || numRows < 2) return s;

    var ret = "";
    for (var r = 0; r < numRows; r++) {
        var index = r;
        var offset0 = 2 * (numRows - r - 1);
        var offset1 = 2 * r;

        ret += s.charAt(index);
        for (; index < len;) {
            if (offset0 > 0) {
                index += offset0;
                if (index < len) {
                    ret += s.charAt(index);
                }
            }
            if (offset1 > 0) {
                index += offset1;
                if (index < len) {
                    ret += s.charAt(index);
                }
            }
        }
    }
    return ret;
};