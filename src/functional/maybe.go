package functional

type MaybeType[A any] struct {
	Value *A
}

type Maybe[A any] interface {
	Just() MaybeType[A]
	Nothing() MaybeType[A]
}

func (m MaybeType[A]) Just() MaybeType[A] {
	return MaybeType[A]{
		Value: m.Value,
	}
}

func (m MaybeType[A]) Nothing() MaybeType[A] {
	return MaybeType[A]{
		Value: nil,
	}
}
