package main

import "fmt"

type otaku int

var x otaku

func main() {

x = 52

fmt.Printf("otaku: %v, %T\n", x, x)

x = 42
fmt.Println(x)






}