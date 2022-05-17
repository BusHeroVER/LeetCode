const int max = INT_MAX / 10;
const int min = INT_MIN / 10;

int reverse(int x){
    int ret = 0;
    
    while (x != 0) {
        int digit = x % 10;

        if (ret > max || (ret == max && digit > 7)) {
			return 0;
		}
		if (ret < min || (ret == min && digit < -8)) {
			return 0;
		}

		ret = ret * 10 + digit;
		x /= 10;
    }
    return ret;
}