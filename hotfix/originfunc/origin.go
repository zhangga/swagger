package originfunc

import "fmt"

//go:noinline
func GetSpecialName(i int) string {
	name := genName(i)
	name = "origin: " + name
	return name
}

//go:noinline
func genName(i int) string {
	return fmt.Sprintf("hello_origin_%d", i)
}
