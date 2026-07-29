func rob(nums []int) int {
    n := len(nums)
    memo := make([]int, n+1)
    for i := 0; i <= n; i++ {
        memo[i] = -1
    }

    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= len(nums) {
            return 0
        }
        if memo[i] != -1 {
            return memo[i]
        }
        memo[i] = max(dfs(i+1), nums[i] + dfs(i+2))
        return memo[i]
    }

    return dfs(0)
}

// do i take the index itself or don't take the index?
// that is the core logic. 
