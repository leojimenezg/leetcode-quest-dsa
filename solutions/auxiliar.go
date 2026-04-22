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

type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type Pair struct {
	I   int
	J   int
	Sum int
}

type PairKey struct {
	I int
	J int
}

type MinHeapPairs []Pair

func (h MinHeapPairs) Len() int { return len(h) }

func (h MinHeapPairs) Less(i, j int) bool { return h[i].Sum < h[j].Sum }

func (h MinHeapPairs) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeapPairs) Push(x any) { *h = append(*h, x.(Pair)) }

func (h *MinHeapPairs) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func heapify(s []int, n, i int) {
	largestIdx := i
	leftIdx := 2*i + 1
	rightIdx := 2*i + 2
	if leftIdx < n && s[leftIdx] > s[largestIdx] {
		largestIdx = leftIdx
	}
	if rightIdx < n && s[rightIdx] > s[largestIdx] {
		largestIdx = rightIdx
	}
	if largestIdx != i {
		s[i], s[largestIdx] = s[largestIdx], s[i]
		heapify(s, n, largestIdx)
	}
}

func buildHeap(s []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(s, n, i)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	inStack  [][]int
	outStack [][]int
	size     int
}

func Constructor() Queue {
	return Queue{}
}

func (this *Queue) Push(x []int) {
	this.inStack = append(this.inStack, x)
	this.size++
}

func (this *Queue) Pop() []int {
	if this.size == 0 {
		return []int{}
	}
	if len(this.outStack) == 0 {
		this.lazyTransfer()
	}
	r := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	this.size--
	return r
}

func (this *Queue) Peek() []int {
	if this.size == 0 {
		return []int{}
	}
	if len(this.outStack) == 0 {
		this.lazyTransfer()
	}
	return this.outStack[len(this.outStack)-1]
}

func (this *Queue) Empty() bool {
	return this.size == 0
}

func (this *Queue) lazyTransfer() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}
