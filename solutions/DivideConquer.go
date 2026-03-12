package solutions

// Beautiful Array
// Patrones:
//   - Divide and Conquer
//
// Útil cuando:
//   - una propiedad se puede garantizar separando en dos mitades
//   - la propiedad se preserva bajo transformaciones lineales
//   - el problema se puede resolver recursivamente con subproblemas del mismo tipo
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(n log n)
func beautifulArray(n int) []int {
	if n == 1 {
		return []int{1}
	}
	odd := beautifulArray((n + 1) / 2)
	even := beautifulArray(n / 2)
	for i, v := range odd {
		odd[i] = 2*v - 1
	}
	for i, v := range even {
		even[i] = 2 * v
	}
	nums := make([]int, 0, n)
	nums = append(nums, odd...)
	nums = append(nums, even...)
	return nums
}
