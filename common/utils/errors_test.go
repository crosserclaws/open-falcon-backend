package utils

import (
	"log"
	"fmt"
	. "gopkg.in/check.v1"
)

type TestErrorsSuite struct{}

var _ = Suite(&TestErrorsSuite{})

// Tests the capture of panic to error object
func (suite *TestErrorsSuite) TestPanicToSimpleError(c *C) {
	sampleFunc := func() (err error) {
		defer PanicToSimpleError(&err)()

		panic("Sample Error 1")
	}

	c.Assert(sampleFunc(), NotNil)
}

func ExamplePanicToError() {
	sampleFunc := func() (err error) {
		defer PanicToError(
			&err,
			func(p interface{}) error {
				return fmt.Errorf("Customized: %v", p)
			},
		)()

		panic("Good Error!!")
	}

	err := sampleFunc()
	fmt.Println(err)

	// Output:
	// Customized: Good Error!!
}

func ExamplePanicToSimpleError() {
	sampleFunc := func() (err error) {
		defer PanicToSimpleError(&err)()

		panic("Novel Error!!")
	}

	err := sampleFunc()
	fmt.Println(err)

	// Output:
	// Novel Error!!
}

func ExamplePanicToSimpleErrorWrapper() {
	sampleFunc := func(n int) {
		log.Panicf("Value: %d", n)
	}

	err := PanicToSimpleErrorWrapper(
		func() { sampleFunc(918) },
	)()
	fmt.Println(err)

	// Output:
	// Value: 918
}
