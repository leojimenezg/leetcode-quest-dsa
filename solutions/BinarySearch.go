// ============================================================================
// Binary Seach
// ============================================================================

package solutions

// Peak Index in Mountain Array
// Patrones:
//   - Binary Search
//
// Útil cuando:
//   - se busca un punto de inflexión en una secuencia unimodal
//   - se pueden descartar mitades basándose en la pendiente local
//   - se requiere O(log n) en lugar de O(n)
//
// Complejidad:
//   - Tiempo: O(log n)
//   - Espacio: O(1)
func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr)
	for (right - left) > 1 {
		mid := (left+right)/2 - 1
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid + 1
		}
	}
	return left
}

// Binary Search
// Patrones:
//   - Binary Search
//
// Útil cuando:
//   - se busca un elemento específico en un arreglo ordenado
//   - se pueden descartar mitades basándose en la comparación con el target
//   - se requiere O(log n) en lugar de O(n)
//
// Complejidad:
//   - Tiempo: O(log n)
//   - Espacio: O(1)
func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	for (right - left) > 1 {
		mid := (left + right) / 2 - 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}
