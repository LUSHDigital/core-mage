# Env
The `env` package is used to load environment variables in the context of using the core-mage targets and development environment. It provides convenient functions for loading or overloading.

## Example

### Loading development environment config

```go
env.TryLoadDev()
```

### Loading development environment config together with other source

```go
env.TryLoadDev("package/other.env")
```


### Loading test environment config

```go
env.TryLoadTest()
```

### Loading test environment config together with other source

```go
env.TryLoadTest("package/other.env")
```
