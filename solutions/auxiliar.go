package solutions

// Implement Queue using Stacks
// Patrones:
//   - Two-Stack Queue (Lazy Transfer)
//   - Amortized Analysis
//
// Útil cuando:
//   - solo se permiten stacks
//   - se puede amortizar el costo de inversión
//
// Complejidad:
//   - Tiempo: O(1) amortizado
//   - Espacio: O(n)
type MyQueue struct {
	inStack  []int
	outStack []int
	size     int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
	this.size++
}

func (this *MyQueue) Pop() int {
	if this.size == 0 {
		return 0
	}
	if len(this.outStack) == 0 {
		this.lazyTransfer()
	}
	r := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	this.size--
	return r
}

func (this *MyQueue) Peek() int {
	if this.size == 0 {
		return 0
	}
	if len(this.outStack) == 0 {
		this.lazyTransfer()
	}
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	return this.size == 0
}

// lazyTransfer transfiere todos los elementos de un stack a otro.
// Esto invierte el orden: el elemento botton se convierte en el top
func (this *MyQueue) lazyTransfer() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}
