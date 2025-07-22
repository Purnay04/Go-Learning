package main

import "fmt"

func main()  {

    s:= []struct{
        x int
        y int
    }{
        {1,2},
        {11,12},
        {21,20},
        {11,21},
    }

    fmt.Println(s)

}