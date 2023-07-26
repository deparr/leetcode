func reverseVowels(s string) string {
    rev := []byte(s)
    i, j := 0, len(rev) - 1
    for i < j {
        for ! isVowel(rev[i]) && i < j {
            i++
        }

        for ! isVowel(rev[j]) && i < j {
            j--
        }

        rev[i], rev[j] = rev[j], rev[i]
        i++
        j--
    }

    return string(rev)
}

func isVowel(b byte) bool {
    if b < 0x61 {
        b += 0x20 
    }

    return b == 0x61 || b == 0x65 || b == 0x69 || b == 0x6f || b == 0x75
}

