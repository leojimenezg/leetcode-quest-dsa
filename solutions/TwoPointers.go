// ============================================================================
// Two Pointers
// ============================================================================

package solutions

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
