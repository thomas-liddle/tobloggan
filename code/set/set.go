package set

type Set[T comparable] map[T]struct{}

func New[T comparable](all ...T) Set[T] {
	result := make(Set[T])
	for _, t := range all {
		result.Add(t)
	}
	return result
}
func (this Set[T]) Add(values ...T) {
	for _, value := range values {
		this[value] = struct{}{}
	}
}
func (this Set[T]) Contains(t T) bool {
	_, ok := this[t]
	return ok
}
