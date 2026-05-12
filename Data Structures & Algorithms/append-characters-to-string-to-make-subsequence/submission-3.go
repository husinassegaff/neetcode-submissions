func appendCharacters(s string, t string) int {
    arrS, lenS := stringToArray(s)
    arrT, lenT := stringToArray(t)
    pS := 0
    pT := 0
    match := 0

    for pT < lenT {

        if pS < lenS {
            if arrT[pT] == arrS[pS] {
                match++
                pT++
                pS++
                continue
            } else {
                pS++
            }
        } else {
            break
        }
       
    }

    return lenT - match

}

func stringToArray(letters string) ([]string, int) {
    temp := strings.Split(letters, "")
    length := len(temp)
    return temp, length
}