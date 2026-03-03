// ============================================================================
// Stack
// ============================================================================

package solutions

import (
	"strconv"
	"strings"
)

// Build an Array With Stack Operations
// Patrones:
//   - Simulation
//   - Two Pointers
//
// Útil cuando:
//   - se simula un proceso paso a paso con reglas fijas
//   - se sincronizan dos secuencias crecientes
//   - un puntero avanza condicionado al otro
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func buildArray(target []int, n int) []string {
	ops := make([]string, 0, 2*n)
	idx := 0
	for v := 1; v <= n; v++ {
		if v == target[idx] {
			ops = append(ops, "Push")
			idx++
		} else {
			ops = append(ops, "Push", "Pop")
		}
		if idx == len(target) {
			break
		}
	}
	return ops
}

// Evaluate Reverse Polish Notation
// Patrones:
//   - Stack (LIFO)
//
// Útil cuando:
//   - se evalúan expresiones postfix/prefix
//   - las operaciones dependen de los últimos operandos ingresados
//   - se requiere deshacer consumo parcial de datos
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		n := len(stack)
		op := 0
		switch token {
		case "+":
			op = stack[n-2] + stack[n-1]
		case "-":
			op = stack[n-2] - stack[n-1]
		case "*":
			op = stack[n-2] * stack[n-1]
		case "/":
			if stack[n-1] != 0 {
				op = int(stack[n-2] / stack[n-1])
			}
		default:
			op, _ = strconv.Atoi(token)
			stack = append(stack, op)
			continue
		}
		stack = stack[:n-1]
		stack[len(stack)-1] = op
	}
	return stack[0]
}

// Exclusive Time of Functions
// Patrones:
//   - Stack (LIFO)
//   - Interval / Time Tracking
//
// Útil cuando:
//   - se modela ejecución anidada de procesos o funciones
//   - hay eventos start/end bien definidos
//   - se requiere excluir tiempo de ejecuciones hijas
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func exclusiveTime(n int, logs []string) []int {
	times := make([]int, n)
	stack := make([]int, 0, len(logs))
	prevTime := 0
	for _, log := range logs {
		info := strings.Split(log, ":")
		id, _ := strconv.Atoi(info[0])
		tm, _ := strconv.Atoi(info[2])
		switch info[1] {
		case "start":
			if len(stack) > 0 {
				prevFunc := stack[len(stack)-1]
				times[prevFunc] += tm - prevTime
			}
			prevTime = tm
			stack = append(stack, id)
		case "end":
			times[id] += tm - prevTime + 1
			prevTime = tm + 1
			stack = stack[:len(stack)-1]
		}
	}
	return times
}
