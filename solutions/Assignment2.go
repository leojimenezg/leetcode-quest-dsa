// ============================================================================
// Assignment 2
// ============================================================================

package solutions

// Remove Duplicate Letters
// Patrones:
//   - Greedy
//   - Monotonic Stack (increasing)
//
// Útil cuando:
//   - se construye una secuencia lexicográficamente mínima
//   - se permiten rollback de decisiones previas
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func RemoveDuplicateLetters(s string) string {
	uniques := make(map[rune]int)
	for i, l := range s {
		uniques[l] = i
	}
	used := make(map[rune]bool)
	stack := make([]rune, 0, len(s))
	for idx, ltr := range s {
		if used[ltr] {
			continue
		}
		for len(stack) > 0 && ltr < stack[len(stack)-1] && uniques[stack[len(stack)-1]] > idx {
			used[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, ltr)
		used[ltr] = true
	}
	return string(stack)
}
