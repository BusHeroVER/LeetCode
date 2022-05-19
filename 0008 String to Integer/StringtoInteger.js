var max = 2147483647;
var min = -2147483648;
/**
 * @param {string} s
 * @return {number}
 */
var myAtoi = function(s) {
    var ret = 0;
    var isReading = false;
	var isPostive = true;

    for (var i = 0; i < s.length; i++){
        var c = s.charAt(i);

        if (c == ' ') {
			if (isReading) {
				return ret;
			}
		} else if (c == '+') {
			if (isReading) {
				return ret;
			} else {
				isReading = true;
			}
		} else if (c == '-') {
            if (isReading) {
				return ret;
			} else {
				isReading = true;
				isPostive = false;
			}
		} else if ('0' <= c && c <= '9') {
			isReading = true;

			ret *= 10;
			if (isPostive) {
				ret += c - '0';
			} else {
				ret -= c - '0';
			}

			if (ret < min) {
				return min;
			} else if (max < ret) {
				return max;
			}
        } else {
            return ret;
        }
    }
    return ret;
};