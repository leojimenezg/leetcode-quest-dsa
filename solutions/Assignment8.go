// ============================================================================
// Assignment 8
// ============================================================================

package solutions

// Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold
//
// Patrones:
//   - Sliding Window (Fixed Size)
//
// Útil cuando:
//   - se evalúa una condición sobre subarrays de tamaño fijo
//   - se puede mantener la suma de la ventana incrementalmente
//   - se evita recalcular la suma desde cero en cada paso
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func numOfSubarrays(arr []int, k int, threshold int) int {
	start, sum, total := 0, 0, 0
	for i := range arr {
		sum += arr[i]
		if i >= k-1 {
			if sum/k >= threshold {
				total++
			}
			sum -= arr[start]
			start++
		}
	}
	return total
}
