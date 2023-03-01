package main

import (
    "fmt"
    "github.com/sigstore/rekor/pkg/client"
)

func main() {
    fmt.Println("Hello world!")
    rekorClient, _ := client.GetRekorClient("https://rekor.sigstore.dev")
    logInfo, _ := rekorClient.Tlog.GetLogInfo(nil)
    fmt.Printf("LogInfo Ok\nTree Size: %v!\n", *logInfo.Payload.TreeSize)
}
