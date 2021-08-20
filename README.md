# Go Generics By Example

## Slice utils
### mapSlice
```go
mapSlice([]int{1, 2, 3, 4}, func(v int) int {
	return v * 2
})
// => [2 4 6 8]
```

```go
mapSlice([]string{"foobar", "baz"}, func(v string) int {
	return len(v)
})
// => [6 3]
```

### flatMapSlice
```go
flatMapSlice([][]int{{1, 2}, {3, 4}}, func(v []int) []int {
	results := []int{}
	results = append(results, v...)
	results = append(results, 100)
	return results
})
// => [1 2 100 3 4 100]
```

### filterSlice
```go
filterSlice([]int{1, 2, 3, 4, 5, 6}, func(v int) bool {
	return v/2 == 0
})
// => [1]
```

```go
type User struct {
	ID    int
	Name  string
	Admin bool
}
users := []*User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob", Admin: true},
	{ID: 3, Name: "Carol"},
}
nonAdminUsers := filterNotSlice(
	users,
	func(user *User) bool { return user.Admin },
)
mapSlice(
	nonAdminUsers,
	func(user *User) string { return user.Name },
)
// => [Alice Carol]
```

### filterMapSlice
```go
filterMapSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(v int) (int, bool) {
	return v * 2, v%2 == 0
})
// => [4 8 12 16 20]
```

### findSlice
```go
type User struct {
	Name string
	Age  int
}
teens, _ := findSlice([]*User{{Name: "Alice", Age: 18}, {Name: "Bob", Age: 27}, {Name: "Carol", Age: 24}}, func(u *User) bool {
	return u.Age >= 20
})
teens
// => &{Name:Bob Age:27}
```

### indexSlice / lastIndexSlice
```go
indexSlice([]string{"foo", "bar", "baz"}, func(v string) bool {
	return strings.HasPrefix(v, "b")
})
// => 1
```

```go
indexSlice([]string{"foo", "bar", "baz"}, func(v string) bool {
	return strings.HasPrefix(v, "a")
})
// => -1
```

```go
lastIndexSlice([]string{"foo", "bar", "baz"}, func(v string) bool {
	return strings.HasPrefix(v, "b")
})
// => 2
```

```go
lastIndexSlice([]string{"foo", "bar", "baz"}, func(v string) bool {
	return strings.HasPrefix(v, "a")
})
// => -1
```

### containsSlice
```go
containsSlice([]int{1, 5, 6}, 4)
// => false
```

### allSlice
```go
allSlice([]string{"ant", "bear", "cat"}, func(v string) bool {
	return len(v) >= 3
})
// => true
```

```go
allSlice([]string{"ant", "bear", "cat"}, func(v string) bool {
	return len(v) >= 4
})
// => false
```

### someSlice
```go
someSlice([]string{"ant", "bear", "cat"}, func(v string) bool {
	return len(v) >= 4
})
// => true
```

```go
someSlice([]string{"ant", "bear", "cat"}, func(v string) bool {
	return len(v) >= 5
})
// => false
```

### maxSlice / minSlice
```go
maxSlice([]string{"albatross", "dog", "horse"}, func(v string) int {
	return len(v)
})
// => albatross
```

```go
minSlice([]string{"albatross", "dog", "horse"}, func(v string) int {
	return len(v)
})
// => dog
```

### compactSlice
```go
compactSlice([]string{"foo", "bar", "", "baz", ""})
// => [foo bar baz]
```

```go
type User struct {
	ID   int
	Name string
}
mapSlice(
	compactSlice([]User{{ID: 1, Name: "Alice"}, {}, {ID: 2, Name: "Bob"}}),
	func(user User) string { return user.Name },
)
// => [Alice Bob]
```

### uniqSlice
```go
uniqSlice([]string{"a", "a", "b", "b", "c"})
// => [a b c]
```

### reduceSlice
```go
reduceSlice([]int{1, 2, 3, 4, 5}, func(n int, acc int) int {
	return acc + n
}, 0)
// => 15
```

### groupSliceBy
```go
groupSliceBy([]int{1, 2, 3, 4, 5, 6}, func(v int) int { return v % 3 })
// => map[0:[3 6] 1:[1 4] 2:[2 5]]
```

### reverseSlice
```go
reverseSlice([]int{1, 2, 3, 4, 5})
// => [5 4 3 2 1]
```

### differenceSlice
```go
differenceSlice([]int{1, 1, 2, 2, 3, 3, 4, 5}, []int{1, 2, 4})
// => [3 3 5]
```

### intersectionSlice
```go
intersectionSlice([]int{1, 2, 3}, []int{0, 1, 2})
// => [1 2]
```


## Map utils
### mapKeys / mapValues
```go
mapKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3})
// => [baz foo bar]
```

```go
mapValues(map[string]int{"foo": 1, "bar": 2, "baz": 3})
// => [1 2 3]
```

### mapTransformKeys / mapTransformValues
```go
transformMapKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string) string {
	return k + k
})
// => map[barbar:2 bazbaz:3 foofoo:1]
```

```go
transformMapValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(v int) int {
	return v * 2
})
// => map[bar:4 baz:6 foo:2]
```


## Set
### setFromSlice
```go
set := setFromSlice([]int{1, 2, 3})
[]bool{
	set.Contains(2),
	set.Contains(4),
}
// => [true false]
```

### makeSet
```go
set := makeSet[int](3)
set.Add(1)
set.Add(2)
set.Add(3)
[]bool{
	set.Contains(2),
	set.Contains(4),
}
// => [true false]
```


## Sqlx utils
### SqlxSelect
```go
type User struct {
	ID   uint64
	Name string
}
ctx := context.Background()
users, err := sqlxSelect[*User](ctx, db, "SELECT * FROM users")
if err != nil {
	panic(err)
}
users
// => [0xc00048d308]
```


