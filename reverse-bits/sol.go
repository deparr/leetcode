func reverseBits(num uint32) uint32 {
    var res uint32 = 0
    for i := 31; i > -1; i-- {
        res |= (num % 2) << i
        num >>= 1
    }
    return res
}
