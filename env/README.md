# Env
The `env` package is used to load environment variables in the context of using the core-mage targets and development environment. It provides convenient functions for loading or overloading.

## Example

### Loading default environment

```go
func main() {
    env.LoadDefault()
}
```

### Loading specific environment files

```go
func main() {
    env.Load("infa/does-not-override.env")
    env.Overload("infa/will-override.env")
}
```

### Loading default test configuration

```go
func TestMain(m *testing.M) {
	env.MustLoadDefaultTest(m)
	os.Exit(m.Run())
}
```

### Loading specific files during the testing

```go
func TestMain(m *testing.M) {
    env.LoadTest(m, "infa/does-not-override.env")
    env.OverloadTest(m, "infa/will-override.env")
    os.Exit(m.Run())
}
```
