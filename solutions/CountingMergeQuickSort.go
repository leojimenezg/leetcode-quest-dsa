package solutions

import "container/heap"

// Find Kth Largest Element in an Array
// Patrones:
//   - Min Heap
//
// Útil cuando:
//   - se busca el k-ésimo elemento más grande sin ordenar
//   - se mantiene un conjunto de tamaño fijo con los k mayores vistos
//   - el top del heap siempre es el candidato a eliminar
//
// Complejidad:
//   - Tiempo: O(n log k)
//   - Espacio: O(k)
func findKthLargest(nums []int, k int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, v := range nums {
		heap.Push(minHeap, v)
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}
	return heap.Pop(minHeap).(int)
}
