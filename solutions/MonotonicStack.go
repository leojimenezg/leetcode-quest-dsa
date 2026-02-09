// ============================================================================
// Monotonic Stack
// ============================================================================

package solutions

// Final Prices With a Special Discount in a Shop
// Patrones:
//   - Monotonic Stack (Increasing)
//
// Útil cuando:
//   - se busca el siguiente elemento menor o igual
//   - se quieren evitar comparaciones redundantes
//   - se necesita resolver en O(n) en lugar de O(n²)
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func FinalPrices(prices []int) []int {
	n := len(prices)
	answer := make([]int, n)
	copy(answer, prices)
	stack := make([]int, 0, n)
	for current := range n {
		for len(stack) > 0 && prices[current] <= prices[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			answer[last] -= prices[current]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, current)
	}
	return answer
}

// Daily Temperatures
// Patrones:
//   - Monotonic Stack (Decreasing)
//
// Útil cuando:
//   - se busca el siguiente elemento mayor
//   - se requiere distancia entre índices
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	stack := make([]int, 0, n)
	for current := range n {
		for len(stack) > 0 && temperatures[current] > temperatures[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			answer[last] = current - last
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, current)
	}
	return answer
}

// Largest Rectangle in Histogram
// Patrones:
//   - Monotonic Stack (Increasing)
//
// Útil cuando:
//   - se necesita calcular áreas máximas bajo restricciones locales
//   - se buscan límites izquierdo y derecho para cada elemento
//   - una decisión depende del primer elemento menor a ambos lados
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func LargestRectangleArea(heights []int) int {
	n := len(heights)
	res := 0
	stack := make([]int, 0, n)
	for i, current := range heights {
		for len(stack) > 0 && current < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := height * width
			if area > res {
				res = area
			}
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		height := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		width := n
		if len(stack) > 0 {
			width = n - stack[len(stack)-1] - 1
		}
		area := height * width
		if area > res {
			res = area
		}
	}
	return res
}
