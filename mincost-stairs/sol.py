def minstairs(cost: list[int]) -> int:
    idx = 0 if cost[0] < cost[1] else 1
    total = cost[idx]
    while idx < len(cost):
        if idx + 2 >= len(cost) or idx + 1 >= len(cost) or idx == len(cost):
            break
        jump = 1 if cost[idx+1] < cost[idx+2] else 2
        total += cost[idx+jump]
        idx += jump

    return total


print(minstairs([10, 15, 20]))
print(minstairs([1, 100, 1, 1, 1, 100, 1, 1, 100, 1]))