package functional

import (
	"fmt"
	"testing"
)

func TestCompose(t *testing.T) {

	// example embellished functions

	Add1 := func(input int) (float32, MonoidType) {
		output := float32(input) + 1
		return output, MonoidType(fmt.Sprintf("%d+1=%f", input, output))
	}

	DivideBy3 := func(input float32) (float64, MonoidType) {
		output := float64(input) / 3.0
		return output, MonoidType(fmt.Sprintf("%f/3=%f", input, output))
	}

	// Compose the embellished functions
	composedFunc := ComposeEmbellished(Add1, DivideBy3)

	// Use the composed function
	composeOnValue := 5
	valueOut, embellishmentChain := composedFunc(composeOnValue)
	fmt.Printf("composing on value: %v\n", composeOnValue)
	fmt.Printf("value out %v\n", valueOut)
	fmt.Printf("embellishment chain: %v\n", embellishmentChain)
}

func TestErrorCompose(t *testing.T) {

	ErrorAdd1 := func(input int) (float32, error) {
		output := float32(input) + 1
		return output, nil //fmt.Errorf("error when adding one")
	}

	ErrorDivideBy3 := func(input float32) (float64, error) {
		output := float64(input) / 3.0
		return output, fmt.Errorf("error when diving by three")
	}

	composedErrorFunc := ComposeErrorEmbellished(ErrorAdd1, ErrorDivideBy3)

	composeOnValue := 5
	valueOut, errorChain := composedErrorFunc(composeOnValue)
	fmt.Printf("error was: %v\n", errorChain)
	fmt.Printf("composing on value: %v\n", composeOnValue)
	fmt.Printf("value out %v\n", valueOut)
}

func TestMap(t *testing.T) {
	add1 := func(i int) int {
		return i + 1
	}
	testSlice := []int{1, 2, 3}
	fmt.Printf("expecting %v map add1 is %v\n", testSlice, Slice[int](testSlice).Map(add1))
}

func TestEmbellishMap(t *testing.T) {

	MakeCustomZeroSlice := func(size int) ([]int, error) {
		return make([]int, size), nil
	}

	ErrorAdd1 := func(input int) (int, error) {
		output := input + 1
		return output, nil //fmt.Errorf("error when adding one")
	}

	composedMappedErrorFunc := ComposeErrorEmbellished(MakeCustomZeroSlice, ErrorEmbellishedTransformFunc[int, int, error](ErrorAdd1).EmbellishMap())

	composeOnValue := 3
	valueOut, errorChain := composedMappedErrorFunc(composeOnValue)
	fmt.Printf("error was: %v\n", errorChain)
	fmt.Printf("composing on value: %v\n", composeOnValue)
	fmt.Printf("value out: %v\n", valueOut)

}
