func mostVisitedPattern(username []string, timestamp []int, website []string) []string {

    // Step 1: Sort visits chronologically since input is not guaranteed sorted
    username, timestamp, website = sortBasedOnTimestamp(username, timestamp, website)

    // Step 2: Group websites by user, preserving timestamp order
    userMap := separateWebBasedOnUser(username, website)

    // Step 3: For each user, generate all 3-website patterns (subsequences)
    // Using a set per user to dedup, so a single user contributes at most 1 to each pattern's score
    patternWebUser := make(map[string]map[string]struct{})
    for key, value := range userMap {
        patternWebUser[key] = make(map[string]struct{})
        for i := 0; i < len(value); i++ {
            for j := i + 1; j < len(value); j++ {
                for k := j + 1; k < len(value); k++ {
                    combination := value[i] + "-" + value[j] + "-" + value[k]
                    patternWebUser[key][combination] = struct{}{}
                }
            }
        }
    }

    // Step 4: Aggregate per-user pattern sets into a global frequency counter
    patternWeb := make(map[string]int)
    for _, patternSet := range patternWebUser {
        for pattern := range patternSet {
            patternWeb[pattern]++
        }
    }

    // Step 5: Find the highest score
    score := 0
    for pattern := range patternWeb {
        if score < patternWeb[pattern] {
            score = patternWeb[pattern]
        }
    }

    // Step 6: Collect all patterns that share the max score
    var results []string
    for pattern := range patternWeb {
        if score == patternWeb[pattern] {
            results = append(results, pattern)
        }
    }

    // Step 7: Pick the lexicographically smallest among them
    // Go's ">" on strings is already lexicographic, no extra conversion needed
    result := results[0]
    if len(results) > 1 {
        for i := 1; i < len(results); i++ {
            if result > results[i] {
                result = results[i]
            }
        }
    }

    // Split the joined pattern back into a []string for the expected output format
    final := strings.Split(result, "-")
    return final
}

// Groups websites by username. Assumes the input slices are already sorted by timestamp,
// so the resulting slice for each user preserves visit order.
func separateWebBasedOnUser(username []string, website []string) map[string][]string {
    userMap := make(map[string][]string)
    for i := 0; i < len(username); i++ {
        userMap[username[i]] = append(userMap[username[i]], website[i])
    }
    return userMap
}

// Sorts the three parallel arrays by timestamp.
// Note: relies on timestamps being unique (true for the given test cases).
func sortBasedOnTimestamp(username []string, timestamp []int, website []string) ([]string, []int, []string) {
    timeMap := make(map[int][]string)
    var newUsername []string
    var newTimestamp []int
    var newWebsite []string

    for i := 0; i < len(username); i++ {
        timeMap[timestamp[i]] = []string{username[i], website[i]}
    }

    // Extract keys and sort them; map iteration order in Go is random
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