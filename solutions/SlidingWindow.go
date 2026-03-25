package solutions

// Contains Duplicate II
// Patrones:
//   - Hash Table
//
// Útil cuando:
//   - se busca un duplicado dentro de una distancia máxima de índices
//   - se requiere acceso O(1) al último índice visto de cada valor
//   - alternativa: sliding window mantiene un conjunto de tamaño k donde
//     la distancia está implícita en el tamaño de la ventana
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func containsNearbyDuplicate(nums []int, k int) bool {
	lastIndexes := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := lastIndexes[v]; ok {
			if abs(i-j) <= k {
				return true
			}
		}
		lastIndexes[v] = i
	}
	return false
}
