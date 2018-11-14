package main

import (
	"fmt"
)

func main() {
	hello()
}

func hello() {
	fmt.Println("Hello, WebAssembly!")
}

//func main() {
//	c := make(chan struct{}, 0)
//	add := func(i []js.Value) {
//		js.Global.Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
//	}
//	js.Global.Set("add", js.NewCallback(add))
//	<-c
//}

//func jsDom() {
//	c := make(chan struct{}, 0)
//	cb = js.NewCallback(func(args []js.Value) {
//		move := js.Global.Get("document").Call("getElementById", "body").Get("value").Int()
//		fmt.Println(move)
//	})
//	js.Global.Get("document").Call("getElementById", "myText").Call("addEventListener", "input", cb)
//	// The goal of the channel is to wait indefinitly
//	// Otherwise, the main function ends and the wasm modules stops
//	<-c
//}

