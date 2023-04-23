package php

import (
	"fmt"
)

func Echo(a ...any) {
	fmt.Println(a...)
}

func Var_dump(a any) {
	fmt.Printf("%+v\r\n", a)
}
