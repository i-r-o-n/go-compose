package functional

// compression function type
type CompressFunc[A any] func(...A) A

// transformation function type
type TransformFunc[A, B any] func(A) B

func Compose[A, B, C any](f TransformFunc[A, B], g TransformFunc[B, C]) {

}

type EmbellishedTransformFunc[A, B any, E Monoid[E]] func(A) (B, E)

// unrestricted embellishment
type EmbellishedTransformFunc_[A, B, C any] func(A) (B, C)

type CombinableFunc[A any] func(...A) A

type Monoid[A any] interface {
	Combine(...A) A
	// Identity(A) A
}

// takes two functions with embellished errors and returns their composition with combined errors
// this function circumvents the need to require functions to take errors as inputs
func ComposeEmbellished[A, B, C any, E Monoid[E]](
	f func(A) (B, E),
	g func(B) (C, E),
) func(A) (C, E) {
	return func(a A) (C, E) {
		b, e1 := f(a)
		c, e2 := g(b)
		return c, e1.Combine(e2)
	}
}

// using variadic to take Unit type for A
// would only be used at head of compose chain?
func ComposeEmbellished_[A, B, C any, E Monoid[E]](
	f func(...A) /*can be Unit*/ (B, E),
	g func(B) (C, E),
) func(...A) (C, E) {
	return func(a ...A) (C, E) {
		b, e1 := f(a...)
		c, e2 := g(b)
		return c, e1.Combine(e2)
	}
}

// embellishment E needs to be a monoid
// doesn't work b/c compiler thinks all input types need to implement Combinable
func ComposeEmbellished__[A, B, C, any, E Monoid[E]](
	f EmbellishedTransformFunc[A, B, E],
	g EmbellishedTransformFunc[B, C, E],
) EmbellishedTransformFunc[A, C, E] {
	return func(a A) (C, E) {
		b, e1 := f(a)
		c, e2 := g(b)
		return c, e1.Combine(e2)
	}
}

// example monoid type
type MonoidType string

func (m MonoidType) Combine(in ...MonoidType) (out MonoidType) {
	// return strings.Join(in..., "")
	for _, el := range in {
		out += " " + el
	}
	return
}
