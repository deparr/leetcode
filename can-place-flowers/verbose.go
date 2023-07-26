func canPlaceFlowers(flowerbed []int, n int) bool {
    placed, i := 0, 1
    if flowerbed[0] == 0 && ((len(flowerbed) > 1 && flowerbed[1] == 0) || len(flowerbed) == 1)  {
        flowerbed[0] = 1
        placed++
    }
    for i < len(flowerbed) - 1 {
        if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
            flowerbed[i] = 1
            placed++
            i += 2
        } else {
            i++
        }
    }

    if flowerbed[len(flowerbed)-1] == 0 && len(flowerbed) > 1 && flowerbed[len(flowerbed)-2] == 0 {
        flowerbed[len(flowerbed)-1] = 1
        placed++
    }
    
    return placed >= n
}

