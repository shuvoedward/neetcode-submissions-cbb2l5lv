func shipWithinDays(weights []int, days int) int {
	// find max, total
	l, r := 0, 0
	for _, w := range weights{
		if w > l{
			l = w
		}
		r += w
	}

	canShip := func(cap int)bool{
		ship, currCap := 1, cap
		for _, w := range weights{
			if currCap - w < 0{
				ship++
				if ship > days{
					return false
				}

				currCap = cap
				
			}
			currCap -= w
		}

		return true 
	}

	res := r

	for l <= r{
		cap := l + (r-l)/2
		if canShip(cap){
			if cap < res{
				res = cap
			}
			r = cap - 1
		}else{
			l = cap + 1
		}
	}

	return res

}
