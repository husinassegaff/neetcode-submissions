func hasDuplicate(nums []int) bool {
    hashValue := make(map[int]bool)

    for i := 0; i < len(nums); i++ {
        if hashValue[nums[i]] == false {
            hashValue[nums[i]] = true
        } else {
            return true
        } 
    }

    return false

}
