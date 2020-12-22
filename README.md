```go
func main() {
	var m Map

	m.Store("a", "aaa")
	m.Range(func(key, value interface{}) bool {
		fmt.Println("1 range ", key, value)
		return true
	})
	fmt.Println("1 length: ", *m.Length())

	m.Store("b", "bbb")
	m.Range(func(key, value interface{}) bool {
		fmt.Println("2 range ", key, value)
		return true
	})
	fmt.Println("2 length: ", *m.Length())

	m.Delete("b")
	m.Range(func(key, value interface{}) bool {
		fmt.Println("3 range ", key, value)
		return true
	})
	fmt.Println("3 length: ", *m.Length())

	m.Store("b", "bbb")
	fmt.Println("4 length: ", *m.Length())
	m.Range(func(key, value interface{}) bool {
		fmt.Println("4 range ", key, value)
		return true
	})

	m.Delete("b")
	m.Range(func(key, value interface{}) bool {
		fmt.Println("4 range ", key, value)
		return true
	})
	fmt.Println("4 length: ", *m.Length())

	m.Range(func(key, value interface{}) bool {
		fmt.Println("5 range ", key, value)
		return true
	})

	fmt.Println("5 length: ", *m.Length())

	fmt.Println("done")
}
```