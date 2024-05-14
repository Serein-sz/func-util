package util

func Filter[E any](slice []E, predict func(E) bool) (filterResult []E) {
	for _, v := range slice {
		if predict(v) {
			filterResult = append(filterResult, v)
		}
	}
	return
}

func Map[E any, R any](slice []E, f func(E) R) (mapResult []R) {
	for _, v := range slice {
		mapResult = append(mapResult, f(v))
	}
	return
}
