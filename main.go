package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
    flag.Parse()
    args := flag.Args()
    if len(args) != 1 {
        log.Fatalln("You need to select a file!")
    }

    f, err := os.Open(args[0])
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    markdown, err := io.ReadAll(f)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(markdown))
}
