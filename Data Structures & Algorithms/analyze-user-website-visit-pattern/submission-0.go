

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	
    username, timestamp, website = sortBasedOnTimestamp(username, timestamp, website)

    userMap := separateWebBasedOnUser(username, website)

    patternWebUser := make(map[string]map[string]struct{})

    for key, value := range userMap {
        patternWebUser[key] = make(map[string]struct{})
        for i := 0; i < len(value); i++ {
            for j := i+1; j < len(value); j++ {
                for k := j+1; k < len(value); k++ {
                    combination := value[i] + "-" + value[j] + "-" + value[k]    
                    patternWebUser[key][combination] = struct{}{}
                }
            }
        }
    }

    patternWeb := make(map[string]int)

    for _, patternSet := range patternWebUser {
        for pattern := range patternSet {
            patternWeb[pattern]++
        }
    }

    score := 0
    for pattern := range patternWeb {
        if score < patternWeb[pattern] {
            score = patternWeb[pattern]
        }
    }

    var results []string

    for pattern := range patternWeb {
        if score == patternWeb[pattern] {
            results = append(results, pattern)
        }
    }

    result := results[0]
    if len(results) > 1 {
        for i := 1; i < len(results); i++ {
            if result > results[i] {
                result = results[i]
            }   
        }
    }

    final := strings.Split(result, "-")

    return final
}

func separateWebBasedOnUser(username []string, website []string) map[string][]string {
    
    userMap := make(map[string][]string)

    for i := 0; i < len(username); i++ {
        userMap[username[i]] = append(userMap[username[i]], website[i])
    }

    return userMap
}

func sortBasedOnTimestamp(username []string, timestamp []int, website []string) ([]string, []int, []string) {
    timeMap := make(map[int][]string)
    var newUsername []string
    var newTimestamp []int
    var newWebsite []string

    for i := 0; i < len(username); i++ {
        timeMap[timestamp[i]] = []string{username[i], website[i]}
    }


    keys := make([]int, 0, len(timeMap))

    for k := range timeMap {
        keys = append(keys, k)
    }

    sort.Ints(keys)

    for i := 0; i < len(keys); i++ {
        newUsername = append(newUsername, timeMap[keys[i]][0])
        newWebsite = append(newWebsite, timeMap[keys[i]][1])
        newTimestamp = append(newTimestamp, keys[i])
    }

    return newUsername, newTimestamp, newWebsite

}
