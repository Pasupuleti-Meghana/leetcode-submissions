import "math"
func findMin(nums []int) int {
    left := 0
    right := len(nums) - 1 
    res := math.MaxInt

    for left <= right {
        mid := (left + right) / 2

        if nums[left] <= nums[mid] {
            if nums[left] < res {
                res = nums[left]
            }
            left = mid + 1
        } else {
            if nums[mid] < res {
                res = nums[mid]
            }
            right = mid - 1
        }
    }

    return res 

}