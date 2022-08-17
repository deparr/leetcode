def stairs(n: int) -> int:
    cache = [0, 1, 2]
    for i in range(3, n+1):
        cache.append(cache[i - 1] + cache[i - 2])
    print(cache)
    return cache[n]

print(stairs(4))
print(stairs(10))
print(stairs(20))
# 1x4 -- 1 1 2 -- 2 1 1 -- 1 2 1 -- 2 2