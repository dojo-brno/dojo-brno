package circularprimes

func IsCircularPrime(n int) bool {
	if !isPrime(n) {
		return false
	}
	return true
}

func isPrime(n int) bool {
	primes := []int{2, 3, 5}
	for _, p := range primes {
		if n != p && n%p == 0 {
			return false
		}
	}
	return true
}
