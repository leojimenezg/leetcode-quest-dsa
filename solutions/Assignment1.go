// ============================================================================
// Assignment 1
// ============================================================================

package solutions

// Plus One
// Patrones:
//   - Digit-by-Digit Arithmetic
//   - Simulation
//
// Útil cuando:
//   - se simulan operaciones aritméticas manuales
//   - hay propagación de acarreo
//   - los números están representados como arrays o strings
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func PlusOne(digits []int) []int {
	n := len(digits)
	carry := true
	for i := n - 1; i >= 0; i-- {
		if carry {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i]++
				carry = false
				break
			}
		}
	}
	if carry {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}

// Valid Mountain Array
// Patrones:
//   - State Machine
//
// Útil cuando:
//   - una secuencia cambia de dirección solo una vez
//   - se validan transiciones entre estados (subida → bajada)
//   - se requiere detectar un único punto de inflexión
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func ValidMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	up := false
	down := false
	prev := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] == prev {
			return false
		}
		if arr[i] < prev {
			if !up {
				return false
			}
			down = true
		}
		if arr[i] > prev {
			if down {
				return false
			}
			up = true
		}
		prev = arr[i]
	}
	return down
}
