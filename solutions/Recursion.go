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

// Find Kth Bit in Nth Binary String
//
// Patrones:
//   - Divide and Conquer
//   - Recursion
//
// Útil cuando:
//   - el string tiene estructura simétrica recursiva
//   - se puede determinar en qué mitad cae k sin construir el string
//   - la mitad derecha es inverse(reverse) de la izquierda
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	size := (1 << n) - 1
	mid := (size + 2 - 1) / 2
	if k == mid {
		return '1'
	} else if k < mid {
		return findKthBit(n-1, k)
	} else {
		bit := findKthBit(n-1, size-k+1)
		if bit == '0' {
			return '1'
		}
		return '0'
	}
}
