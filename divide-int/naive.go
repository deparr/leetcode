func divide(dividend int, divisor int) int {
    if dividend == divisor {
        return 1
    }


    neg := false
    if (dividend >= 0) != (divisor >= 0) {
        neg = true
    }

    if dividend < 0 {
        dividend *= -1
    }
    if divisor < 0 {
        divisor *= -1
    }

    if divisor == 1 {
        if neg {
            dividend *= -1
            if dividend < -(1 << 31) {
                dividend = -(1 << 31)
            }
        }

        if dividend > ((1 << 31)-1) {
            dividend = ((1 << 31)-1)
        }
        return dividend
    }

    count := 0
    for dividend > 0 {
        dividend -= divisor
        count++
    }

    if dividend < 0 {
        count -= 1
    }

    if neg {
        count *= -1
        if count < -(1 << 31) {
            count = -(1 << 31)
        }
    } else if count > ((1 << 31)-1) {
        count = ((1 << 31)-1)
    }

    return count
}

