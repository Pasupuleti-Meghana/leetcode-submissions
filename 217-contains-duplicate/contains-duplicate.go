func containsDuplicate(nums []int) bool {
    mp := make(map[int]int)
    for i,num := range nums {
        if _, found := mp[num]; found {
            return true
        }
        mp[num] = i
    }
    return false
}