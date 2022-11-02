# gp-cb

`go-cb` is just a library to help you to use the [sony/gobreaker](https://github.com/sony/gobreaker) without need
deal with type conversions thus making your code safer.

## Example

Using `gobreaker` your code would look like:

```go
var circuitbreaker *breaker.CircuitBreaker

func Get(url string) ([]byte, error) {
	body, err := circuitbreaker.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}
```

Using this library:

```go
var circuitbreaker *breaker.CircuitBreaker

func Get(url string) ([]byte, error) {
	body, err := cb.Execute(circuitbreaker, func() ([]byte, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}
	return body, nil
}
```

## But ... WHY!?!? 😫😫😫

> Well, it didn't reduced that much my code, why should I bother?

Yes, you are totally right. Using this light wrapper won't shorten your code enough to justify its usage. However, it will
give you more safety since it will ensure you will always return the same type you are expecting. Since on the original
implementation uses `interface{}` you can easily return a type that is not the one you are expecting. Usually it happens
when you should be returning a pointer instead of a value, or vice-versa.

## License

MIT