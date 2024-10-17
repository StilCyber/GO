func runningSum(nums []int) []int {
    var result []int
    sum := 0

    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        result = append(result, sum)
    }

    return result
}