// ============================================================================
// Queue
// ============================================================================

package solutions

// Number of Students Unable to Eat Lunch
// Patrones:
//   - Counting
//   - Greedy Consumption
//
// Útil cuando:
//   - el orden puede rotarse indefinidamente
//   - solo importa la disponibilidad total
//   - el consumo ocurre de forma secuencial y estricta
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func CountStudents(students []int, sandwiches []int) int {
	wanted := [2]int{0, 0}
	for _, s := range students {
		wanted[s]++
	}
	for _, s := range sandwiches {
		if wanted[s] > 0 {
			wanted[s]--
		} else {
			break
		}
	}
	return wanted[0] + wanted[1]
}

// Time Needed to Buy Tickets
// Patrones:
//   - Mathematical Simulation
//
// Útil cuando:
//   - un proceso repetitivo puede expresarse como suma directa
//   - hay un punto de terminación específico
//   - los elementos antes y después del punto se comportan distinto
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func TimeRequiredToBuy(tickets []int, k int) int {
	res := 0
	for i := range tickets {
		if i < k {
			res += min(tickets[k], tickets[i])
		} else if i == k {
			res += tickets[k]
		} else {
			res += min(tickets[k]-1, tickets[i])
		}
	}
	return res
}
