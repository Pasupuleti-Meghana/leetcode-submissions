func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i,num := range nums{
        needed := target - num
        if ind,found := m[needed]; found {
            return []int{ind,i}
        }
        m[num]=i
    }
    return []int{}
}