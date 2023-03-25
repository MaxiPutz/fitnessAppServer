package fn

func Map[V any, R any](arr []V, fn func(ele V) R) []R {

	resultArr := []R{}
	for _, ele := range arr {
		resultArr = append(resultArr, fn(ele))
	}

	return resultArr
}

func Filter[V any](arr []V, fn func(ele V) bool) []V {
	resultArr := []V{}
	for _, ele := range arr {
		if fn(ele) {
			resultArr = append(resultArr, ele)
		}
	}

	return resultArr
}
