func lengthOfLastWord(s string) int {

    fields := strings.Fields(s)
    return len(fields[len(fields)-1])
}
