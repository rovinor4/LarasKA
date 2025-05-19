package controller

func LinearSearch[T any](data []T, target T, cmp func(a, b T) int) []T {
	result := []T{}
	for _, v := range data {
		if cmp(v, target) == 0 {
			result = append(result, v)
		}
	}
	return result
}

func BinarySearch[T any](data []T, target T, cmp func(a, b T) int) []T {
	result := []T{}
	left, right := 0, len(data)-1
	for left <= right {
		mid := (left + right) / 2
		cmpRes := cmp(data[mid], target)
		if cmpRes == 0 {
			for i := mid; i >= left && cmp(data[i], target) == 0; i-- {
				result = append(result, data[i])
			}
			for i := mid + 1; i <= right && cmp(data[i], target) == 0; i++ {
				result = append(result, data[i])
			}
			break
		} else if cmpRes < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}