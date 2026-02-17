// ============================================================================
// Assignment 3
// ============================================================================

package solutions

import "container/heap"

type Apples struct {
	Total  int
	ExpDay int
}

type MinHeapApples []Apples

func (h MinHeapApples) Len() int { return len(h) }

func (h MinHeapApples) Less(i, j int) bool { return h[i].ExpDay < h[j].ExpDay }

func (h MinHeapApples) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeapApples) Push(x any) { *h = append(*h, x.(Apples)) }

func (h *MinHeapApples) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// Maximum Number of Eaten Apples
// Patrones:
//   - Priority Queue (MinHeap)
//   - Greedy
//
// Útil cuando:
//   - siempre conviene consumir primero el recurso con menor deadline
//   - el input es finito pero el consumo puede extenderse más allá
//   - es necesario mantener dinámicamente el “más urgente”
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(n)
func EatenApples(apples []int, days []int) int {
	n := len(apples)
	count := 0
	minHeap := &MinHeapApples{}
	heap.Init(minHeap)
	idx := 0
	for idx < n || minHeap.Len() > 0 {
		if idx < n && apples[idx] > 0 {
			heap.Push(minHeap, Apples{Total: apples[idx], ExpDay: days[idx] + idx})
		}
		for minHeap.Len() > 0 {
			top := (*minHeap)[0]
			if top.ExpDay > idx {
				break
			}
			heap.Pop(minHeap)
		}
		if minHeap.Len() > 0 {
			latest := heap.Pop(minHeap).(Apples)
			latest.Total--
			if latest.Total > 0 {
				heap.Push(minHeap, latest)
			}
			count++
		}
		idx++
	}
	return count
}

// Design Circular Queue
// Patrones:
//   - Circular Queue
//   - Modular Indexing
//   - Dual Index Tracking
//
// Útil cuando:
//   - se necesita una queue FIFO con capacidad fija
//   - el espacio debe reutilizarse de forma controlada
//   - es necesario mantener índices dentro de un rango circular fijo
//   - se necesita mantener un índice al inicio y otro al final
//
// Complejidad:
//   - Tiempo: O(1) por operación
//   - Espacio: O(k), donde k es la capacidad fija de la queue
type MyCircularQueue struct {
	Items    []int
	FrontIdx int
	RearIdx  int
	Cap      int
	Size     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{Items: make([]int, k), Cap: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.Size == this.Cap {
		return false
	}
	this.Items[this.RearIdx] = value
	this.RearIdx = (this.RearIdx + 1) % this.Cap
	this.Size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.Size == 0 {
		return false
	}
	this.FrontIdx = (this.FrontIdx + 1) % this.Cap
	this.Size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.Size == 0 {
		return -1
	}
	return this.Items[this.FrontIdx]
}

func (this *MyCircularQueue) Rear() int {
	if this.Size == 0 {
		return -1
	}
	idx := (this.RearIdx - 1 + this.Cap) % this.Cap
	return this.Items[idx]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Size == this.Cap
}
