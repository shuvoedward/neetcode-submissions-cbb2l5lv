func rearrangeArray(nums []int) []int {
i, j := 0, 1
    res := make([]int, len(nums))
    for k := 0; k < len(nums); k++ {
        if nums[k] > 0 {
            res[i] = nums[k]
            i += 2
        } else {
            res[j] = nums[k]
            j += 2
        }
    }
    return res
}
