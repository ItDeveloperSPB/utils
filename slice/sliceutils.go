package slice

// RemoveAllInPlace removes all occurrences of a given value from a slice in place
func RemoveAllInPlace(nums []int, value int) []int {
    newLength := 0
    for _, v := range nums {
        if v != value {
            nums[newLength] = v
            newLength++
        }
    }
    return nums[:newLength]
}
