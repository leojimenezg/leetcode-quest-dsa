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

// Reduction Operations to Make the Array Elements Equal
// Patrones:
//   - Sorting
//   - Linear Scan
//
// Útil cuando:
//   - se cuenta el costo acumulado de reducir elementos a un mínimo común
//   - cada nivel distinto incrementa el costo de todos los elementos superiores
//   - el orden entre elementos es clave para detectar cambios de nivel
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(log n)
func reductionOperations(nums []int) int {
	sort.Ints(nums)
	steps := 0
	ops := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			steps++
			ops += steps
		} else {
			ops += steps
		}
	}
	return ops
}

// Merge Intervals
// Patrones:
//   - Sorting
//   - Linear Scan
//
// Útil cuando:
//   - se fusionan intervalos solapados
//   - ordenar por inicio garantiza que solapamientos son siempre contiguos
//   - se necesita mantener el último resultado para comparaciones sucesivas
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(log n)
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	res := make([][]int, 0, n)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, current := range intervals {
		m := len(res)
		if m > 0 {
			prev := res[m-1]
			if prev[1] >= current[0] {
				res[m-1][1] = max(prev[1], current[1])
				continue
			}
		}
		res = append(res, current)
	}
	return res
}
