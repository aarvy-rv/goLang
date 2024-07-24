package is_prime

func Isprime(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		} else {
			continue
		}
	}
	return true
}
