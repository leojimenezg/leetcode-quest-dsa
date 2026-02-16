// ============================================================================
// String Matching
// ============================================================================

package solutions

import "strings"

// Repeated Substring Pattern
// Patrones:
//   - String Periodicity
//   - Prefix-based Validation
//   - Modulo Indexing
//
// Útil cuando:
//   - se necesita detectar si una secuencia completa está formada por repeticiones de un mismo bloque
//   - el orden global de los elementos es crítico
//   - contar frecuencias no es suficiente para validar la estructura
//   - la longitud del patrón debe dividir exactamente la longitud total
//
// Complejidad:
//   - Tiempo: O(n^2)
//   - Espacio: O(1)
func RepeatedSubstringPattern(s string) bool {
	n := len(s)
	for length := 1; length <= n/2; length++ {
		if n%length != 0 {
			continue
		}
		sub := s[:length]
		valid := true
		for i := length; i < n; i++ {
			if s[i] != sub[i%length] {
				valid = false
				break
			}
		}
		if valid {
			return true
		}
	}
	return false
}

// Rotate String
// Patrones:
//   - String Rotation
//   - Substring Search
//   - Property-based Reduction
//
// Útil cuando:
//   - se necesita verificar si una secuencia es una rotación de otra
//   - simular transformaciones paso a paso resulta costoso o innecesario
//   - existe una propiedad global que agrupa todos los estados posibles
//   - el problema puede reducirse a una búsqueda de substring
//
// Complejidad:
//   - Tiempo: O(n) en promedio (dependiente del algoritmo de búsqueda de substring)
//   - Espacio: O(n) por la concatenación de strings
func RotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	duplicated := s + s
	return strings.Contains(duplicated, goal)
}

// Repeated String Match
// Patrones:
//   - String Repetition
//   - Substring Search
//   - Ceiling Division
//
// Útil cuando:
//   - un string base puede repetirse un número arbitrario de veces
//   - se necesita verificar si otro string aparece como substring
//   - el problema permite acotar matemáticamente el número de repeticiones
//   - el caso difícil ocurre cuando el match cruza fronteras entre repeticiones
//   - simular repeticiones infinitas es innecesario o incorrecto
//
// Complejidad:
//   - Tiempo: O(n * m) en el peor caso
//   - Espacio: O(n) por la construcción del string repetido
func RepeatedStringMatch(a string, b string) int {
	minReps := (len(b) + len(a) - 1) / len(a) // Ceiling division
	var reps strings.Builder
	for range minReps {
		reps.WriteString(a)
	}
	if strings.Contains(reps.String(), b) {
		return minReps
	}
	reps.WriteString(a)
	if strings.Contains(reps.String(), b) {
		return minReps + 1
	}
	return -1
}
