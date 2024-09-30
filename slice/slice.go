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

func RemoveFirst(nums []int, value int) []int {
	for i, v := range nums {
		if v == value {
			// Удаляем элемент с индексом i
			return append(nums[:i], nums[i+1:]...)
		}
	}
	// Если элемент не найден, возвращаем оригинальный срез
	return nums
}

func RemoveAllWithNewSlice(nums []int, value int) []int {
	result := []int{}
	for _, v := range nums {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}
