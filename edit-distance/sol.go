func minDistance(word1 string, word2 string) int {
	arr := make([][]int, len(word2)+1)
	for i := range arr {
		arr[i] = make([]int, len(word1)+1)
	}

	for i := range arr[0] {
		arr[0][i] = i
	}

	for i := range arr {
		arr[i][0] = i
	}

	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[i]); j++ {

			if word1[j-1] == word2[i-1] {
				arr[i][j] = arr[i-1][j-1]

			} else {
				sc := arr[i-1][j] + 1

				if arr[i-1][j-1]+1 < sc {
					sc = arr[i-1][j-1] + 1
				}

				if arr[i][j-1]+1 < sc {
					sc = arr[i][j-1] + 1
				}

				arr[i][j] = sc
			}
		}
	}

	return arr[len(word2)][len(word1)]

}