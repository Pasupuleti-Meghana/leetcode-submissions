func searchRange(nums []int, target int) []int {
    n := len(nums)
    lb := lowerbound(nums, target)
    if len(nums) == 0 {
        return []int{-1,-1}
    }
    if lb == n || nums[lb] != target {
        return []int{-1,-1}
    }
    return []int{lb, upperbound(nums, target)-1}
}

func lowerbound(nums []int, target int) int {
    n := len(nums)
    left := 0
    right := n-1
    lb := n

    for left <= right {
        mid := (left + right) / 2
        if nums[mid] >= target {
            lb = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return lb 

}

func upperbound(nums []int, target int) int {
    n := len(nums)
    left := 0
    right := n-1
    up := n

    for (left <= right) {
        mid := (left + right) / 2
        if nums[mid] > target {
            up = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return up

} 