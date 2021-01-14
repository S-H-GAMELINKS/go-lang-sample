package main

import (
        "fmt"
)

type X struct {
    Name string
}

func (x *X) hasName() bool {
    if x.Name != "" {
        return true
    }
    return false
}

func main() {
    x1 := X{Name: "hoge"}
    fmt.Println(x1.Name)
    x2 := X{}
    if x2.hasName() {
        fmt.Println(x2.hasName())
    }
}
