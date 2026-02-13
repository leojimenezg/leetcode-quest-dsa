// ============================================================================
// Heap
// ============================================================================

package solutions

import "container/heap"

// Last Stone Weight
// Patrones:
//   - Priority Queue (Max Heap)
//
// Útil cuando:
//   - se necesita acceder repetidamente al mayor elemento
//   - el conjunto cambia dinámicamente
//   - mantener orden total no es necesario
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(n)
func LastStoneWeight(stones []int) int {
	maxHeap := MaxHeap(stones)
	heap.Init(&maxHeap)
	for maxHeap.Len() > 1 {
		y := heap.Pop(&maxHeap).(int)
		x := heap.Pop(&maxHeap).(int)
		if x != y {
			heap.Push(&maxHeap, y-x)
		}
	}
	if maxHeap.Len() == 0 {
		return 0
	}
	return heap.Pop(&maxHeap).(int)
}

// Find K Pairs with Smallest Sums
// Patrones:
//   - Priority Queue (Min Heap)
//
// Útil cuando:
//   - se explora una matriz virtual ordenada
//   - solo se necesitan los k mejores resultados
//
// Complejidad:
//   - Tiempo: O(k log k)
//   - Espacio: O(k)
func KSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pairs := make([][]int, 0, k)
	visited := make(map[PairKey]bool)
	minHeap := &MinHeapPairs{}
	heap.Init(minHeap)
	heap.Push(minHeap, Pair{I: 0, J: 0, Sum: nums1[0] + nums2[0]})
	visited[PairKey{0, 0}] = true
	for len(pairs) < k {
		pair := heap.Pop(minHeap).(Pair)
		pairs = append(pairs, []int{nums1[pair.I], nums2[pair.J]})
		next := PairKey{I: pair.I, J: pair.J + 1}
		if next.J < len(nums2) && !visited[next] {
			heap.Push(minHeap, Pair{I: next.I, J: next.J, Sum: nums1[next.I] + nums2[next.J]})
			visited[next] = true
		}
		next = PairKey{I: pair.I + 1, J: pair.J}
		if next.I < len(nums1) && !visited[next] {
			heap.Push(minHeap, Pair{I: next.I, J: next.J, Sum: nums1[next.I] + nums2[next.J]})
			visited[next] = true
		}
	}
	return pairs
}

// Construct Target Array With Multiple Sums
// Patrones:
//   - Reverse Greedy
//   - Priority Queue (Max Heap)
//   - Mathematical Reduction (Module)
//
// Útil cuando:
//   - un solo elemento domina el estado final
//   - la simulación directa no escala
//   - el proceso puede revertirse de forma determinista
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(n)
func IsPossible(target []int) bool {
	maxHeap := MaxHeap(target)
	heap.Init(&maxHeap)
	currentSum := 0
	for _, v := range target {
		currentSum += v
	}
	for {
		currentMax := heap.Pop(&maxHeap).(int)
		if currentMax == 1 {
			return true
		}
		restSum := currentSum - currentMax
		if restSum <= 0 || restSum >= currentMax {
			return false
		}
		prevNum := currentMax % restSum
		if prevNum == 0 && restSum != 1 {
			return false
		}
		heap.Push(&maxHeap, prevNum)
		currentSum = prevNum + restSum
	}
}
