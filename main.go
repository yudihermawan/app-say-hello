package main

import (
	"fmt"
	mod_say_hello "github.com/yudihermawan/mod-say-hello"
)

func main() {
	hello := mod_say_hello.SayHello()
	fmt.Println(hello)
}