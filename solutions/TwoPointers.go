// ============================================================================
// Two Pointers
// ============================================================================

package solutions

import "sort"

// Linked List Cycle
// Patrones:
//   - Two Pointers
//
// Útil cuando:
//   - se detecta un ciclo en una lista enlazada
//   - se requiere O(1) espacio sin modificar la estructura
//   - dos pointers a distinta velocidad inevitablemente convergen en un ciclo
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 3Sum Closest
// Patrones:
//   - Sorting
//   - Two Pointers
//
// Útil cuando:
//   - se busca una tripleta cuya suma se acerca a un target
//   - ordenar permite descartar candidatos moviendo pointers
//   - se fija un elemento y se reduce el problema a un par
//
// Complejidad:
//   - Tiempo: O(n^2)
//   - Espacio: O(log n)
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	sum := nums[0] + nums[1] + nums[2]
	for i := range nums {
		left := i + 1
		right := n - 1
		for left < right {
			threeSum := nums[i] + nums[left] + nums[right]
			if threeSum < target {
				left++
			} else if threeSum > target {
				right--
			} else {
				return threeSum
			}
			diff := abs(threeSum - target)
			if diff < abs(sum-target) {
				sum = threeSum
			}
		}
	}
	return sum
}
