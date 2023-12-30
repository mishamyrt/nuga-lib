// Package slices contains utils for slice manipulation
package slices

// Cast numbers collection type
func Cast[F, T number](collection []F) []T {
	result := make([]T, len(collection))
	for i := range collection {
		result[i] = T(collection[i])
	}
	return result
}
