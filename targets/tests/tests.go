package tests

import "fmt"

const deprecated = `DEPRECATED: Remove import of tests package from your Magefile since it's since been baked into the regular targets.

// mage:import test
_ "github.com/LUSHDigital/core-mage/targets/tests"
`

func init() {
	Deprecated()
}

// Deprecated prints the deprecation warning from the tests mage targets
func Deprecated() {
	fmt.Println(deprecated)
}
