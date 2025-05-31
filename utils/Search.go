package utils

func FindMany[T any](data []T, target T, cmp func(a, b T) int) []T {
	result := []T{}
	for _, v := range data {
		if cmp(v, target) == 0 {
			result = append(result, v)
		}
	}
	return result
}

func BinaryFindMany[T any](data []T, target T, cmp func(a, b T) int) []T {
	result := []T{}
	left, right := 0, len(data)-1
	found := false
	mid := 0

	for left <= right && !found {
		mid = (left + right) / 2
		cmpRes := cmp(data[mid], target)
		if cmpRes == 0 {
			found = true
		} else if cmpRes < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if found {
		for i := mid; i >= 0 && cmp(data[i], target) == 0; i-- {
			result = append(result, data[i])
		}

		for i := mid + 1; i < len(data) && cmp(data[i], target) == 0; i++ {
			result = append(result, data[i])
		}
	}

	return result
}

func FindOne[T any](data []T, target T, cmp func(a, b T) int) (T, bool, int) {
	for i, v := range data {
		if cmp(v, target) == 0 {
			return v, true, i
		}
	}
	var zero T
	return zero, false, -1
}

func BinaryFindOne[T any](data []T, target T, cmp func(a, b T) int) (T, bool, int) {
	left, right := 0, len(data)-1
	for left <= right {
		mid := (left + right) / 2
		cmpRes := cmp(data[mid], target)
		if cmpRes == 0 {
			return data[mid], true, mid
		} else if cmpRes < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	var zero T
	return zero, false, -1
}
