// ============================================================================
// Linked List
// ============================================================================

package solutions

// Remove Duplicates from Sorted List
// Patrones:
//   - Single-Pass Traversal
//   - In-place Transformation
//
// Útil cuando:
//   - los duplicados son contiguos (ordenada)
//   - se permite modificar la estructura original
//   - se requiere O(1) espacio extra
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	var prev *ListNode
	for current := head; current != nil; current = current.Next {
		if prev != nil {
			if prev.Val == current.Val {
				prev.Next = current.Next
				continue
			}
		}
		prev = current
	}
	return head
}

// Odd Even Linked List
//
// Patrones:
//   - Single-Pass Traversal
//   - In-place Transformation
//   - Pointer Rewiring
//
// Útil cuando:
//   - se requiere reagrupar nodos según su posición (odd / even)
//   - la estructura solo permite avanzar en una dirección (listas enlazadas)
//   - no se puede regresar, reiniciar o usar accesos aleatorios
//   - no se permite usar memoria extra proporcional a n
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddHead := head
	evenHead := head.Next
	evenCurrent := evenHead
	for evenCurrent != nil && evenCurrent.Next != nil {
		oddHead.Next = evenCurrent.Next
		oddHead = oddHead.Next
		evenCurrent.Next = oddHead.Next
		evenCurrent = evenCurrent.Next
	}
	oddHead.Next = evenHead
	return head
}

// Reverse Linked List
//
// Patrones:
//   - Single-Pass Traversal
//   - Pointer Reversal (Rewiring)
//   - In-place Transformation
//
// Útil cuando:
//   - se requiere invertir una lista enlazada
//   - la estructura solo permite avanzar en una dirección
//   - no se permite usar memoria extra proporcional a n
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func reverseList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	for current != nil {
		next := current.Next
		current.Next = prev
		if next == nil {
			head = current
		}
		prev = current
		current = next
	}
	return head
}
