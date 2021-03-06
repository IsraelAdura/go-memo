package main

import (
	"fmt"

	memoize "github.com/go-memo"
)

func main() {
	// memoize for 60secs
	foo := memoize.Memoize("cache-foo", func() interface{} {
		return foobar("foo")
	}, 60)

	// memoize for 20secs
	bar := memoize.Memoize("cache-bar", func() interface{} {
		return foobar("bar")
	}, 20)

	fmt.Println(foo) // foo
	fmt.Println(foo) // prints cached result //foo

	fmt.Println(bar) // bar
	fmt.Println(bar) // prints cached result //bar

	result := memoize.Get("cache-foo")
	fmt.Println(result) // foo

	memoize.UnMemoize("cache-foo")
	foo = memoize.Get("cache-foo")
	bar = memoize.Get("cache-bar")
	fmt.Println(foo) // nil
	fmt.Println(bar) // bar

	memoize.ClearCache()
	foo = memoize.Get("cache-foo")
	bar = memoize.Get("cache-bar")
	fmt.Println(foo) // nil
	fmt.Println(bar) // nil
}

func foobar(x string) string {
	return x
}
