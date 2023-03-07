package main

import (
    "fmt"
    "github.com/SantiagoTorres/vip-rekor-monitor/pkg/monitoring" // TODO: how to import getlogsize
)


func main() {
    fmt.Println("Hello world!")
    logsize := monitoring.GetLogSize()
    fmt.Printf("LogInfo Ok\nTree Size: %v!\n", logsize)
}
