def tribonacci(n: int) -> int:
    if n == 0:
        return 0
    elif n == 1 or n == 2:
        return 1
    cache = [0] * (n + 1)
    cache[1] = 1
    cache[2] = 1

    for i in range(3, n + 1):
        cache[i] = cache[i - 1] + cache[i - 2] + cache[i - 3]

    return cache[n]

print(tribonacci(4))
print(tribonacci(10))
print(tribonacci(25))