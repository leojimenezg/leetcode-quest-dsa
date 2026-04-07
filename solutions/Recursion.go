// ============================================================================
// Recursion
// ============================================================================

package solutions

// Merge Two Sorted Lists
//
// Patrones:
//   - Recursion
//   - Pointer Rewiring
//
// Útil cuando:
//   - se fusionan dos listas ordenadas nodo por nodo
//   - se puede delegar el resto del problema a una llamada recursiva
//   - el caso base es cuando una lista se agota
//
// Complejidad:
//   - Tiempo: O(m+n)
//   - Espacio: O(m+n)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
