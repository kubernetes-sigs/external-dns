package or

import "fmt"

// Panic panics with panicArgs if test is false. It does nothing otherwise.
func Panic(test bool, panicArgs ...any) {
	if !test {
		panic(fmt.Sprint(panicArgs...))
	}
}
