package v4tests

import "fmt"

func printOne[R any](fn func(string) (R, error), arg string) {
	result, err := fn(arg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func printMany[R any](fn func() ([]R, error)) {
	result, err := fn()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
