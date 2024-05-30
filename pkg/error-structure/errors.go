package errorstructure

import "fmt"

type XErrors map[string]XError

// Get one error by key
func (e XErrors) Error() string {
	output := ""
	for key, err := range e {
		output += fmt.Sprintf("key: %v, code: %v; message: %v; expected value: %v; given value: %v \n", key, err.Code, err.Message, err.ExpectedVal, err.GivenVal)
	}
	return output
}

// Get the first error by key
func (e XErrors) GetFirst() error {
	for _, err := range e {
		return err
	}
	return nil
}

// Check if a key exists
func (e XErrors) KeyExists(key string) bool {
	if _, ok := e[key]; ok {
		return true
	}
	return false
}

// Import list from other XErrors
func (e XErrors) Import(src XErrors) {
	for idx, val := range src {
		e[idx] = val
	}
}
