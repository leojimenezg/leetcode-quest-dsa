// ============================================================================
// Backtracking
// ============================================================================

package solutions

import (
	"strconv"
)

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

// Restore IP Addresses
//
// Patrones:
//   - Backtracking
//
// Útil cuando:
//   - se particiona un string en segmentos con restricciones estrictas
//   - el número de segmentos es fijo (4 octetos)
//   - se puede podar temprano cuando un segmento es inválido (> 255 o leading zero)
//
// Complejidad:
//   - Tiempo: O(1) — el espacio de búsqueda está acotado (máximo 3^4 = 81 combinaciones)
//   - Espacio: O(1) — profundidad del call stack es siempre 4
func restoreIpAddresses(s string) []string {
	n := len(s)
	res := make([]string, 0)
	if n > 12 {
		return res
	}
	var backtrack func(idx, dots int, curIP string)
	backtrack = func(idx, dots int, curIP string) {
		if dots == 4 && idx == n {
			res = append(res, curIP[:len(curIP)-1])
			return
		} else if dots > 4 {
			return
		}
		for i := idx; i < min(i+3, n); i++ {
			seg, _ := strconv.Atoi(s[idx : i+1])
			if seg < 256 && (s[idx] != '0' || i == idx) {
				backtrack(i+1, dots+1, curIP+s[idx:i+1]+".")
			}
		}
	}
	backtrack(0, 0, "")
	return res
}

// Get the K-th Lexicographically Smallest Happy String
//
// Patrones:
//   - Backtracking
//
// Útil cuando:
//   - se generan strings con restricciones entre caracteres adyacentes
//   - se busca el k-ésimo elemento en orden lexicográfico
//   - se puede podar temprano cuando ya se encontró el resultado
//
// Complejidad:
//   - Tiempo: O(n * 3 * 2^(n-1)) — 3 opciones iniciales, 2 en cada nivel siguiente
//   - Espacio: O(n)
func getHappyString(n int, k int) string {
	happy := [3]byte{'a', 'b', 'c'}
	count, res := 0, ""
	current := make([]byte, 0, n)
	var backtrack func()
	backtrack = func() {
		m := len(current)
		if m == n {
			count++
			if count == k {
				res = string(current)
			}
			return
		}
		for _, v := range happy {
			if m > 0 && current[m-1] == v {
				continue
			}
			current = append(current, v)
			backtrack()
			if res != "" {
				return
			}
			current = current[:len(current)-1]
		}
	}
	backtrack()
	return res
}
