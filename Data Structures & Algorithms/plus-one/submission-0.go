func plusOne(digits []int) []int {
	lastIndex := len(digits)-1
    for i := lastIndex; i >= 0; i--{
		if digits[i] == 9{
			digits[i] = 0
		}else{
			digits[i] = digits[i] + 1
			break
		}		
	}

	if digits[0] == 0{
		res := make([]int, len(digits)+1)
		res[0] = 1
		return res
	}
	
	return digits
}
