func containsDuplicate(nums []int) bool {
    result := false

    mapHash := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        if _, key := mapHash[nums[i]]; key {
            result = true
            break
        } else {
            mapHash[nums[i]] = i
        }
    }
    
    return result
}