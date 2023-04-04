package main

import (
    "github.com/SantiagoTorres/vip-rekor-monitor/pkg/monitoring" // TODO: how to import getlogsize
)


func main() {
	logInfoResp, err := monitoring.GetLogInfoResp()
	if err != nil {
		panic(err)
	}
	err = monitoring.InsertIntoDB(logInfoResp)
	if err != nil {
		panic(err)
	}
}
