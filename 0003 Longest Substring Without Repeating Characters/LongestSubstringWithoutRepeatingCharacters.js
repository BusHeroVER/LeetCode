/**
 * @param {string} s
 * @return {number}
 */
 var lengthOfLongestSubstring = function(s) {
    var s_len = s.length;
    if (s_len <= 1)
        return s_len;

    var head = 0, ret = 0, map = {};
    
    for (var index = 0; index < s_len;) {
        var c = s.charAt(index);
        if (map[c] != undefined) {
            ret = Math.max(index - head, ret);
            head = Math.max(head, map[c]);
        }
        map[c] = ++index;
    }
    return Math.max(s_len - head, ret);    
};