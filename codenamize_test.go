package codenam_test

import (
	"fmt"

	"github.com/azr/codenam"
)

func Example() {
	type digest string
	c := codenam.Ize(digest("xyz"))
	fmt.Println(c)

	c = codenam.Ize("abc")
	fmt.Println(c)

	cs := codenam.IzeSlice([]digest{"xyz", "abc"})
	fmt.Println(cs)

	cs = codenam.IzeSlice([]string{"abc", "def"})
	fmt.Println(cs)
	// Output: left-classroom
	// testy-give
	// [left-classroom testy-give]
	// [testy-give subsequent-win]
}
