// ============================================================================
// Assignment 9
// ============================================================================

package solutions

// Reorder List
//
// Patrones:
//   - Array as Index (LinkedList to Array)
//   - Two Pointers
//
// Útil cuando:
//   - se necesita acceso aleatorio a nodos de una linked list
//   - se intercalan nodos del inicio y del final
//   - convertir a array simplifica el rewiring de punteros
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	nodes := make([]*ListNode, n)
	cur := head
	for i := range n {
		nodes[i] = cur
		cur = cur.Next
	}
	i, j := 0, n-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i >= j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

// All Paths From Source to Target
//
// Patrones:
//   - BFS
//   - Graph Traversal
//
// Útil cuando:
//   - se buscan todos los caminos posibles en un DAG
//   - se exploran rutas nivel por nivel
//   - el grafo no tiene ciclos (DAG) por lo que no se necesita visited
//
// Complejidad:
//   - Tiempo: O(2^n * n) — puede haber 2^n caminos de longitud n
//   - Espacio: O(2^n * n) por el almacenamiento de todos los caminos
func allPathsSourceTarget(graph [][]int) [][]int {
	goalNode := len(graph) - 1
	combinations := make([][]int, 0)
	queue := Queue{}
	queue.Push([]int{0})
	for !queue.Empty() {
		path := queue.Pop()
		lastNode := path[len(path)-1]
		if lastNode == goalNode {
			combinations = append(combinations, path)
		} else {
			for i := range graph[lastNode] {
				newPath := make([]int, len(path)+1)
				copy(newPath, path)
				newPath[len(newPath)-1] = graph[lastNode][i]
				queue.Push(newPath)
			}
		}
	}
	return combinations
}
