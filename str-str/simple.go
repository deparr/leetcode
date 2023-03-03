func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {
        return -1
    }

    for i := range haystack {
        found := true
        for j := 0; j < len(needle); j++ {
            if i+j >= len(haystack) || haystack[i+j] != needle[j] {
                found = false
                break
            }
        }
        if found {
            return i
        }
    }


    return -1
}
