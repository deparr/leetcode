func compress(chars []byte) int {
    if len(chars) == 1 {
        return 1
    }
    prev := chars[0]
    count := 1
    writeIdx := 0
    for i := 1; i < len(chars); i++ {
        if chars[i] == prev {
            count++
        } else {
            writeIdx += write(chars, writeIdx, count, prev)
            prev = chars[i]
            count = 1
        }
    }
    writeIdx += write(chars, writeIdx, count, prev)

    return writeIdx
}

func write(buf []byte, i int, count int, b byte) int {
    buf[i] = b
    j := i + 1
    start := j

    if count == 1 {
        return 1
    }

    for count > 0 {
        buf[j] = byte((count % 10)) | 0x30
        j++
        count /= 10
    }

    bytes := j - i
    j--
    for start < j {
        buf[start], buf[j] = buf[j], buf[start]
        j--
        start++
    }
    return bytes
}

