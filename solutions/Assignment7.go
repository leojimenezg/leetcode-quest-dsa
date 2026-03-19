// ============================================================================
// Assignment 7
// ============================================================================

package solutions

import "slices"

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
