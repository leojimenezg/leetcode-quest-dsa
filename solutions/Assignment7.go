// ============================================================================
// Assignment 7
// ============================================================================

package solutions

import (
	"slices"
	"sort"
)

// Minimum Number of Days to Make m Bouquets
// Patrones:
//   - Binary Search
//   - Linear Scan
//
// Útil cuando:
//   - se busca el mínimo valor en un rango que satisface una condición monótona
//   - la condición divide el rango en [F, F, ..., T, T, ...] sin excepciones
//   - verificar si un valor candidato es válido es más fácil que encontrarlo directamente
//
// Complejidad:
//   - Tiempo: O(n log d) donde d = max(bloomDay)
//   - Espacio: O(1)
func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if n < m*k {
		return -1
	}
	left := 0
	right := slices.Max(bloomDay)
	for left < right {
		day := (left + right) / 2
		adjacent := 0
		bouquets := 0
		for _, v := range bloomDay {
			if v <= day {
				adjacent++
				if adjacent == k {
					bouquets++
					adjacent = 0
				}
			} else {
				adjacent = 0
			}
		}
		if bouquets < m {
			left = day + 1
		} else {
			right = day
		}
	}
	return left
}

// Merge k Sorted Lists
// Patrones:
//   - Sorting
//
// Útil cuando:
//   - se unifican múltiples estructuras en una sola
//   - el orden relativo entre elementos es lo único relevante
//   - el origen de los elementos no importa tras la unificación
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(n)
func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}
	n := 0
	for _, head := range lists {
		for current := head; current != nil; current = current.Next {
			n++
		}
	}
	if n == 0 {
		return nil
	}
	unifiedNodes := make([]*ListNode, 0, n)
	for _, head := range lists {
		for current := head; current != nil; current = current.Next {
			unifiedNodes = append(unifiedNodes, current)
		}
	}
	sort.Slice(unifiedNodes, func(i, j int) bool {
		return unifiedNodes[i].Val < unifiedNodes[j].Val
	})
	for i := 0; i < n-1; i++ {
		unifiedNodes[i].Next = unifiedNodes[i+1]
	}
	unifiedNodes[n-1].Next = nil
	return unifiedNodes[0]
}
