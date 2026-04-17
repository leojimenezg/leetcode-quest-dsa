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

// Grumpy Bookstore Owner
//
// Patrones:
//   - Sliding Window (Fixed Size)
//
// Útil cuando:
//   - hay una base fija de satisfacción independiente de la técnica
//   - se busca maximizar el beneficio adicional en una ventana de tamaño fijo
//   - se pueden separar los dos componentes del resultado y calcularlos independientemente
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	left, sumWindow := 0, 0
	satisfied, technique := 0, 0
	for right := range n {
		if grumpy[right] == 1 {
			sumWindow += customers[right]
		} else {
			satisfied += customers[right]
		}
		if right-left+1 > minutes {
			if grumpy[left] == 1 {
				sumWindow -= customers[left]
			}
			left++
		}
		technique = max(technique, sumWindow)
	}
	return satisfied + technique
}
