package to

// List2Map list convert to map
func List2Map[T any, Filed comparable](list []T, callback func(T) Filed) (map[Filed]T, error) {
	m := make(map[Filed]T)
	for _, t := range list {
		field := callback(t)
		m[field] = t
	}
	return m, nil
}
