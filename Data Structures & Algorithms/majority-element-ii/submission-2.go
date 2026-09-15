func majorityElement(nums []int) []int {
	num1, num2 := 0, 0
	cnt1, cnt2 := 0, 0

	for _, num := range nums{
		if num1 == num{
			cnt1++
		}else if num2 == num{
			cnt2++
		}else if cnt1 == 0{
			num1 = num
			cnt1 = 1 
		}else if cnt2 == 0{
			num2 = num
			cnt2 = 1
		}else{
			cnt1--
			cnt2--
		}
	}

	cnt1, cnt2 = 0, 0
	for _, num := range nums{
		if num1 == num{
			cnt1++
		}else if num2 == num{
			cnt2++
		}
	}
	res := []int{}
	if cnt1 > len(nums)/3{
		res = append(res, num1)
	}
	if cnt2 > len(nums)/3{
		res = append(res, num2)
	}
	return res
}
