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
			if i-j <= k {
				return true
			}
		}
		lastIndexes[v] = i
	}
	return false
}

// Number of Substrings Containing All Three Characters
// Patrones:
//   - Sliding Window
//   - Array as Hash
//
// Útil cuando:
//   - se cuentan substrings que satisfacen una condición sobre caracteres
//   - el conjunto de caracteres es pequeño y conocido
//   - todos los substrings a la derecha de una ventana válida también son válidos
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func numberOfSubstrings(s string) int {
	n := len(s)
	present := [3]int{} // [a, b, c]
	left, right, count := 0, 0, 0
	for right < n {
		present[s[right]-'a']++
		for present[0] > 0 && present[1] > 0 && present[2] > 0 {
			count += n - right
			present[s[left]-'a']--
			left++
		}
		right++
	}
	return count
}
