func twoSum(numbers []int, target int) []int {
    mp := make(map[int]int)

    for i,num := range numbers {
        needed := target - num
        if ind, found := mp[needed]; found {
            return []int{ind+1,i+1}
        }
        mp[num] = i
    }
    return []int{}
}