package originfunc_hotfix

//go:noinline
func GetSpecialName(i int) string {
	name := genName(i)
	name = "origin: " + name
	return name
}
