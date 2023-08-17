package functional

type ReflexiveFunc[A any] func(A) A

type Slice[A any] []A

func (s Slice[A]) Map(m func(A) A) (a []A) {
	for _, v := range a {
		a = append(a, m(v))
	}
	return
}

func Map[A any](input []A, m func(A) A) (a []A) {
	output := make([]A, len(input)) // performance considerations?
	for i, element := range input {
		output[i] = m(element)
	}
	return output
}

type TransformSlice[A, B any] []A

func (a TransformSlice[A, B]) Map(f func(A) B) (b []B) {
	for _, v := range a {
		b = append(b, f(v))
	}
	return
}

func EmbellishMap[A, B any, E Monoid[E]](a []A, f func(A) (B, E)) func([]A) ([]B, E) {
	return func(a []A) (b []B, e E) {
		for _, _a := range a {
			_b, _e := f(_a)
			e = e.Combine(_e)
			b = append(b, _b)
		}
		return
	}
}

// produces a function that can take a slice of A from a function that takes A
func (f EmbellishedTransformFunc[A, B, E]) EmbellishMap() func([]A) ([]B, E) {
	return func(a []A) (b []B, e E) {
		for _, _a := range a {
			_b, _e := f(_a)
			e = e.Combine(_e)
			b = append(b, _b)
		}
		return
	}
}

type EmbellishedTransformSlice[A, B any, E Monoid[E]] []A

// produces a function from a slice
func (a EmbellishedTransformSlice[A, B, E]) EmbellishMap(f func(A) (B, E)) func([]A) ([]B, E) {
	return func(a []A) (b []B, e E) {
		for _, _a := range a {
			_b, _e := f(_a)
			e = e.Combine(_e)
			b = append(b, _b)
		}
		return
	}
}
