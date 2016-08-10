package main

import "fmt"

func main() {
    var a []uint32
    a = []uint32{}
    for i := 0; i<2; i++ {
        var b uint32
        fmt.Scanf("%v", &b)
        a = append(a, b)
    }
    fmt.Println("---------")
    fmt.Println(a)
}
