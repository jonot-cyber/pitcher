package main

import (
	"flag"
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

    Parse(f)
}
