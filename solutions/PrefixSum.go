// ============================================================================
// Prefix Sum
// ============================================================================

package solutions

// Find the Highest Altitude
//
// Patrones:
//   - Prefix Sum (optimized)
//   - Running Maximum
//   - Single Pass Traversal
//
// Útil cuando:
//   - se calculan sumas acumuladas progresivamente
//   - solo importa el valor máximo de las sumas
//   - se puede procesar en una sola pasada sin almacenar histórico
//   - se busca optimizar espacio evitando arrays auxiliares
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func largestAltitude(gain []int) int {
	prev := 0
	res := prev
	for i := range len(gain) {
		prev = prev + gain[i]
		if prev > res {
			res = prev
		}
	}
	return res
}

// Make Sum Divisible by P
//
// Patrones:
//   - Prefix Sum
//   - Hash Table
//   - Mathematical Reduction
//
// Útil cuando:
//   - se busca un subarray cuya suma cumple una condición de módulo
//   - se necesita encontrar el subarray más corto que satisface la condición
//   - se puede reformular como búsqueda de prefix sums previos
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func minSubarray(nums []int, p int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	remainder := total % p
	if remainder == 0 {
		return 0
	}
	mods := make(map[int]int)
	mods[0] = -1
	sum := 0
	res := len(nums)
	for i, v := range nums {
		sum += v
		currentMod := sum % p
		targetMod := (currentMod - remainder + p) % p
		if idx, ok := mods[targetMod]; ok {
			length := i - idx
			if length < res {
				res = length
			}
		}
		mods[currentMod] = i
	}
	if res >= len(nums) {
		return -1
	}
	return res
}

// Ways to Make a Fair Array
//
// Patrones:
//   - Prefix Sum
//   - Running Totals
//   - Position Swap Analysis
//
// Útil cuando:
//   - remover un elemento afecta las posiciones de elementos posteriores
//   - se necesita evaluar todas las posiciones sin simulación O(n²)
//   - los elementos se dividen en dos categorías (par/impar, positivo/negativo)
//   - se puede precalcular totales y ajustarlos dinámicamente
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func waysToMakeFair(nums []int) int {
	totalEven, totalOdd := 0, 0
	for i, v := range nums {
		if i%2 == 0 {
			totalEven += v
		} else {
			totalOdd += v
		}
	}
	res := 0
	leftEven, leftOdd := 0, 0
	for i, v := range nums {
		if i%2 == 0 {
			totalEven -= v
		} else {
			totalOdd -= v
		}
		newEven := leftEven + totalOdd
		newOdd := leftOdd + totalEven
		if newEven == newOdd {
			res++
		}
		if i%2 == 0 {
			leftEven += v
		} else {
			leftOdd += v
		}
	}
	return res
}
