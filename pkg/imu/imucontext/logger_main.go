package imucontext

import (
	"fmt"
	"maps"
)

func Mmain() {
	r := gett()

	fmt.Printf("%v\n", r)
}

type MyMap map[string]string

func gett() (result MyMap) {
	a := MyMap{}
	a["r"] = "1"
	maps.Copy(result, a)
	return
}
