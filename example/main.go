package main

func main() {
	// Because of a change in go 1.13 of `go list`, mage is broken if this file doesn't exist.
	// We need to keep a regular go file without build tags in the same directory as the magefile.
	// TODO: Remove this file after this issue is fixed: https://github.com/magefile/mage/issues/262
}
