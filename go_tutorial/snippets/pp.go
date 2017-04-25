package main

import (
	"github.com/k0kubun/pp"
)

func main(){
	m := map[string]string{"foo": "bar", "hello": "world"}
	pp.Print(m)
}
