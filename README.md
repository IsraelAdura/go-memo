# go-memo

A simple go package to memoize go functions with a cache key for a specified time.

## Installation

Install [go](https://golang.org/doc/install)

clone repo

```go install```

## Usage

```
import (
	memoize "github.com/israeladura/go-memo"
)

func main() {
	memoize.Memoize("cache-key", func() interface{} {
		result := callthis("foo")
		return result
	}, 60)

	memoize.Memoize("cache-key2", func() interface{} {
		result := callthis("bar")
		return result
	}, 20)

	result := memoize.Get("cache-key")
	result2 := memoize.Get("cache-key2")
    
	fmt.Println(result)
	fmt.Println(result2)

    memoize.UnMemoize("cache-key")
    fmt.Println(result)
    
    memoize.clearCache()
    fmt.Println(result2)
}

func callthis(x string) string {
	return x
}

```

## Contributing
Pull requests are welcome. 

### TODO
* Tests
* Add more features.
* Improve docs
* Add example directory.

## License
[MIT](https://choosealicense.com/licenses/mit/)
