package util

func Divisors(n int) []int {
	if n == 1 {
		return []int{1}
	}
	d := []int{1, n}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			d = append(d, i)
		}
	}
	return d
}

func Digits(n int) []int {
	d := []int{}
	for n > 0 {
		d = append(d, n%10)
		n /= 10
	}
	for i := 0; i < len(d)/2; i++ {
		d[i], d[len(d)-i-1] = d[len(d)-i-1], d[i]
	}
	return d
}
