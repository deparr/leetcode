func myPow(x float64, n int) float64 {
    var helper func (x float64, n int) float64
    var fixed_x float64

    if n < 0 {
        fixed_x = 1 / x
    } else {
      fixed_x = x
    }

    helper = func (x float64, n int) float64 {
        if n == 0 {
            return 1
        }
        z := helper(x, n / 2)
        if n % 2 == 0 {
            return z * z
        }
        return x * z * z
    }

    return helper(fixed_x, n)
}
