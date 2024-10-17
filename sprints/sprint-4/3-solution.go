func twoSum(nums []int, target int) []int {
    var result []int 

    numsHash := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        if _, key := numsHash[nums[i]]; key {
            result = append(result, numsHash[nums[i]], i)        
            } else {
            numsHash[target - nums[i]] = i
        }
    }

    return result
}