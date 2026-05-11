func hasDuplicate(nums []int) bool {
    hashValue := make(map[int]struct{})

    for i := 0; i < len(nums); i++ {
        if _, exists := hashValue[nums[i]]; exists {
            return true
        } else {
            hashValue[nums[i]] = struct{}{} 
        } 
    }

    return false

}
