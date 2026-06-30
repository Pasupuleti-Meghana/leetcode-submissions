func searchInsert(nums []int, target int) int {
    res := 0 
    n := len(nums)
    left := 0
    right := n-1

    for left <= right {
        mid := (left + right) / 2

        if nums[mid] >= target {
            res = mid
            right = mid - 1
        } else {
            left = mid + 1
            res = left
        }
    }

    return res 

}