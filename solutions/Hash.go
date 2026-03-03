// ============================================================================
// Hash
// ============================================================================

package solutions

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
func twoSum(nums []int, target int) []int {
	uniques := make(map[int]int)
	for i, v := range nums {
		uniques[v] = i
	}
	for i, v := range nums {
		diff := target - v
		if idx, ok := uniques[diff]; ok && i != idx {
			return []int{i, idx}
		}
	}
	return nil
}

// Copy List with Random Pointer
//
// Patrones:
//   - Hash Table
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
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodes := make(map[*Node]*Node)
	for original := head; original != nil; original = original.Next {
		nodes[original] = &Node{Val: original.Val}
	}
	for original := head; original != nil; original = original.Next {
		cpy := nodes[original]
		cpy.Next = nodes[original.Next]
		cpy.Random = nodes[original.Random]
	}
	return nodes[head]
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
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i, v := range nums {
		if v <= 0 || v > n {
			nums[i] = n + 1
		}
	}
	for _, v := range nums {
		v = abs(v)
		if v <= n {
			if nums[v-1] > 0 {
				nums[v-1] = -nums[v-1]
			}
		}
	}
	for i, v := range nums {
		if v > 0 {
			return i + 1
		}
	}
	return n + 1
}
