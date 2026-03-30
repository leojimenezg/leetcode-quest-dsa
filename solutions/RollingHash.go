package solutions

// Shortest Palindrome
// Patrones:
//   - Rolling Hash
//   - Polynomial Hashing
//
// Útil cuando:
//   - se busca el prefijo palindrómico más largo de un string
//   - se necesita comparar substrings eficientemente en O(1)
//   - la verificación carácter por carácter daría O(n^2)
//
// Notas:
//   - hash    = hash*p + val(s[i])          → lectura izquierda a derecha
//   - hashRev = hashRev + val(s[i])*p^i     → lectura derecha a izquierda
//   - se usa mod 1_000_000_007 (primo) para evitar overflow y reducir colisiones
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(n)
func shortestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	p, mod := 31, 1_000_000_007
	hash, hashRev, power := 0, 0, 1
	longestIdx := 0
	for i, v := range s {
		hash = (hash*p + int(v)) % mod
		hashRev = (hashRev + int(v)*power) % mod
		power = (power * p) % mod
		if hash == hashRev {
			longestIdx = i
		}
	}
	reversed := []rune(s[longestIdx+1:])
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return string(reversed) + s
}
