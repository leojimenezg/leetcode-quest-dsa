// ============================================================================
// Assignment 4
// ============================================================================

package solutions

import "strings"

// Reformat Date
//
// Patrones:
//   - Input Normalization
//   - Field Reordering
//
// Útil cuando:
//   - la entrada es semiestructurada
//   - el output exige formato estricto
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func reformatDate(date string) string {
	elements := strings.Split(date, " ")
	switch elements[1] {
	case "Jan":
		elements[1] = "01"
	case "Feb":
		elements[1] = "02"
	case "Mar":
		elements[1] = "03"
	case "Apr":
		elements[1] = "04"
	case "May":
		elements[1] = "05"
	case "Jun":
		elements[1] = "06"
	case "Jul":
		elements[1] = "07"
	case "Aug":
		elements[1] = "08"
	case "Sep":
		elements[1] = "09"
	case "Oct":
		elements[1] = "10"
	case "Nov":
		elements[1] = "11"
	case "Dec":
		elements[1] = "12"
	}
	day := []rune(elements[0])
	day = day[:len(day)-2]
	if len(day) < 2 {
		elements[0] = "0" + string(day)
	} else {
		elements[0] = string(day)
	}
	return elements[2] + "-" + elements[1] + "-" + elements[0]
}

// Maximum Repeating Substring
//
// Patrones:
//   - Incremental String Construction
//   - Substring Search
//
// Útil cuando:
//   - se busca el máximo número de repeticiones consecutivas
//   - el patrón base es fijo
//   - el límite de repeticiones está acotado por el tamaño de la entrada
//
// Complejidad:
//   - Tiempo: O(n^2)
//   - Espacio: O(n)
func maxRepeating(sequence string, word string) int {
	k := len(sequence) / len(word)
	var repeated strings.Builder
	for i := range k {
		repeated.WriteString(word)
		if !strings.Contains(sequence, repeated.String()) {
			return i
		}
	}
	return k
}
