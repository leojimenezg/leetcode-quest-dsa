// ============================================================================
// Divide and Conquer
// ============================================================================

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

// Construct Binary Tree from Inorder and Postorder Traversal
// Patrones:
//   - Divide and Conquer
//
// Útil cuando:
//   - se reconstruye un árbol a partir de dos traversals complementarios
//   - el último elemento del postorder identifica la raíz del subárbol actual
//   - la posición de la raíz en inorder delimita los subárboles izquierdo y derecho
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func buildTree(inorder []int, postorder []int) *TreeNode {
	nodeIdxs := make(map[int]int)
	for i, v := range inorder {
		nodeIdxs[v] = i
	}
	var auxiliar func(inStart, inEnd, poEnd int) *TreeNode
	auxiliar = func(inStart, inEnd, poEnd int) *TreeNode {
		if inEnd < inStart {
			return nil
		}
		rootIdx := nodeIdxs[postorder[poEnd]]
		rightSize := inEnd - rootIdx
		leftNode := auxiliar(inStart, rootIdx-1, poEnd-1-rightSize)
		rightNode := auxiliar(rootIdx+1, inEnd, poEnd-1)
		return &TreeNode{Val: postorder[poEnd], Left: leftNode, Right: rightNode}
	}
	return auxiliar(0, len(inorder)-1, len(postorder)-1)
}

// Reverse Pairs
// Patrones:
//   - Merge Sort
//   - Divide and Conquer
//
// Útil cuando:
//   - se cuentan pares con condición entre elementos de distintas mitades
//   - el ordenamiento parcial permite descartar comparaciones redundantes
//   - una solución O(n^2) daría TLE y se necesita O(n log n)
//
// Complejidad:
//   - Tiempo: O(n log n)
//   - Espacio: O(n)
func reversePairs(nums []int) int {
	count := 0
	var mergeSort func(arr []int) []int
	mergeSort = func(arr []int) []int {
		n := len(arr)
		if n == 1 {
			return arr
		}
		leftArr := make([]int, n/2)
		copy(leftArr, arr[:n/2])
		rightArr := make([]int, n-n/2)
		copy(rightArr, arr[n/2:])
		left := mergeSort(leftArr)
		right := mergeSort(rightArr)
		for i, j := 0, 0; i < len(left) && j < len(right); {
			if left[i] > 2*right[j] {
				count = count + (len(left) - i)
				j++
			} else {
				i++
			}
		}
		arr = arr[:0]
		i, j := 0, 0
		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				arr = append(arr, left[i])
				i++
			} else {
				arr = append(arr, right[j])
				j++
			}
		}
		for ; i < len(left); i++ {
			arr = append(arr, left[i])
		}
		for ; j < len(right); j++ {
			arr = append(arr, right[j])
		}
		return arr
	}
	mergeSort(nums)
	return count
}
