package functional

import (
	"fmt"
)

// would need to define a function that composes errors to defined the following way
// type ErrorEmbellishedTransformFunc[A, B any, Error error] ComposableEmbellishedFunc[A, B, Error]

type ErrorEmbellishedTransformFunc[A, B any, Error error] func(A) (B, Error)

func ComposeErrorEmbellished[A, B, C any, E error](
	f func(A) (B, E), g func(B) (C, E),
) func(A) (C, error) {
	return func(a A) (C, error) {
		b, err1 := f(a)
		c, err2 := g(b)
		return c, CombineErrors(err1, err2)
	}
}

func (f ErrorEmbellishedTransformFunc[A, B, E]) EmbellishMap() func([]A) ([]B, error) {
	return func(a []A) (b []B, e error) {
		for _, _a := range a {
			_b, _e := f(_a)
			e = CombineErrors(e, _e)
			b = append(b, _b)
		}
		return
	}
}

// not composable
// func ComposeUnitErrorEmbellished[B, C any, E error](
// 	f func() (B, E), g func(B) (C, E),
// ) func() (C, error) {
// 	return func() (C, error) {
// 		b, err1 := f()
// 		c, err2 := g(b)
// 		return c, CombineErrors(err1, err2)
// 	}
// }

func CombineErrors(err1, err2 error) (err error) {
	if err1 == nil && err2 == nil {
		return nil
	}
	if err1 != nil && err != nil {
		return fmt.Errorf("%w %w", err1, err2)
	}
	if err1 != nil {
		return err1
	}
	return err2
}

//func CombineErrors(errs ...error) (err error) {}
