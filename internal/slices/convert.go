package slices

func Cast[F, T Number](collection []F) []T {
	result := make([]T, len(collection))
	for i := range collection {
		result[i] = T(collection[i])
	}
	return result
}
