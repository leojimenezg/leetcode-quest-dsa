// ============================================================================
// LeetCode Quest: Data Structures and Algorithms
// ============================================================================

package problems

import (
	"container/heap"
	"sort"
	"strconv"
	"strings"
)

// ============================================================================
// Array 1
// ============================================================================

// Concatenation of Array
// Patrones:
//   - Array Construction
//
// Útil cuando:
//   - se necesita crear un nuevo arreglo a partir de uno existente
//   - el tamaño final es conocido de antemano
//   - no se requiere lógica condicional compleja
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
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
// Patrones:
//   - Array Construction
//   - Interleaving
//
// Útil cuando:
//   - se combinan dos secuencias con posiciones conocidas
//   - se requiere acceso indexado directo
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func Shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)
	for i := range n {
		ans[2*i] = nums[i]
		ans[2*i+1] = nums[i+n]
	}
	return ans
}

// Max Consecutive Ones
// Patrones:
//   - Single Pass
//   - Running Counter
//
// Útil cuando:
//   - se busca la subsecuencia consecutiva más larga
//   - una condición rompe el conteo y obliga a reiniciar
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
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
// Patrones:
//   - Hash Table
//   - Mathematical Invariants
//
// Útil cuando:
//   - se detectan duplicados y valores faltantes
//   - se conoce el rango esperado de valores
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
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
// Patrones:
//   - Sorting
//   - Value Compression
//
// Útil cuando:
//   - se necesita ranking relativo
//   - O(n log n) es aceptable
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(n)
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
// Patrones:
//   - Index as Hash
//   - Presence Tracking
//
// Útil cuando:
//   - el rango de valores es conocido (1...n)
//   - se puede usar el índice como estructura auxiliar
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
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
// Patrones:
//   - Stack (Call Stack Simulation)
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
// Patrones:
//   - Monotonic Stack (decreciente)
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
// Patrones:
//   - Single Pass
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
// Patrones:
//   - Hash Table
//   - Complement Lookup
//
// Útil cuando:
//   - se busca un par con relación matemática fija
//   - se requiere acceso O(1) por valor
//   - el orden de los elementos no es relevante
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
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
// Patrones:
//   - Counting
//   - Greedy Consumption
//
// Útil cuando:
//   - el orden puede rotarse indefinidamente
//   - solo importa la disponibilidad total
//   - el consumo ocurre de forma secuencial y estricta
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
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
// Patrones:
//   - Mathematical Simulation
//
// Útil cuando:
//   - un proceso repetitivo puede expresarse como suma directa
//   - hay un punto de terminación específico
//   - los elementos antes y después del punto se comportan distinto
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
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
// Patrones:
//   - Two-Stack Queue
//   - Amortized Analysis
//
// Útil cuando:
//   - solo se permiten stacks
//   - se puede amortizar el costo de inversión
//
// Complejidad:
//   - Tiempo: O(1) amortizado
//   - Espacio: O(n)
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

// ============================================================================
// Heap
// ============================================================================

// MaxHeap implementa heap.Interface para un max heap de enteros
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Last Stone Weight
// Patrones:
//   - Max Heap
//   - Priority Queue
//
// Útil cuando:
//   - se necesita acceder repetidamente al mayor elemento
//   - el conjunto cambia dinámicamente
//   - mantener orden total no es necesario
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(n)
func LastStoneWeight(stones []int) int {
	h := MaxHeap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		y := heap.Pop(&h).(int)
		x := heap.Pop(&h).(int)
		if x != y {
			heap.Push(&h, y-x)
		}
	}
	if h.Len() > 0 {
		return h[0]
	}
	return 0
}

type Pair struct {
	I   int
	J   int
	Sum int
}

type MinHeapPairs []Pair

func (h MinHeapPairs) Len() int {
	return len(h)
}

func (h MinHeapPairs) Less(i, j int) bool { return h[i].Sum < h[j].Sum }

func (h MinHeapPairs) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeapPairs) Push(x any) { *h = append(*h, x.(Pair)) }

func (h *MinHeapPairs) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type PairKey struct {
	I int
	J int
}

// Find K Pairs with Smallest Sums
// Patrones:
//   - Min Heap
//   - BFS-like Expansion
//
// Útil cuando:
//   - se explora una matriz virtual ordenada
//   - solo se necesitan los k mejores resultados
//
// Complejidad:
//   - Tiempo: O(k log k)
//   - Espacio: O(k)
func KSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	res := make([][]int, 0, k)
	visited := make(map[PairKey]bool)
	minHeap := MinHeapPairs{}
	heap.Init(&minHeap)
	heap.Push(&minHeap, Pair{0, 0, nums1[0] + nums2[0]})
	visited[PairKey{0, 0}] = true
	for len(res) < k {
		pair := heap.Pop(&minHeap).(Pair)
		res = append(res, []int{nums1[pair.I], nums2[pair.J]})
		if j := pair.J + 1; j < len(nums2) && !visited[PairKey{pair.I, j}] {
			heap.Push(&minHeap, Pair{pair.I, j, nums1[pair.I] + nums2[j]})
			visited[PairKey{pair.I, j}] = true
		}
		if i := pair.I + 1; i < len(nums1) && !visited[PairKey{i, pair.J}] {
			heap.Push(&minHeap, Pair{i, pair.J, nums1[i] + nums2[pair.J]})
			visited[PairKey{i, pair.J}] = true
		}
	}
	return res
}

// Construct Target Array With Multiple Sums
// Patrones:
//   - Reverse Greedy
//   - Max Heap
//   - Mathematical Reduction (modulo)
//
// Útil cuando:
//   - un solo elemento domina el estado final
//   - la simulación directa no escala
//   - el proceso puede revertirse de forma determinista
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(n)
func IsPossible(target []int) bool {
	maxHeap := MaxHeap(target)
	heap.Init(&maxHeap)
	currentSum := 0
	for _, v := range target {
		currentSum += v
	}
	for {
		currentMax := heap.Pop(&maxHeap).(int)
		if currentMax == 1 {
			return true
		}
		othersSum := currentSum - currentMax
		if othersSum <= 0 || othersSum >= currentMax {
			return false
		}
		prevNum := currentMax % othersSum
		if prevNum == 0 && othersSum != 1 {
			return false
		}
		heap.Push(&maxHeap, prevNum)
		currentSum = othersSum + prevNum
	}
}

// ============================================================================
// String
// ============================================================================

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
func DetectCapitalUse(word string) bool {
	allLower := strings.ToLower(word)
	if allLower == word {
		return true
	}
	allUpper := strings.ToUpper(word)
	if allUpper == word {
		return true
	}
	onlyFirst := strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	if onlyFirst == word {
		return true
	}
	return false
}

// License Key Formatting
// Patrones:
//   - String Normalization
//   - Reverse Traversal
//   - Fixed-size Grouping
//   - Post-processing (Reverse Output)
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
func LicenseKeyFormatting(s string, k int) string {
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ToUpper(s)
	count := 0
	var res []byte
	for i := len(s) - 1; i >= 0; i-- {
		if count == k {
			res = append(res, '-')
			count = 0
		}
		res = append(res, s[i])
		count++
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
