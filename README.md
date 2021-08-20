# Go Generics By Example
## Slice utils
### mapSlice
```go
mapSlice([]int{1, 2, 3}, func(v int) int {
	return v * 2
})
// => [2 4 6]
```

```go
mapSlice([]string{"foobar", "baz"}, func(v string) int {
	return len(v)
})
// => [6 3]
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

### containsSlice
```go
containsSlice([]int{1, 5, 6}, 4)
// => false
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

### groupSliceBy
```go
groupSliceBy([]int{1, 2, 3, 4, 5, 6}, func(v int) int { return v % 3 })
// => map[0:[3 6] 1:[1 4] 2:[2 5]]
```

### differenceSlice
```go
differenceSlice([]int{1, 1, 2, 2, 3, 3, 4, 5}, []int{1, 2, 4})
// => [3 3 5]
```


