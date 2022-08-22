def power_of_four(n: int) -> bool:
    if n <= 0:
        return False
    while n != 1:
        print(n)
        if (n % 4) != 0:
            return False
        n = n // 4
    return True

print(power_of_four(16))