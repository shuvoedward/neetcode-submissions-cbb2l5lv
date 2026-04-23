type MinStack struct {
	stack []int
	minS []int
}

func Constructor() MinStack {
    return MinStack{stack : []int{}, minS:[]int{}}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0{
		this.stack = append(this.stack, val)
		this.minS = append(this.minS, val)
		return 
	}

	this.stack = append(this.stack, val)
	lastMin := this.minS[len(this.minS)-1]
	if val < lastMin{
		this.minS = append(this.minS, val)
	}else{
		this.minS = append(this.minS, lastMin)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minS = this.minS[:len(this.minS)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minS[len(this.minS)-1]
}
