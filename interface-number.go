package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num interface{} = 20
	fmt.Println(reflect.TypeOf(num))
}
