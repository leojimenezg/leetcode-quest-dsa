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

// Magical String
// Patrones:
//   - Two Pointers
//   - Self-referential Generation
//
// Útil cuando:
//   - una secuencia se construye usando sus propios valores como instrucciones
//   - un índice lee la secuencia mientras otro la genera
//   - los valores se alternan entre dos opciones fijas
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func magicalString(n int) int {
	s := make([]int, 0, 2*n)
	s = append(s, 1, 2, 2)
	onesCount := 1
	idx := 2
	group := 1
	for idx < n {
		for range s[idx] {
			if len(s) < n && group == 1 {
				onesCount++
			}
			s = append(s, group)
		}
		group = 3 - group
		idx++
	}
	return onesCount
}
