package main 

import "sort"

func singleNumber(nums []int) int {

    sort.Ints(nums)
    result := nums[0]

    for i := 1; i < len(nums) - 1; i += 2 {
        if nums[i] != result {
            break
        }
        result = nums[i + 1]
    }

    return result
}