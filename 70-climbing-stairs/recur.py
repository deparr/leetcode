def stairsR(n: int) -> int:
    if n == 0:
        return 0
    elif n == 1:
        return 1
    elif n == 2:
        return 2

    return stairsR(n - 1) + stairsR(n - 2)
