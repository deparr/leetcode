func detectCapitalUse(word string) bool {

    upperCount, lowerCount := 0, 0
    for _, char := range word {
        if char < 0x5b {
            upperCount++
        } else {
            lowerCount++
        }
    }

    if upperCount == len(word) { return true }
    if lowerCount == len(word) { return true }
    if word[0] < 0x5b && upperCount == 1 { return true }

    return false
}
