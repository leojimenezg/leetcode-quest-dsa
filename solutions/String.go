// ============================================================================
// String
// ============================================================================

package solutions

import "strings"

// Detect Capital
// Patrones:
//   - String Validation
//   - Case Analysis
//
// Útil cuando:
//   - se valida el formato de una cadena según reglas explícitas
//   - existen pocos estados válidos y son mutuamente excluyentes
//   - la claridad es más importante que la optimización prematura
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n) por creación de strings auxiliares
func detectCapitalUse(word string) bool {
	upper := strings.ToUpper(word)
	if upper == word {
		return true
	}
	lower := strings.ToLower(word)
	if lower == word {
		return true
	}
	capitalize := upper[:1] + lower[1:]
	if capitalize == word {
		return true
	}
	return false
}

// License Key Formatting
// Patrones:
//   - String Normalization
//   - Reverse Traversal
//   - Fixed-size Grouping
//
// Útil cuando:
//   - la entrada necesita limpieza previa (eliminar caracteres, unificar formato)
//   - el agrupamiento depende del final de la secuencia
//   - procesar desde adelante introduce casos especiales innecesarios
//   - es más simple construir el resultado en orden inverso y corregir al final
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func licenseKeyFormatting(s string, k int) string {
	n := len(s)
	upper := strings.ToUpper(s)
	res := make([]byte, 0, n)
	for idx, count := n-1, 0; idx >= 0; idx-- {
		if upper[idx] == '-' {
			continue
		}
		if count == k {
			res = append(res, '-')
			count = 0
		}
		res = append(res, upper[idx])
		count++
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

// Masking Personal Information
// Patrones:
//   - Case Analysis
//   - Input Normalization (Structure Cleaning)
//   - String Validation
//   - String Masking / Redaction
//
// Útil cuando:
//   - la entrada puede representar distintos tipos de datos
//   - cada tipo requiere reglas de formato distintas
//   - es necesario limpiar símbolos irrelevantes antes de procesar
//   - se debe ocultar información sensible manteniendo partes visibles
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func maskPII(s string) string {
	atIdx := strings.Index(s, "@")
	var res strings.Builder
	if atIdx > 0 {
		lower := strings.ToLower(s)
		res.WriteString(lower[:1] + "*****" + lower[atIdx-1:])
	} else {
		phone := []byte(s)
		for i, l := range phone {
			if l == '+' || l == '-' || l == '(' || l == ')' {
				phone[i] = ' '
			}
		}
		cleaned := strings.ReplaceAll(string(phone), " ", "")
		n := len(cleaned)
		cc := n - 10
		if cc > 0 {
			res.WriteRune('+')
			for range cc {
				res.WriteRune('*')
			}
			res.WriteRune('-')
		}
		res.WriteString("***-***-" + cleaned[n-4:])
	}
	return res.String()
}
