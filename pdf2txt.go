package main

import (
        "fmt"
        "os"

        "code.sajari.com/docconv"
)

func main() {
        arg := os.Args
        f := arg[1]
        res, err := docconv.ConvertPath(f)
        if err != nil {
                fmt.Printf("error %s\n", err)
        }
        fmt.Println(res)
}
