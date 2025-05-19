LarasKA (Layanan Reservasi Kereta Api Sederhana)

## Dokumentasi Sorting

```go
// Ascending (menaik)
sorted := SelectionSort(data, func(a, b T) bool { return a < b })

// Descending (menurun)
sorted := SelectionSort(data, func(a, b T) bool { return a > b })

// Insertion sort sama cara pakainya:
insertion := InsertionSort(data, func(a, b T) bool { return a < b })
```

## Dokumentasi Search

```go
// LinearSearch pada slice int, mengembalikan nilai yang cocok
res1 := LinearSearch([]int{1, 3, 5, 3, 7}, 3, func(a, b int) int { return a - b })
fmt.Println(res1) // Output: [3 3]

// BinarySearch pada slice int yang sudah terurut, mengembalikan nilai yang cocok
res2 := BinarySearch([]int{1, 3, 3, 5, 7}, 3, func(a, b int) int { return a - b })
fmt.Println(res2) // Output: [3 3]

// LinearSearch pada slice struct, mengembalikan elemen struct yang cocok
type User struct {
    Name string
    Age  int
}

users := []User{
    {Name: "Andi", Age: 25},
    {Name: "Budi", Age: 20},
    {Name: "Andi", Age: 30},
}
target := User{Name: "Andi"}
res3 := LinearSearch(users, target, func(a, b User) int {
    if a.Name < b.Name {
        return -1
    } else if a.Name > b.Name {
        return 1
    }
    return 0
})
fmt.Println(res3) // Output: [{Andi 25} {Andi 30}]

// BinarySearch pada slice struct yang sudah terurut berdasarkan Name, mengembalikan elemen struct yang cocok
sortedUsers := SelectionSort(users, func(a, b User) bool { return a.Name < b.Name })
res4 := BinarySearch(sortedUsers, target, func(a, b User) int {
    if a.Name < b.Name {
        return -1
    } else if a.Name > b.Name {
        return 1
    }
    return 0
})
fmt.Println(res4) // Output: [{Andi 25} {Andi 30}]
```
