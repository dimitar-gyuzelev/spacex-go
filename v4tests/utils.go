package v4tests

import (
	"fmt"
	"net/http"
)

func handler(resp string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(resp)); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}

// TODO: refactor
// This function and printMany[R any]... below have identical bodies.
// The only difference is the signature -- printMany returns a slice.
// So if I change the signature of printMany, I wouldn't need this one.
func printOneNoArgs[R any](fn func() (R, error)) {
	result, err := fn()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

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
