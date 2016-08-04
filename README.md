# GoLang

## Html / template
Set up a nice way to work with templates.
* http://stackoverflow.com/questions/17206467/go-how-to-render-multiple-templates-in-golang

#### Neat set of Html5 Templates
* http://html5up.net/


## Cross - Compile
```bash 
env GOOS=linux GOARCH=386 go build -v -o branching *go
```

## Read input data
``` go
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
```

