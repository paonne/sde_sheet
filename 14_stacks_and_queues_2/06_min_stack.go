type MinStack struct {
	stack      []int
	mini_stack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if n := len(this.mini_stack); n == 0 {
		this.mini_stack = append(this.mini_stack, val)
	} else {
		this.mini_stack = append(this.mini_stack, min(val, this.mini_stack[n-1]))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.mini_stack = this.mini_stack[:len(this.mini_stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mini_stack[len(this.mini_stack)-1]
}
