package main

import "fmt"

func main()  {
    for i := 1; i <= 10; i ++{
        if i == 5{
            return
        }
        fmt.Println("hello",i)
    }
    fmt.Println("hello,world")
}