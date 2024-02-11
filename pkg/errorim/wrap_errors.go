package errorim

import (
	"fmt"

	"github.com/pkg/errors"
)

func WrapErrorsAndPrintCallstack() {
	err := fmt.Errorf("standalone %v", "Message")
	err1 := errors.New("Inner Message")
	err2 := errors.Wrap(err1, "Outer Message")
	fmt.Printf("%v", err)
	fmt.Println("")
	fmt.Printf("%v", err1)
	fmt.Println("")
	fmt.Printf("%v", err2)
	fmt.Println("")
	fmt.Println("------------------------------ 1")
	fmt.Println(err)
	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println("------------------------------ 2")
	fmt.Printf("%+v", err)
	fmt.Println("------------------------------ 3")
	fmt.Printf("%+v", err1)
	fmt.Println("------------------------------ 4")
	fmt.Printf("%+v", err2)
}
