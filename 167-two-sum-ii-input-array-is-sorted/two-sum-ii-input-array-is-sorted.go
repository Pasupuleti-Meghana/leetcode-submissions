func twoSum(numbers []int, target int) []int {
    length := len(numbers)
    left := 0
    right := length - 1
    
    for left < right {
        sum := numbers[left] + numbers[right]
        if sum == target {
            return []int{left+1, right+1}
        } else if sum > target {
            right--
        } else {
            left++
        }
    }
    return []int{}
}