func pushDominoes(dominoes string) string {
    var (left byte; right byte)
    bytes := []byte(dominoes)
    done := false
    for ! done {
        done = true
        left = ' '
        for i, d := range bytes {
            if i < len(bytes) - 1 {
                right = bytes[i+1]
            }
            _new := d
            if d == '.' {
                if left == 'R' && right != 'L' {
                    _new = 'R'
                    done = false
                } else if right == 'L' && left != 'R' {
                    _new = 'L'
                    done = false
                }
            }
            left = d
            bytes[i] = _new
        }
    }
    return string(bytes)
}
