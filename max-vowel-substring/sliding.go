func maxVowels(s string, k int) int {
    max := 0
    for i := 0; i < k; i++ {
        max += isVowel(s[i])
    }

    total := max
    for i := k; i < len(s); i++ {
        total += isVowel(s[i])
        total -= isVowel(s[i-k])

        if total >= max {
            if total == k {
                return k
            }
            max = total
        }

    }

    return max 
}

func isVowel(b byte) int {
    if b == 0x61 || b == 0x65 || b == 0x69 || b == 0x6f || b == 0x75 {
        return 1
    }

    return 0
}

