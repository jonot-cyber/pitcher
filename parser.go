package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

func Parse(markdown io.Reader) {
    scanner := bufio.NewScanner(markdown)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
