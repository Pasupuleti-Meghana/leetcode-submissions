func commonFactors(a int, b int) int {
    min := 0

    if a<b {
        min = a
    } else {
        min = b
    }

    c:=0

    for i:=1; i<=min; i++ {
        if a%i==0 && b%i == 0 {
            c++
        }
    }

    return c

}