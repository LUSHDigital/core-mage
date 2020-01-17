package targets

// goget will get and/or upgrade given go packages
func goget(libs ...string) error {
	return Exec(GoBin, append([]string{"get", "-u"}, libs...)...)
}
