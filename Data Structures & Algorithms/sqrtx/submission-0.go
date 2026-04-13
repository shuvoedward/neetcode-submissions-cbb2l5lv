func mySqrt(x int) int {
	res := 0
	l, r := 1, x

	for l <= r{
		m := l + (r - l)/2
		if int64(m)* int64(m) > int64(x){
			r = m - 1
		}else if int64(m)* int64(m) < int64(x){
			l = m + 1
			res = m
		}else{
			return m
		}
	}

	return res
}
