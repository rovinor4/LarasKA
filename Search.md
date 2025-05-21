# Dokumentasi Fungsi Pencarian

Dokumentasi ini menjelaskan cara menggunakan empat fungsi generik untuk pencarian data pada slice di Go: `FindMany`, `BinaryFindMany`, `FindOne`, dan `BinaryFindOne`.

---

## 1. FindMany

**Signature**
```go
func FindMany[T any](data []T, target T, cmp func(a, b T) int) []T
```

**Deskripsi**  
Mencari semua elemen dalam `data` yang cocok dengan `target` menggunakan fungsi perbandingan `cmp`. Mengembalikan slice berisi semua elemen yang ditemukan (bisa kosong).

**Parameter**
- `data []T` : Slice data yang akan dicari.  
- `target T` : Nilai yang dicari.  
- `cmp func(a, b T) int` : Fungsi perbandingan yang mengembalikan:
  - `< 0` jika `a < b`
  - `0` jika `a == b`
  - `> 0` jika `a > b`

**Return**
- `[]T` : Slice elemen yang cocok dengan `target`.

**Contoh Penggunaan**
```go
numbers := []int{1, 2, 3, 2, 4, 2}
cmp := func(a, b int) int { return a - b }
matches := FindMany(numbers, 2, cmp)
// matches == []int{2, 2, 2}
```

---

## 2. BinaryFindMany

**Signature**
```go
func BinaryFindMany[T any](data []T, target T, cmp func(a, b T) int) []T
```

**Deskripsi**  
Mencari semua elemen pada slice terurut `data` yang cocok dengan `target` menggunakan algoritma binary search. Mengembalikan slice berisi semua elemen yang ditemukan.

**Parameter**
- `data []T` : Slice data **terurut**.  
- `target T` : Nilai yang dicari.  
- `cmp func(a, b T) int` : Fungsi perbandingan seperti pada `FindMany`.

**Return**
- `[]T` : Slice elemen yang cocok dengan `target`.

**Contoh Penggunaan**
```go
numbers := []int{1, 2, 2, 2, 3, 4}
cmp := func(a, b int) int { return a - b }
matches := BinaryFindMany(numbers, 2, cmp)
// matches == []int{2, 2, 2}
```

---

## 3. FindOne

**Signature**
```go
func FindOne[T any](data []T, target T, cmp func(a, b T) int) (T, bool)
```

**Deskripsi**  
Mencari **satu** elemen pada `data` yang cocok dengan `target`. Mengembalikan elemen pertama yang ditemukan dan `true`; jika tidak ditemukan, mengembalikan nilai nol (`zero value`) dan `false`.

**Parameter**
- `data []T` : Slice data.  
- `target T` : Nilai yang dicari.  
- `cmp func(a, b T) int` : Fungsi perbandingan.

**Return**
- `T` : Elemen pertama yang cocok atau nilai nol.  
- `bool` : `true` jika ditemukan, `false` jika tidak.

**Contoh Penggunaan**
```go
names := []string{"Alice", "Bob", "Charlie"}
cmp := func(a, b string) int { return strings.Compare(a, b) }
name, found := FindOne(names, "Bob", cmp)
if found {
    // name == "Bob"
}
```

---

## 4. BinaryFindOne

**Signature**
```go
func BinaryFindOne[T any](data []T, target T, cmp func(a, b T) int) (T, bool)
```

**Deskripsi**  
Mencari **satu** elemen pada slice terurut `data` menggunakan binary search. Mengembalikan elemen yang ditemukan dan `true`; jika tidak, nilai nol dan `false`.

**Parameter**
- `data []T` : Slice data **terurut**.  
- `target T` : Nilai yang dicari.  
- `cmp func(a, b T) int` : Fungsi perbandingan.

**Return**
- `T` : Elemen yang cocok atau nilai nol.  
- `bool` : Status ditemukan.

**Contoh Penggunaan**
```go
numbers := []int{1, 2, 3, 4, 5}
cmp := func(a, b int) int { return a - b }
value, ok := BinaryFindOne(numbers, 3, cmp)
if ok {
    // value == 3
}
```

---

## 5. Contoh Penggunaan dengan Struct

**Definisi Struct dan Data**
```go
type Person struct {
    ID   int
    Name string
}

people := []Person{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 2, Name: "Bobby"},
    {ID: 3, Name: "Charlie"},
}
```

**Comparator Berdasarkan ID**
```go
cmpByID := func(a, b Person) int {
    return a.ID - b.ID
}
```

**FindMany**
```go
matches := FindMany(people, Person{ID: 2}, cmpByID)
// matches == []Person{{ID: 2, Name: "Bob"}, {ID: 2, Name: "Bobby"}}
```

**BinaryFindMany**  
*Pastikan `people` terurut berdasarkan `ID` sebelum digunakan.*
```go
sortedPeople := []Person{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 2, Name: "Bobby"},
    {ID: 3, Name: "Charlie"},
}
matchesBinary := BinaryFindMany(sortedPeople, Person{ID: 2}, cmpByID)
// matchesBinary == []Person{{ID: 2, Name: "Bob"}, {ID: 2, Name: "Bobby"}}
```

**FindOne**
```go
person, found := FindOne(people, Person{ID: 3}, cmpByID)
if found {
    // person == Person{ID: 3, Name: "Charlie"}
}
```

**BinaryFindOne**  
*Pastikan `sortedPeople` terurut berdasarkan `ID`.*
```go
personBinary, ok := BinaryFindOne(sortedPeople, Person{ID: 1}, cmpByID)
if ok {
    // personBinary == Person{ID: 1, Name: "Alice"}
}
```

---

## 5. Contoh Penggunaan dengan Struct

**Definisi Struct dan Data**
```go
type Person struct {
    ID   int
    Name string
}

people := []Person{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 2, Name: "Bobby"},
    {ID: 3, Name: "Charlie"},
}
```

**Comparator Berdasarkan ID**
```go
cmpByID := func(a, b Person) int {
    return a.ID - b.ID
}
```

**FindMany**
```go
matches := FindMany(people, Person{ID: 2}, cmpByID)
// matches == []Person{{ID: 2, Name: "Bob"}, {ID: 2, Name: "Bobby"}}
```

**BinaryFindMany**  
*Pastikan `people` terurut berdasarkan `ID` sebelum digunakan.*
```go
sortedPeople := []Person{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 2, Name: "Bobby"},
    {ID: 3, Name: "Charlie"},
}
matchesBinary := BinaryFindMany(sortedPeople, Person{ID: 2}, cmpByID)
// matchesBinary == []Person{{ID: 2, Name: "Bob"}, {ID: 2, Name: "Bobby"}}
```

**FindOne**
```go
person, found := FindOne(people, Person{ID: 3}, cmpByID)
if found {
    // person == Person{ID: 3, Name: "Charlie"}
}
```

**BinaryFindOne**  
*Pastikan `sortedPeople` terurut berdasarkan `ID`.*
```go
personBinary, ok := BinaryFindOne(sortedPeople, Person{ID: 1}, cmpByID)
if ok {
    // personBinary == Person{ID: 1, Name: "Alice"}
}
```