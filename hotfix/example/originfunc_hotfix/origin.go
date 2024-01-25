package originfunc_hotfix

import "fmt"

//go:noinline
func GetSpecialName(i int) string {
	name := genName(i)
	name = "hotfix: " + name
	return name
}

//go:noinline
func genName(i int) string {
	return fmt.Sprintf("hello_hotfix_%d", i)
}
