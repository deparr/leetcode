func mergeAlternately(word1 string, word2 string) string {
    var b strings.Builder
    b.Grow(len(word1) + len(word2))
    i, cnt := 0, 0
    var char byte
    for i < len(word1) && i < len(word2) {
        if cnt % 2 == 0 {
            char = word1[i]
        } else {
            char = word2[i]
            i++
        }
        b.WriteByte(char)
        cnt++
    }

    if len(word1) > len(word2) {
        b.WriteString(word1[i:])
    } else if len(word2) > len(word1) {
        b.WriteString(word2[i:])
    }

    return b.String()
}

