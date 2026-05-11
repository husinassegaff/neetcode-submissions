func hasDuplicate(nums []int) bool {
    hashValue := make(map[int]bool)
    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        if hashValue[nums[i]] == false {
            hashValue[nums[i]] = true
        } else {
            return true
        } 
    }

    return false

}
