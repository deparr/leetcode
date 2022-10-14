// TIMES OUT - SOLUTION NEEDS TO BE LOG(N)

func myPow(x float64, n int) float64 {
    var fixed_x float64
    if n == 0 {
        return 1.0
    }

    if n < 0 {
        fixed_x = 1 / x
        n *= -1
    } else {
      fixed_x = x
    }
	num := fixed_x
	for i := 0; i < n-1; i++ {
		num *= fixed_x
	}
	return num
}

