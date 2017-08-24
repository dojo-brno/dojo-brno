package circularprimes

import "sort"

func HowManyBelow(n int) int {
	count := 0
	for _, p := range primes {
		if p >= n {
			break
		}
		if IsCircularPrime(p) {
			count++
		}
	}
	return count
}

func IsCircularPrime(n int) bool {
	for _, r := range rotations(n) {
		if !isPrime(r) {
			return false
		}
	}
	return true
}

func isPrime(n int) bool {
	i := sort.SearchInts(primes[:], n)
	return i < len(primes) && primes[i] == n
}

func rotations(n int) []int {
	if 10 <= n && n < 100 {
		return []int{n, n%10*10 + n/10}
	}
	if 100 <= n && n < 1000 {
		o := n%100*10 + n/100
		p := o%100*10 + o/100
		return []int{n, o, p}
	}
	if 1000 <= n && n < 10000 {
		o := n%1000*10 + n/1000
		p := o%1000*10 + o/1000
		q := p%1000*10 + p/1000
		return []int{n, o, p, q}
	}
	if 10000 <= n && n < 100000 {
		o := n%10000*10 + n/10000
		p := o%10000*10 + o/10000
		q := p%10000*10 + p/10000
		r := q%10000*10 + q/10000
		return []int{n, o, p, q, r}
	}
	if 100000 <= n && n < 1000000 {
		o := n%100000*10 + n/100000
		p := o%100000*10 + o/100000
		q := p%100000*10 + p/100000
		r := q%100000*10 + q/100000
		s := r%100000*10 + r/100000
		return []int{n, o, p, q, r, s}
	}
	return []int{n}
}
