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

// ============================================================================
// Monotonic Stack
// ============================================================================

// Final Prices With a Special Discount in a Shop
// Patrón: Monotonic Stack
// Útil cuando:
//   - se busca el próximo menor/mayor elemento
//   - evitar comparaciones redundantes en peor caso
//   - garantizar O(n) en lugar de O(n²)
func FinalPrices(prices []int) []int {
	n := len(prices)
	res := make([]int, n)
	copy(res, prices)
	stack := make([]int, 0)
	for currentProd := range n {
		for len(stack) > 0 && prices[currentProd] <= prices[stack[len(stack)-1]] {
			prevProd := stack[len(stack)-1]
			res[prevProd] = prices[prevProd] - prices[currentProd]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, currentProd)
	}
	return res
}

// Daily Temperatures
// Patrón: Monotonic Stack (Decreasing)
// Útil cuando:
//   - se busca próximo mayor elemento
//   - calcular distancia hasta ese elemento
//   - garantizar O(n)
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	stack := make([]int, 0)
	for i := range n {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			answer[idx] = i - idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return answer
}

// Largest Rectangle in Histogram
// Patrón: Monotonic Stack (Increasing)
// Útil cuando:
//   - se calcula área/ancho máximo con restricción de altura
//   - se necesita encontrar límites izquierdo y derecho para cada elemento
//   - garantizar O(n)
func LargestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([]int, 0)
	for i, currentHeight := range heights {
		for len(stack) > 0 && currentHeight < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			var width int
			if len(stack) < 1 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		height := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		var width int
		if len(stack) == 0 {
			width = len(heights)
		} else {
			width = len(heights) - stack[len(stack)-1] - 1
		}
		area := height * width
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// ============================================================================
// Assignment 1
// ============================================================================

// Plus One
// Patrón: Digit-by-Digit Arithmetic (Simulation)
// Útil cuando:
//   - se simulan operaciones aritméticas manuales (suma, resta, multiplicación)
//   - se procesan dígitos de derecha a izquierda con propagación de acarreo
//   - se trabaja con números representados como arrays/strings/linked lists
func PlusOne(digits []int) []int {
	n := len(digits)
	if digits[n-1] != 9 {
		digits[n-1]++
		return digits
	}
	carry := false
	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			carry = true
		} else {
			if carry {
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
// Patrón: Single Pass + State Machine (Fase de subida → Fase de bajada)
// Útil cuando:
//   - se verifica una secuencia que cambia de dirección una sola vez
//   - se necesita detectar transiciones entre estados (subir/bajar)
//   - se valida monotonía con un punto de inflexión
func ValidMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	down := false
	amount := 0
	for i := 1; i < n; i++ {
		idx := i - 1
		if arr[i] == arr[idx] {
			return false
		} else if arr[i] > arr[idx] {
			if down {
				return false
			}
			amount++
		} else {
			down = true
			if amount < 1 {
				return false
			}
		}
	}
	if !down {
		return false
	}
	return true
}

// ============================================================================
// Assignment 2
// ============================================================================

// Remove Duplicate Letters
// Patrón: Greedy + Monotonic Stack (Increasing)
// Útil cuando:
//   - se construye una secuencia lexicográficamente óptima
//   - se pueden "deshacer" decisiones si aparece algo mejor después
//   - se necesita eliminar duplicados manteniendo orden relativo
func RemoveDuplicateLetters(s string) string {
	lastIndex := make(map[rune]int)
	for i, l := range s {
		lastIndex[l] = i
	}
	used := make(map[rune]bool)
	stack := make([]rune, 0)
	for idx, lt := range s {
		if used[lt] {
			continue
		}
		for len(stack) > 0 &&
			lt < stack[len(stack)-1] &&
			lastIndex[stack[len(stack)-1]] > idx {
			stackTop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			used[stackTop] = false
		}
		stack = append(stack, lt)
		used[lt] = true
	}
	return string(stack)
}

// ============================================================================
// Why not?
// ============================================================================

// Two Sum
// Patrón: Hash Table
// Útil cuando:
//   - se busca un par de elementos con relación matemática (suma, diferencia, producto)
//   - se necesita acceso O(1) a elementos por valor
//   - el orden relativo no importa
func TwoSum(nums []int, target int) []int {
	numbers := make(map[int]int)
	for i, v := range nums {
		numbers[v] = i
	}
	for idx, num := range nums {
		diff := target - num
		if i, ok := numbers[diff]; ok && i != idx {
			return []int{idx, i}
		}
	}
	return []int{}
}

// ============================================================================
// Queue
// ============================================================================

// Number of Students Unable to Eat Lunch
// Patrón: Counting + Sequential Matching
// Útil cuando:
//   - elementos pueden reordenarse libremente (rotación infinita)
//   - solo importa disponibilidad total, no orden específico
//   - hay procesamiento secuencial estricto del otro lado
func CountStudents(students []int, sandwiches []int) int {
	countEach := [2]int{}
	for _, sandwichType := range students {
		countEach[sandwichType]++
	}
	for _, sandwichType := range sandwiches {
		if countEach[sandwichType] > 0 {
			countEach[sandwichType]--
		} else {
			break
		}
	}
	return countEach[0] + countEach[1]
}

// Time Needed to Buy Tickets
// Patrón: Mathematical Simulation (cálculo directo sin simular paso a paso)
// Útil cuando:
//   - se simula un proceso circular pero se puede calcular el resultado directamente
//   - hay un punto de terminación específico (persona k termina)
//   - personas antes/después del punto tienen comportamientos diferentes
func TimeRequiredToBuy(tickets []int, k int) int {
	totalTime := 0
	for i := range tickets {
		if i < k {
			totalTime += min(tickets[i], tickets[k])
		} else if i == k {
			totalTime += tickets[i]
		} else {
			totalTime += min(tickets[i], tickets[k]-1)
		}
	}
	return totalTime
}

// Implement Queue using Stacks
// Patrón: Two-Stack Queue (lazy transfer)
// Útil cuando:
//   - se necesita implementar Queue (FIFO) usando solo Stacks (LIFO)
//   - se quiere amortizar el costo de inversión de orden
//   - las operaciones Pop/Peek son menos frecuentes que Push
type MyQueue struct {
	StackInput  []int
	StackOutput []int
	Size        int
}

func Constructor() MyQueue {
	return MyQueue{
		StackInput:  make([]int, 0),
		StackOutput: make([]int, 0),
		Size:        0,
	}
}

func (this *MyQueue) Push(x int) {
	this.StackInput = append(this.StackInput, x)
	this.Size++
}

func (this *MyQueue) Pop() int {
	if len(this.StackOutput) < 1 {
		this.revealBottom()
	}
	first := this.StackOutput[len(this.StackOutput)-1]
	this.StackOutput = this.StackOutput[:len(this.StackOutput)-1]
	this.Size--
	return first
}

func (this *MyQueue) Peek() int {
	if len(this.StackOutput) < 1 {
		this.revealBottom()
	}
	return this.StackOutput[len(this.StackOutput)-1]
}

func (this *MyQueue) Empty() bool {
	return this.Size == 0
}

// revealBottom transfiere todos los elementos de un stack a otro
// Esto invierte el orden: el más antiguo queda en top del stack
// También conocido como "lazy transfer" en el patrón Two-Stack Queue
func (this *MyQueue) revealBottom() {
	for len(this.StackInput) > 0 {
		this.StackOutput = append(
			this.StackOutput, this.StackInput[len(this.StackInput)-1],
		)
		this.StackInput = this.StackInput[:len(this.StackInput)-1]
	}
}
