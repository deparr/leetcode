func videoStitching(clips [][]int, time int) int {
    sort.Slice(clips, func(x, y int) bool {return clips[x][0] < clips[y][0]})
    if clips[0][0] != 0 {
        return -1
    }
    
    curr_end := clips[0][1]
    idx := 0
    for i, val := range clips {
        if val[0] != 0 {
            break
        }
        if val[1] > curr_end {
            curr_end = val[1]
            idx = i
        }
    }

    idx++
    numclips := 1
    for curr_end < time {
        if idx >= len(clips) || clips[idx][0] > curr_end {
            return -1
        }
        new_end := curr_end
        for idx < len(clips) && clips[idx][0] <= curr_end {
            if clips[idx][1] > new_end {
                new_end = clips[idx][1]
            }
            idx++
        }

        numclips++
        curr_end = new_end
    }
    return numclips
}
