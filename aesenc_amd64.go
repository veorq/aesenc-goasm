package main

import "fmt"

func aesenc(k, s*byte)

// go build 
func main() {
    var k = make([]byte, 16)
    var s = make([]byte, 16)

    for i := range k {
        k[i] = byte(i)
        s[i] = k[i]
    }

    fmt.Printf("in:\t%v, %v\n", k, s)

    aesenc(&k[0], &s[0])

    fmt.Printf("out:\t%v, %v\n", k, s)
}

