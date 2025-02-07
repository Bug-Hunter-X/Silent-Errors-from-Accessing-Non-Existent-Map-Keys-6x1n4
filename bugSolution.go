```go
func main() {
    var m = map[string]int{"a": 1, "b": 2}
    value, ok := m["c"]
    if ok {
        fmt.Println("Value found:", value)
    } else {
        fmt.Println("Key not found")
    }
}
```