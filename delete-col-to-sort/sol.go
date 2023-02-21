func minDeletionSize(strs []string) int {
    wordLen := len(strs[0])
    colToDelete := 0
    for i := 0; i < wordLen; i++ {
        for j := 0; j < len(strs) - 1; j++ {
            if strs[j][i] > strs[j+1][i] {
                colToDelete++
                break
            }
        }
    }

    return colToDelete
}
