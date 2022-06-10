## Non-struct types

You've previously seen defining struct types.
It is also possible to define non-struct types which you can use as an alias for a built-in type declaration, and you can define receiver functions on them to extend them in the same way as struct types.

```go
type Name string
func SayHello(n Name) {
  fmt.Printf("Hello %s\n", n)
}
n := Name("Fred")
SayHello(n)
// Output: Hello Fred
```

You can also define non-struct types composed of arrays and maps.

```go
type Names []string
func SayHello(n Names) {
  for _, name := range n {
    fmt.Printf("Hello %s\n", name)
  }
}
n := Names([]string{"Fred", "Bill"})
SayHello(n)
// Output:
// Hello Fred
// Hello Bill
```
