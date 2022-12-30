func numberOfWeakCharacters(properties [][]int) int {
    sort.Slice(properties, func (x, y int) bool { 
        if properties[x][0] == properties[y][0] {
            return properties[x][1] > properties[y][1]
        }
        return properties[x][0] < properties[y][0]
    })

    num_weak := 0
    max := properties[len(properties)-1][1]

    for i := len(properties)-2; i > -1; i-- {
        if properties[i][1] < max {
            num_weak++
        } else {
            max = properties[i][1]
        }
    }

    return num_weak

}
