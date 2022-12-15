package main

import (
	"flag"
	"log"
)

func main() {
    args := flag.Args()
    if len(args) != 1 {
        log.Fatalln("You need to select a file!")
    }
}
