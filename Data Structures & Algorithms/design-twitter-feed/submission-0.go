// Ordered by the tweet counter at index 0.
type TweetHeap [][]int

func (h TweetHeap) Len() int            { return len(h) }
func (h TweetHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h TweetHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *TweetHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

type Twitter struct {
    count     int
    tweetMap  map[int][][2]int    // userId -> [count, tweetId]
    followMap map[int]map[int]bool // userId -> set of followeeIds
}

func Constructor() Twitter {
    return Twitter{
        count:     0,
        tweetMap:  make(map[int][][2]int),
        followMap: make(map[int]map[int]bool),
    }
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
    if _, exists := t.tweetMap[userId]; !exists {
        t.tweetMap[userId] = make([][2]int, 0, 10)
    }
    t.tweetMap[userId] = append(t.tweetMap[userId], [2]int{t.count, tweetId})
    if len(t.tweetMap[userId]) > 10 {
        t.tweetMap[userId] = t.tweetMap[userId][1:]
    }
    t.count--
}

func (t *Twitter) GetNewsFeed(userId int) []int {
    res := []int{}
    if _, ok := t.followMap[userId]; !ok {
        t.followMap[userId] = make(map[int]bool)
    }
    t.followMap[userId][userId] = true
    minHeap := &TweetHeap{}
    heap.Init(minHeap)

    if len(t.followMap[userId]) >= 10 {
        maxHeap := &TweetHeap{}
        heap.Init(maxHeap)
        for fId := range t.followMap[userId] {
            if tweets, exists := t.tweetMap[fId]; exists && len(tweets) > 0 {
                idx := len(tweets) - 1
                c := tweets[idx][0]
                tId := tweets[idx][1]
                heap.Push(maxHeap, []int{-c, tId, fId, idx - 1})
                if maxHeap.Len() > 10 {
                    heap.Pop(maxHeap)
                }
            }
        }

        for maxHeap.Len() > 0 {
            arr := heap.Pop(maxHeap).([]int)
            negCount := arr[0]
            tId := arr[1]
            fId := arr[2]
            nextIdx := arr[3]
            realCount := -negCount
            heap.Push(minHeap, []int{realCount, tId, fId, nextIdx})
        }
    } else {
        for fId := range t.followMap[userId] {
            if tweets, exists := t.tweetMap[fId]; exists && len(tweets) > 0 {
                idx := len(tweets) - 1
                c := tweets[idx][0]
                tId := tweets[idx][1]
                heap.Push(minHeap, []int{c, tId, fId, idx - 1})
            }
        }
    }

    for minHeap.Len() > 0 && len(res) < 10 {
        arr := heap.Pop(minHeap).([]int)
        tId := arr[1]
        fId := arr[2]
        nextIdx := arr[3]

        res = append(res, tId)
        if nextIdx >= 0 {
            older := t.tweetMap[fId][nextIdx]
            heap.Push(minHeap, []int{older[0], older[1], fId, nextIdx - 1})
        }
    }

    return res
}

func (t *Twitter) Follow(followerId, followeeId int) {
    if _, ok := t.followMap[followerId]; !ok {
        t.followMap[followerId] = make(map[int]bool)
    }
    t.followMap[followerId][followeeId] = true
}

func (t *Twitter) Unfollow(followerId, followeeId int) {
    if _, ok := t.followMap[followerId]; ok {
        delete(t.followMap[followerId], followeeId)
    }
}