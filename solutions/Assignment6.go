// ============================================================================
// Assignment 6
// ============================================================================

package solutions

// Merge Sorted Array
// Patrones:
//   - Two Pointers
//   - Reverse Traversal
//
// Útil cuando:
//   - se fusionan dos arrays ordenados in-place
//   - el array destino tiene espacio suficiente al final
//   - recorrer de derecha a izquierda evita sobrescribir elementos no procesados
//
// Complejidad:
//   - Tiempo: O(m+n)
//   - Espacio: O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i, j, k := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0 && k >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	for ; j >= 0 && k >= 0; j-- {
		nums1[k] = nums2[j]
		k--
	}
}
