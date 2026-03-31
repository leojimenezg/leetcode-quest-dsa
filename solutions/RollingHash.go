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

// Longest Happy Prefix
// Patrones:
//   - Rolling Hash
//   - Polynomial Hashing
//
// Útil cuando:
//   - se busca el prefijo más largo que también es sufijo
//   - se necesita comparar substrings eficientemente en O(1)
//   - la verificación directa de strings daría O(n^2)
//
// Notas:
//   - prefixHash crece de izquierda a derecha: hash = hash*p + val(s[i])
//   - suffixHash crece de derecha a izquierda: hash = hash + val(s[j])*p^i
//   - prefixHash == suffixHash implica que el prefijo s[:i+1] es igual al sufijo s[n-i-1:]
//
// Complejidad:
//   - Tiempo: O(n)
//   - Espacio: O(1)
func longestPrefix(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	p, mod, power := 31, 1_000_000_007, 1
	prefixHash, suffixHash, longestIdx := 0, 0, -1
	for i, j := 0, n-1; i < n-1 && j > 0; i, j = i+1, j-1 {
		prefixHash = (prefixHash*p + int(s[i])) % mod
		suffixHash = (suffixHash + int(s[j])*power) % mod
		power = (power * p) % mod
		if prefixHash == suffixHash {
			longestIdx = i
		}
	}
	if longestIdx < 0 {
		return ""
	}
	return s[:longestIdx+1]
}
