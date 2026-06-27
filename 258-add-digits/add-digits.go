func addDigits(num int) int {
    return add(num)
}
func add(x int) int {
    sum := 0

    if x/10 == 0 {
        return x
    }
    
    for x>0 {
        sum += x%10
        x/=10
    }
    
    if sum/10 != 0 {
        return add(sum)
    }
    
    return sum
    
}