const max = 214748364;
const min = -214748364;
/**
 * @param {number} x
 * @return {number}
 */
 var reverse = function(x) {
    let ret = 0;

    while (x != 0) {
        let digit = x % 10;

        if (ret > max || (ret == max && digit > 7)) {
            return 0;
        }
        if (ret < min || (ret == min && digit < -8)) {
            return 0;
        }

        ret = ret * 10 + digit;
        x = parseInt(x / 10);
    }
    return ret;
};