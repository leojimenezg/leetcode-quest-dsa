// ============================================================================
// LeetCode Quest: Data Structures and Algorithms
// ============================================================================

package problems

import (
	"container/heap"
	"strconv"
	"strings"
)

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
func MaskPII(s string) string {
	symbolIdx := strings.Index(s, "@")
	var res strings.Builder
	if symbolIdx < 0 {
		s = strings.ReplaceAll(s, "+", "")
		s = strings.ReplaceAll(s, "-", "")
		s = strings.ReplaceAll(s, "(", "")
		s = strings.ReplaceAll(s, ")", "")
		s = strings.ReplaceAll(s, " ", "")
		n := len(s)
		res.WriteString("***-***-" + s[n-4:])
		if diff := n - 10; diff > 0 {
			var code strings.Builder
			code.WriteRune('+')
			for range diff {
				code.WriteRune('*')
			}
			code.WriteRune('-')
			code.WriteString(res.String())
			return code.String()
		}
		return res.String()
	} else {
		s = strings.ToLower(s)
		res.WriteString(s[:1] + "*****" + s[symbolIdx-1:])
		return res.String()
	}
}

// ============================================================================
// String Matching
// ============================================================================

// Repeated Substring Pattern
// Patrones:
//   - String Periodicity (core pattern)
//   - Prefix-based Validation (deriving candidate pattern)
//   - Modulo Indexing (efficient cyclic comparison)
//
// Útil cuando:
//   - se necesita detectar si una secuencia completa está formada por repeticiones de un mismo bloque
//   - el orden global de los elementos es crítico
//   - contar frecuencias no es suficiente para validar la estructura
//   - la longitud del patrón debe dividir exactamente la longitud total
//
// Complejidad:
//   - Tiempo: O(n²) en el peor caso (prueba de múltiples longitudes candidatas)
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
//   - String Rotation via Substring Search
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
	ss := s + s
	return strings.Contains(ss, goal)
}

// Repeated String Match
// Patrones:
//   - String Repetition (Bounded)
//   - Substring Search
//   - Ceiling Division (Minimum Blocks to Cover a Length)
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
	minReps := (len(b) + len(a) - 1) / len(a)
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

// ============================================================================
// Assignment 3
// ============================================================================

type Apple struct {
	ExpireDay int
	Count     int
}

type MinHeapApple []Apple

func (h MinHeapApple) Len() int { return len(h) }

func (h MinHeapApple) Less(i, j int) bool { return h[i].ExpireDay < h[j].ExpireDay }

func (h MinHeapApple) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeapApple) Push(x any) { *h = append(*h, x.(Apple)) }

func (h *MinHeapApple) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Maximum Number of Eaten Apples
// Patrones:
//   - MinHeap / Priority Queue
//   - Greedy
//
// Útil cuando:
//   - siempre conviene consumir primero el recurso con menor deadline
//   - el input es finito pero el consumo puede extenderse más allá
//   - es necesario mantener dinámicamente el “más urgente”
//
// Complejidad:
//   - Tiempo: O(n log n), cada lote entra y sale del heap una vez
//   - Espacio: O(n) por los lotes activos en el heap
func EatenApples(apples []int, days []int) int {
	n := len(apples)
	res := 0
	minHeap := MinHeapApple{}
	heap.Init(&minHeap)
	idx := 0
	for idx < n || minHeap.Len() > 0 {
		if idx < n && apples[idx] > 0 {
			ed := idx + days[idx]
			heap.Push(&minHeap, Apple{ExpireDay: ed, Count: apples[idx]})
		}
		for minHeap.Len() > 0 {
			top := minHeap[0]
			if top.ExpireDay > idx {
				break
			}
			heap.Pop(&minHeap)
		}
		if minHeap.Len() > 0 {
			res++
			latest := heap.Pop(&minHeap).(Apple)
			latest.Count--
			if latest.Count > 0 {
				heap.Push(&minHeap, latest)
			}
		}
		idx++
	}
	return res
}

// Design Circular Queue
// Patrones:
//   - Circular Queue
//   - Modular Indexing
//   - Array-based Data Structure
//
// Útil cuando:
//   - se necesita una queue FIFO con capacidad fija
//   - se requiere O(1) para enqueue y dequeue
//   - el espacio debe reutilizarse de forma controlada
//   - es necesario mantener índices dentro de un rango circular fijo
//
// Complejidad:
//   - Tiempo: O(1) por operación
//   - Espacio: O(k), donde k es la capacidad fija de la queue
type MyCircularQueue struct {
	Items []int
	IdxF  int
	IdxR  int
	Size  int
	Cap   int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{Items: make([]int, k), Cap: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.Size == this.Cap {
		return false
	}
	this.Items[this.IdxR] = value
	this.IdxR = (this.IdxR + 1) % this.Cap
	this.Size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.Size == 0 {
		return false
	}
	this.IdxF = (this.IdxF + 1) % this.Cap
	this.Size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.Size == 0 {
		return -1
	}
	return this.Items[this.IdxF]
}

func (this *MyCircularQueue) Rear() int {
	if this.Size == 0 {
		return -1
	}
	idx := (this.IdxR - 1 + this.Cap) % this.Cap
	return this.Items[idx]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Size == this.Cap
}

// ============================================================================
// Assignment 4
// ============================================================================

// Reformat Date
// Patrones:
//   - Input Normalization
//   - Canonical Representation
//   - Finite Mapping
//   - Field Reordering
//
// Útil cuando:
//   - la entrada es semiestructurada
//   - múltiples representaciones apuntan al mismo valor lógico
//   - el output exige formato estricto
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func ReformatDate(date string) string {
	elements := strings.Split(date, " ")
	elements[0] = strings.ReplaceAll(elements[0], "st", "")
	elements[0] = strings.ReplaceAll(elements[0], "nd", "")
	elements[0] = strings.ReplaceAll(elements[0], "rd", "")
	elements[0] = strings.ReplaceAll(elements[0], "th", "")
	if len(elements[0]) < 2 {
		elements[0] = "0" + elements[0]
	}
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
	return elements[2] + "-" + elements[1] + "-" + elements[0]
}

// Maximum Repeating Substring
// Patrones:
//   - Incremental Pattern Construction
//   - Repeated Concatenation Validation
//   - Substring Search
//
// Útil cuando:
//   - se busca el máximo número de repeticiones consecutivas
//   - el patrón base es fijo
//   - el límite de repeticiones está acotado por el tamaño de la entrada
//
// Complejidad:
//   - Tiempo: O(n²) en el peor caso
//   - Espacio: O(n)
func MaxRepeating(sequence string, word string) int {
	k := len(sequence) / len(word)
	var str strings.Builder
	for i := range k {
		str.WriteString(word)
		if !strings.Contains(sequence, str.String()) {
			return i
		}
	}
	return k
}

// ============================================================================
// Linked List
// ============================================================================

type ListNode struct {
	Val  int
	Next *ListNode
}

// Remove Duplicates from Sorted List
// Patrones:
//   - Single-Pass Traversal
//   - In-place Deduplication
//   - Sorted Input Invariant
//
// Útil cuando:
//   - los duplicados son contiguos (ordenada)
//   - se permite modificar la estructura original
//   - se requiere O(1) espacio extra
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func DeleteDuplicates(head *ListNode) *ListNode {
	var prevNode *ListNode
	currentNode := head
	for currentNode != nil {
		if prevNode == nil {
			prevNode = currentNode
			goto next
		}
		if currentNode.Val == prevNode.Val {
			prevNode.Next = currentNode.Next
		} else {
			prevNode = currentNode
		}
	next:
		currentNode = currentNode.Next
	}
	return head
}

// Odd Even Linked List
//
// Patrones:
//   - Single-Pass Traversal
//   - In-place List Partitioning
//   - Pointer Rewiring
//
// Útil cuando:
//   - se requiere reagrupar nodos según su posición (odd / even)
//   - la estructura solo permite avanzar en una dirección (listas enlazadas)
//   - no se puede regresar, reiniciar o usar accesos aleatorios
//   - no se permite usar memoria extra proporcional a n
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddCurrent := head
	evenHead := head.Next
	evenCurrent := evenHead
	for evenCurrent != nil && evenCurrent.Next != nil {
		oddCurrent.Next = evenCurrent.Next
		oddCurrent = oddCurrent.Next
		evenCurrent.Next = oddCurrent.Next
		evenCurrent = evenCurrent.Next
	}
	oddCurrent.Next = evenHead
	return head
}

// Reverse Linked List
//
// Patrones:
//   - Single-Pass Traversal
//   - Pointer Reversal (Rewiring)
//   - In-place Transformation
//
// Útil cuando:
//   - se requiere invertir una lista enlazada
//   - la estructura solo permite avanzar en una dirección
//   - no se permite usar memoria extra proporcional a n
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func ReverseList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	for current != nil {
		next := current.Next
		current.Next = prev
		if next == nil {
			head = current
		}
		prev = current
		current = next
	}
	return head
}

// ============================================================================
// Hash
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

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Copy List with Random Pointer
//
// Patrones:
//   - Hash Table
//   - Two-Pass Traversal
//   - Deep Copy
//
// Útil cuando:
//   - se requiere clonar una estructura enlazada con referencias arbitrarias
//   - existen punteros que no siguen una relación lineal
//   - es necesario mantener una correspondencia 1-a-1 entre nodos originales y copiados
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func CopyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodes := make(map[*Node]*Node) // Original: Copy
	for current := head; current != nil; current = current.Next {
		nodes[current] = &Node{Val: current.Val}
	}
	for original := head; original != nil; original = original.Next {
		cpy := nodes[original]
		cpy.Next = nodes[original.Next]
		cpy.Random = nodes[original.Random]
	}
	return nodes[head]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// First Missing Positive
//
// Patrones:
//   - Index as Hash
//   - In-place Transformation
//   - Multi-Pass Traversal
//
// Útil cuando:
//   - se trabaja con un array no ordenado
//   - el rango de valores relevantes está acotado por la longitud del array
//   - se requiere detectar presencia / ausencia de elementos
//   - el problema exige explícitamente espacio extra O(1)
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func FirstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}
	for i := range nums {
		v := abs(nums[i])
		if v <= n {
			if nums[v-1] > 0 {
				nums[v-1] = -nums[v-1]
			}
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

// ============================================================================
// Prefix Sum
// ============================================================================

// Find the Highest Altitude
//
// Patrones:
//   - Prefix Sum (optimizado)
//   - Running Maximum
//   - Single Pass Traversal
//
// Útil cuando:
//   - se calculan sumas acumuladas progresivamente
//   - solo importa el valor máximo/mínimo de las sumas (no todas)
//   - se puede procesar en una sola pasada sin almacenar histórico
//   - se busca optimizar espacio evitando arrays auxiliares
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func LargestAltitude(gain []int) int {
	prev := 0
	res := prev
	for i := 0; i < len(gain); i++ {
		prev = prev + gain[i]
		if prev > res {
			res = prev
		}
	}
	return res
}

// Make Sum Divisible by P
//
// Patrones:
//   - Prefix Sum (con Módulo)
//   - Hash Table (Complement Lookup)
//   - Mathematical Reduction (propiedades del módulo)
//
// Útil cuando:
//   - se busca un subarray cuya suma cumple una condición de módulo
//   - se necesita encontrar el subarray más corto que satisface la condición
//   - se puede reformular como búsqueda de prefix sums previos
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func MinSubarray(nums []int, p int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	remainder := total % p
	if remainder == 0 {
		return 0
	}
	mods := make(map[int]int)
	mods[0] = -1
	sum := 0
	res := len(nums)
	for i, v := range nums {
		sum += v
		currentMod := sum % p
		targetMod := (currentMod - remainder + p) % p
		if idx, ok := mods[targetMod]; ok {
			length := i - idx
			if length < res {
				res = length
			}
		}
		mods[currentMod] = i
	}
	if res >= len(nums) {
		return -1
	}
	return res
}

// Ways to Make a Fair Array
//
// Patrones:
//   - Prefix Sum (bidireccional)
//   - Running Totals
//   - Position Swap Analysis
//
// Útil cuando:
//   - remover un elemento afecta las posiciones de elementos posteriores
//   - se necesita evaluar todas las posiciones sin simulación O(n²)
//   - los elementos se dividen en dos categorías (par/impar, positivo/negativo)
//   - se puede precalcular totales y ajustarlos dinámicamente
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func WaysToMakeFair(nums []int) int {
	totalEven, totalOdd := 0, 0
	for i, v := range nums {
		if i%2 == 0 {
			totalEven += v
		} else {
			totalOdd += v
		}
	}
	count := 0
	leftEven, leftOdd := 0, 0
	for i, v := range nums {
		if i%2 == 0 {
			totalEven -= v
		} else {
			totalOdd -= v
		}
		newEven := leftEven + totalOdd
		newOdd := leftOdd + totalEven
		if newEven == newOdd {
			count++
		}
		if i%2 == 0 {
			leftEven += v
		} else {
			leftOdd += v
		}
	}
	return count
}

// ============================================================================
// Assignment 5
// ============================================================================

// Next Greater Node In Linked List
//
// Patrones:
//   - Nested Iteration
//   - Linked List Traversal
//   - Array Construction
//
// Útil cuando:
//   - se necesita comparar cada elemento con los siguientes
//   - la estructura es secuencial (linked list, array)
//   - la solución directa es suficientemente eficiente
//
// Nota:
//   - Existe solución O(n) con Monotonic Stack
//
// Complejidad:
//   - Tiempo: O(n^2)
//   - Espacio: O(n)
func NextLargerNodes(head *ListNode) []int {
	n := 0
	for current := head; current != nil; current = current.Next {
		n++
	}
	nodeVals := make([]int, n)
	idx := 0
	for current := head; current != nil; current = current.Next {
		for next := current.Next; next != nil; next = next.Next {
			if next.Val > current.Val {
				nodeVals[idx] = next.Val
				break
			}
		}
		idx++
	}
	return nodeVals
}

// Continuous Subarray Sum
//
// Patrones:
//   - Prefix Sum
//   - Hash Table
//
// Útil cuando:
//   - se busca subarray cuya suma es divisible por k
//   - se requiere longitud mínima
//   - dos prefix sums con mismo mod implican subarray divisible
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(min(n,k))
func CheckSubarraySum(nums []int, k int) bool {
	mods := make(map[int]int)
	mods[0] = -1
	prefixSum := 0
	for i, v := range nums {
		prefixSum += v
		mod := prefixSum % k
		if idx, ok := mods[mod]; ok {
			if i-idx >= 2 {
				return true
			}
		} else {
			mods[mod] = i
		}
	}
	return false
}
