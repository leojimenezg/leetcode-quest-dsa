// ============================================================================
// Backtracking
// ============================================================================

package solutions

// Combinations
//
// Patrones:
//   - Backtracking
//
// Útil cuando:
//   - se generan todas las combinaciones posibles de k elementos
//   - el orden no importa
//   - se puede podar el espacio de búsqueda avanzando el índice de inicio
//
// Complejidad:
//   - Tiempo: O(k * C(n,k)) donde C(n,k) = n! / (k! * (n-k)!)
//   - Espacio: O(k)
func combine(n int, k int) [][]int {
	res := make([][]int, 0, n+k)
	var backtrack func(start int, comb []int)
	backtrack = func(start int, comb []int) {
		if len(comb) == k {
			c := make([]int, k)
			copy(c, comb)
			res = append(res, c)
			return
		}
		for i := start; i <= n; i++ {
			comb = append(comb, i)
			backtrack(i+1, comb)
			comb = comb[:len(comb)-1]
		}
	}
	backtrack(1, []int{})
	return res
}
