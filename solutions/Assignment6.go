// ============================================================================
// Assignment 6
// ============================================================================

package solutions

import "sort"

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

// Find Right Interval
// Patrones:
//   - Binary Search
//   - Hash Table
//
// Útil cuando:
//   - se busca el mínimo valor mayor o igual a un target
//   - los valores son únicos y permiten mapeo directo a índices
//   - ordenar previamente habilita búsqueda eficiente
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(n)
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	startToOriginalIdx := make(map[int]int)
	sortedStarts := make([]int, 0, n)
	for originalIdx, interval := range intervals {
		startToOriginalIdx[interval[0]] = originalIdx
		sortedStarts = append(sortedStarts, interval[0])
	}
	sort.Ints(sortedStarts)
	res := make([]int, n)
	for i, interval := range intervals {
		rightIntervalIdx := sort.SearchInts(sortedStarts, interval[1])
		if rightIntervalIdx == n {
			res[i] = -1
		} else {
			res[i] = startToOriginalIdx[sortedStarts[rightIntervalIdx]]
		}
	}
	return res
}
