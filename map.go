package funcs

func Map[T []U, U any, V []W, W any](values T, f func(u U) W) V {
	ret := make(V, len(values))

	for i, value := range values {
		ret[i] = f(value)
	}
	return ret
}
