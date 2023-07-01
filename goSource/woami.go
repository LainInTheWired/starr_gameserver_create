package main

import (
        "fmt"
        "os"
)

func main() {
        if os.Geteuid() != 0 {
                fmt.Println("Must be run as root")
                os.Exit(1)
        }
        os.Exit(0)
}