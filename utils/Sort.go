package utils

func SelectionSort[T any](arr []T, less func(a, b T) bool) []T {
    n := len(arr)
    a := make([]T, n)
    copy(a, arr)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if less(a[j], a[minIdx]) {
                minIdx = j
            }
        }
        a[i], a[minIdx] = a[minIdx], a[i]
    }
    return a
}

func InsertionSort[T any](arr []T, less func(a, b T) bool) []T {
    n := len(arr)
    a := make([]T, n)
    copy(a, arr)
    for i := 1; i < n; i++ {
        key := a[i]
        j := i - 1
        for j >= 0 && less(key, a[j]) {
            a[j+1] = a[j]
            j--
        }
        a[j+1] = key
    }
    return a
}