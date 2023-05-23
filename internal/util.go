package internal

import "fmt"

func IndexOf[T comparable](collection []T, entity T) (int, error) {
    for i, x := range collection {
        if x == entity {
            return i, nil
        }
    }
    return -1, fmt.Errorf("entity not found")
}

func ExcludeIndex[T any](collection []T, index int) []T {
	result := make([]T, 0, len(collection)-1)
	result = append(result, collection[:index]...)
	result = append(result, collection[index+1:]...)
	return result
}
