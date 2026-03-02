// ============================================================================
// Sorting
// ============================================================================

package solutions

import (
	"math"
	"sort"
)

// Minimum Absolute Difference
// Patrones:
//   - Sorting
//   - Linear Scan
//
// Útil cuando:
//   - se busca la diferencia mínima entre elementos de un arreglo
//   - se necesita comparar proximidad entre elementos
//   - en un arreglo ordenado, la respuesta siempre está entre vecinos contiguos
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(log n)
func minimumAbsDifference(arr []int) [][]int {
	n := len(arr)
	pairs := make([][]int, 0, n)
	sort.Ints(arr)
	mad := math.MaxInt
	for i := 0; i < n-1; i++ {
		diff := abs(arr[i] - arr[i+1])
		mad = min(mad, diff)
	}
	for i := 0; i < n-1; i++ {
		diff := abs(arr[i] - arr[i+1])
		if diff == mad {
			pairs = append(pairs, []int{arr[i], arr[i+1]})
		}
	}
	return pairs
}
