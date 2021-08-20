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
// => [0xc0002629a8]
```


