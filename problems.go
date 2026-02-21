// ============================================================================
// LeetCode Quest: Data Structures and Algorithms
// ============================================================================

package problems

// ============================================================================
// Assignment 5
// ============================================================================

// Next Greater Node In Linked List
//
// Patrones:
//   - Nested Iteration
//   - Linked List Traversal
//   - Array Construction
//
// Útil cuando:
//   - se necesita comparar cada elemento con los siguientes
//   - la estructura es secuencial (linked list, array)
//   - la solución directa es suficientemente eficiente
//
// Nota:
//   - Existe solución O(n) con Monotonic Stack
//
// Complejidad:
//   - Tiempo: O(n^2)
//   - Espacio: O(n)
func NextLargerNodes(head *ListNode) []int {
	n := 0
	for current := head; current != nil; current = current.Next {
		n++
	}
	nodeVals := make([]int, n)
	idx := 0
	for current := head; current != nil; current = current.Next {
		for next := current.Next; next != nil; next = next.Next {
			if next.Val > current.Val {
				nodeVals[idx] = next.Val
				break
			}
		}
		idx++
	}
	return nodeVals
}

// Continuous Subarray Sum
//
// Patrones:
//   - Prefix Sum
//   - Hash Table
//
// Útil cuando:
//   - se busca subarray cuya suma es divisible por k
//   - se requiere longitud mínima
//   - dos prefix sums con mismo mod implican subarray divisible
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(min(n,k))
func CheckSubarraySum(nums []int, k int) bool {
	mods := make(map[int]int)
	mods[0] = -1
	prefixSum := 0
	for i, v := range nums {
		prefixSum += v
		mod := prefixSum % k
		if idx, ok := mods[mod]; ok {
			if i-idx >= 2 {
				return true
			}
		} else {
			mods[mod] = i
		}
	}
	return false
}
