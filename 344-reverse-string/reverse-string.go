func reverseString(s []byte)  {
    reverse(s, 0, len(s)-1)
}

func reverse(s []byte,left int, right int) {
    if left > right {
        return
    }

    s[left], s[right] = s[right], s[left]

    reverse(s, left+1, right-1)
} 