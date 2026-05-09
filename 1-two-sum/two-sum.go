func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        required := target - num
        if ind, found := m[required]; found{
            return []int{ind,i}
        }
        m[num]=i
    }
    return nil
}