/**
 * @param {number} x
 * @return {boolean}
 */
 var isPalindrome = function(x) {
    if (x < 0) return false;

	var org = x;
	var reverse = 0;
	while (x != 0) {
		reverse = reverse*10 + x%10;
		x = parseInt(x / 10);
	}
	return org == reverse;
};