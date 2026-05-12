func appendCharacters(s string, t string) int {
    // Convert both strings into character arrays for easier indexing
    arrS, lenS := stringToArray(s)
    arrT, lenT := stringToArray(t)
    
    // Two pointers: pS walks through s, pT walks through t
    pS := 0
    pT := 0
    
    // Counts how many characters of t we successfully matched as a subsequence of s
    match := 0

    // Keep going until we finish checking every character of t
    for pT < lenT {

        // Still have characters left in s to compare
        if pS < lenS {
            if arrT[pT] == arrS[pS] {
                // Match found: consume one char from both s and t
                match++
                pT++
                pS++
                continue
            } else {
                // No match: only advance pS, keep looking for current t[pT] later in s
                pS++
            }
        } else {
            // s is exhausted, no more matches possible; remaining t chars must be appended
            break
        }
       
    }

    // Unmatched chars of t are exactly what we need to append at the end of s
    return lenT - match
}

func stringToArray(letters string) ([]string, int) {
    // Split string into a slice of single-character strings, also return its length
    temp := strings.Split(letters, "")
    length := len(temp)
    return temp, length
}