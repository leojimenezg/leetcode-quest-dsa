// ============================================================================
// LeetCode Quest: Data Structures and Algorithms
// ============================================================================

package problems

import "strings"

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
