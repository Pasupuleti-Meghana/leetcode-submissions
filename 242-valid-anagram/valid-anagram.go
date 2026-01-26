func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }

    map1 := make(map[rune]int)

    for _, ch := range s {
        map1[ch]++
    }

    for _, ch2 := range t {
        map1[ch2]--
        if map1[ch2] < 0 {
            return false
        }
    }
    return true
}