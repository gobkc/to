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

// List2Maps list convert to maps
func List2Maps[T any, Filed comparable](list []T, callback func(T) Filed) (map[Filed][]T, error) {
	m := make(map[Filed][]T)
	for _, t := range list {
		field := callback(t)
		if q, ok := m[field]; ok {
			m[field] = append(q, t)
		} else {
			m[field] = []T{t}
		}
	}
	return m, nil
}
