type Solution struct {
 prefix []int	
}

func Constructor(w []int) Solution {
	prefix := []int{0}
    for _, wgt := range w {
        prefix = append(prefix, prefix[len(prefix)-1]+wgt)
    }
    return Solution{prefix: prefix}	
}

func (this *Solution) PickIndex() int {
	target := float64(this.prefix[len(this.prefix)-1]) * rand.Float64()
    l, r := 1, len(this.prefix)
    for l < r {
        mid := (l + r) >> 1
        if float64(this.prefix[mid]) <= target {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return l - 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w)
 * param_1 := obj.PickIndex()
 */
 