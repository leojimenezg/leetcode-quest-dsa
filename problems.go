// ============================================================================
// LeetCode Quest: Data Structures and Algorithms
// ============================================================================

package problems

import (
	"sort"
	"strconv"
	"strings"
)

// ============================================================================
// Array 1
// ============================================================================

// Concatenation of Array
// Patrón: Array construction/building
// Útil cuando:
//   - se requiere un nuevo array
//   - el tamaño del nuevo array es predecible
func GetConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, n*2)
	for i, v := range nums {
		ans[i] = v
		ans[i+n] = v
	}
	return ans
}

// Shuffle the Array
// Patrón: Array Construction + Interleaving (intercalado)
// Útil cuando:
//   - se necesita mezclar/intercalar dos secuencias
//   - el tamaño resultante es predecible
func Shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)
	for i := range n {
		ans[2*i] = nums[i]
		ans[2*i+1] = nums[i+n]
	}
	return ans
}

// Max Consecutive Ones
// Patrón: Sliding Window + Running Maximum
// Útil cuando:
//   - se necesita encontrar la secuencia consecutiva más larga
//   - se debe resetear un contador cuando se rompe una condición
//   - se mantiene track del mejor resultado visto hasta ahora
func FindMaxConsecutiveOnes(nums []int) int {
	currentCount := 0
	lastMaxCount := 0
	for _, v := range nums {
		if v == 1 {
			currentCount++
			lastMaxCount = max(lastMaxCount, currentCount)
		} else {
			currentCount = 0
		}
	}
	return lastMaxCount
}

// ============================================================================
// Array 2
// ============================================================================

// Set Mismatch
// Patrón: Hash Table + Math Formula
// Útil cuando:
//   - se necesita detectar duplicados en O(n)
//   - se pueden usar fórmulas matemáticas para encontrar valores faltantes
//   - espacio extra O(n) es aceptable
func FindErrorNums(nums []int) []int {
	n := len(nums)
	expectedSum := n * (n + 1) / 2
	repeatedNum := -1
	actualSum := 0
	seenNums := make(map[int]bool)
	for _, v := range nums {
		if _, ok := seenNums[v]; ok {
			repeatedNum = v
		}
		seenNums[v] = true
		actualSum += v
	}
	lossNum := expectedSum - (actualSum - repeatedNum)
	return []int{repeatedNum, lossNum}
}

// How Many Numbers Are Smaller Than the Current Number
// Patrón: Sorting + Hash Table
// Útil cuando:
//   - se necesita comparar cada elemento con todos los demás
//   - se puede sacrificar posiciones originales temporalmente
//   - O(n log n) es aceptable vs O(n²)
func SmallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)
	count := make(map[int]int)
	for i, v := range sorted {
		if _, exists := count[v]; !exists {
			count[v] = i
		}
	}
	for i, v := range nums {
		sorted[i] = count[v]
	}
	return sorted
}

// Find All Numbers Disappeared in an Array
// Patrón: Index as Hash + Array Construction
// Útil cuando:
//   - rango de valores conocido (1 a n)
//   - se necesita rastrear presencia/ausencia
//   - O(n) tiempo y espacio es aceptable
func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	nonRepeated := make([]bool, n)
	for _, v := range nums {
		nonRepeated[v-1] = true
	}
	missing := make([]int, 0)
	for i, v := range nonRepeated {
		if !v {
			missing = append(missing, i+1)
		}
	}
	return missing
}

// ============================================================================
// Stack
// ============================================================================

// Build an Array With Stack Operations
// Patrón: Simulation + Two Pointers
// Útil cuando:
//   - se simula un proceso paso a paso
//   - se necesita sincronizar dos secuencias ordenadas
//   - los punteros avanzan con lógica relacionada
func BuildArray(target []int, n int) []string {
	idx := 0
	ops := make([]string, 0)
	for v := 1; v <= n; v++ {
		if v == target[idx] {
			ops = append(ops, "Push")
			idx++
			if idx == len(target) {
				break
			}
		} else {
			ops = append(ops, "Push", "Pop")
		}
	}
	return ops
}

// Evaluate Reverse Polish Notation
// Patrón: Stack (LIFO)
// Útil cuando:
//   - evaluación de expresiones postfix/prefix
//   - se necesita procesar elementos en orden inverso al ingreso
//   - operaciones dependen de los últimos N elementos
func EvalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case "+":
			res := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = res
		case "-":
			res := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = res
		case "*":
			res := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = res
		case "/":
			res := int(stack[len(stack)-2] / stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = res
		default:
			n, _ := strconv.Atoi(token)
			stack = append(stack, n)
		}
	}
	return stack[0]
}

// Exclusive Time of Functions
// Patrón: Stack (LIFO) + Time Tracking
// Útil cuando:
//   - se modela call stack de funciones anidadas
//   - se necesita calcular tiempos exclusivos con pausas/resumes
//   - operaciones tienen inicio y fin que deben balancearse
func ExclusiveTime(n int, logs []string) []int {
	result := make([]int, n)
	stack := make([]int, 0)
	prevTime := 0
	for _, log := range logs {
		d := strings.Split(log, ":")
		id, _ := strconv.Atoi(d[0])
		timestamp, _ := strconv.Atoi(d[2])
		switch d[1] {
		case "start":
			if len(stack) > 0 {
				prevFunc := stack[len(stack)-1]
				result[prevFunc] += timestamp - prevTime
			}
			prevTime = timestamp
			stack = append(stack, id)
		case "end":
			prevFunc := stack[len(stack)-1]
			result[prevFunc] += timestamp - prevTime + 1
			prevTime = timestamp + 1
			stack = stack[:len(stack)-1]
		}
	}
	return result
}
