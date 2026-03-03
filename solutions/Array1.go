// ============================================================================
// Array 1
// ============================================================================

package solutions

// Concatenation of Array
// Patrones:
//   - Array Construction
//
// Útil cuando:
//   - se necesita crear un nuevo arreglo a partir de uno existente
//   - el tamaño final es conocido de antemano
//   - no se requiere lógica condicional compleja
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func getConcatenation(nums []int) []int {
	n := len(nums)
	res := make([]int, 2*n)
	for i, v := range nums {
		res[i] = v
		res[n+i] = v
	}
	return res
}

// Shuffle the Array
// Patrones:
//   - Array Construction
//   - Interleaving
//
// Útil cuando:
//   - se necesita crear un nuevo arreglo a partir de uno existente
//   - se combinan dos secuencias con posiciones conocidas
//   - se requiere acceso indexado directo
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i := range n {
		res[2*i] = nums[i]
		res[2*i+1] = nums[n+i]
	}
	return res
}

// Max Consecutive Ones
// Patrones:
//   - Single Pass
//   - Running Counter
//
// Útil cuando:
//   - se busca la subsecuencia consecutiva más larga
//   - una condición rompe el conteo y obliga a reiniciar
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func findMaxConsecutiveOnes(nums []int) int {
	maxFinal := 0
	maxCurrent := 0
	for _, v := range nums {
		if v == 1 {
			maxCurrent++
			maxFinal = max(maxFinal, maxCurrent)
		} else {
			maxCurrent = 0
		}
	}
	return maxFinal
}
