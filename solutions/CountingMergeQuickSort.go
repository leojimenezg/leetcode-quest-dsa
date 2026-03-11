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

// Sort an Array
// Patrones:
//   - Heap Sort
//
// Útil cuando:
//   - se requiere ordenamiento in-place con O(n log n) garantizado
//   - se busca la menor complejidad espacial posible
//   - no se pueden usar funciones built-in de ordenamiento
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(log n)
func sortArray(nums []int) []int {
	n := len(nums)
	buildHeap(nums, n)
	for lastIdx := n - 1; lastIdx > 0; lastIdx-- {
		nums[lastIdx], nums[0] = nums[0], nums[lastIdx]
		heapify(nums, lastIdx, 0)
	}
	return nums
}

// Insertion Sort List
// Patrones:
//   - Insertion Sort
//
// Útil cuando:
//   - se ordena una linked list sin acceso aleatorio a sus elementos
//   - se reutilizan los nodos existentes sin asignar memoria adicional
//   - se mantiene una lista ordenada creciente insertando un nodo a la vez
//
// Complejidad:
//   - Tiempo: O(n^2)
//   - Espacio: O(1)
func insertionSortList(head *ListNode) *ListNode {
	sortHead := &ListNode{}
	for current := head; current != nil; {
		next := current.Next
		for node := sortHead; node != nil; node = node.Next {
			if node.Next == nil {
				current.Next = nil
				node.Next = current
				break
			} else if current.Val < node.Next.Val {
				current.Next = node.Next
				node.Next = current
				break
			}
		}
		current = next
	}
	return sortHead.Next
}
