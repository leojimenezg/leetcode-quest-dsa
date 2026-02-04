// ============================================================================
// Array 2
// ============================================================================

package solutions

import "sort"

// Set Mismatch
// Patrones:
//   - Hash Table
//   - Equation-based Derivation
//
// Útil cuando:
//   - se detectan duplicados y valores faltantes
//   - se conoce el rango esperado de valores
//   - se obtiene el resultado mediante una ecuación/fórmula
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func FindErrorNums(nums []int) []int {
	n := len(nums)
	expectedSum := n * (n + 1) / 2
	actualSum := 0
	repeatedNum := -1
	uniques := make(map[int]bool)
	for _, v := range nums {
		if _, ok := uniques[v]; ok {
			repeatedNum = v
		}
		uniques[v] = true
		actualSum += v
	}
	lostNum := expectedSum - (actualSum - repeatedNum)
	return []int{repeatedNum, lostNum}
}

// How Many Numbers Are Smaller Than the Current Number
// Patrones:
//   - Sorting
//   - Value Compression
//
// Útil cuando:
//   - se necesita ranking relativo
//   - O(n log n) es aceptable
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(n)
func SmallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)
	counts := make(map[int]int)
	for i, v := range sorted {
		if _, ok := counts[v]; !ok {
			counts[v] = i
		}
	}
	for i, v := range nums {
		sorted[i] = counts[v]
	}
	return sorted
}

// Find All Numbers Disappeared in an Array
// Patrones:
//   - Index as Hash
//
// Útil cuando:
//   - el rango de valores es conocido (1...n)
//   - se puede usar el índice como estructura auxiliar
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	uniques := make([]bool, n)
	for _, v := range nums {
		uniques[v-1] = true
	}
	missing := make([]int, 0, n)
	for i, v := range uniques {
		if !v {
			missing = append(missing, i+1)
		}
	}
	return missing
}
