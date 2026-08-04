type NumArray struct {
   nums []int 
}


func Constructor(nums []int) NumArray {
	res := make([]int, len(nums))
	prefix := 0
	for i, num := range nums{
		prefix += num
		res[i] = prefix
	}
	return NumArray{nums: res}
}


func (this *NumArray) SumRange(left int, right int) int {
	if left == 0{
		return this.nums[right]
	}
    return this.nums[right] - this.nums[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */