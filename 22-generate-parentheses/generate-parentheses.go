func generateParenthesis(n int) []string {
    result := []string{}

    build("", 0, 0, n, &result)

    return result
}

func build(currentString string, openUsed, closeUsed, n int, result *[]string) {
    // Base case
    if openUsed == n && closeUsed == n {
        *result = append(*result, currentString)
        return
    }

    // Try adding '('
    if openUsed < n {
        build(currentString + "(", openUsed+1, closeUsed, n, result)
    }

    // Try adding ')'
    if closeUsed < openUsed {
        build(currentString + ")", openUsed, closeUsed+1, n, result)
    }
}