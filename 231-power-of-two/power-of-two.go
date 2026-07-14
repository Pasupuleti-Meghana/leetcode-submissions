func isPowerOfTwo(n int) bool {
    res := Power(n)
    return res

}

func Power(n int) bool {

    if n == 1 {
        return true
    }

    if n <= 0 {
        return false
    }

    if n%2 == 0 {
        return Power(n/2)
    }

    return false
    
}