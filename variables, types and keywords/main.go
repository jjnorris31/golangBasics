package main

import (
	"fmt"
)

func main(){
	// asignación clásica
	var a string
	a = "hola"
	// duck type
	b := 20
	// group declaration
	var (
		y bool
		z int
	)
	y = false
	z = 16
	// discard some value
	_, g := 21, 22

	fmt.Printf("%s %d %t %d %d \n", a, b, y, z, g)

	// strings are inmutable
	var s string
	s = "something"
	// s[0] = 'v' makes an error, instead
	q := []byte(s)
	q[0] = 'v'
	s = string(q)
	fmt.Println(s)
}
