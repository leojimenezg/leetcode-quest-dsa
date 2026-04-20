// ============================================================================
// Assignment 9
// ============================================================================

package solutions

// Reorder List
//
// Patrones:
//   - Array as Index (LinkedList to Array)
//   - Two Pointers
//
// Útil cuando:
//   - se necesita acceso aleatorio a nodos de una linked list
//   - se intercalan nodos del inicio y del final
//   - convertir a array simplifica el rewiring de punteros
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	nodes := make([]*ListNode, n)
	cur := head
	for i := range n {
		nodes[i] = cur
		cur = cur.Next
	}
	i, j := 0, n-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i >= j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
