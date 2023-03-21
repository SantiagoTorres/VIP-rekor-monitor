package monitoring

import (
    "github.com/sigstore/rekor/pkg/client"
)

// TODO: no error handling?!
// TODO: hardcoded parameter (URL)
func GetLogSize() int64 {
    rekorClient, _ := client.GetRekorClient("https://rekor.sigstore.dev")
    logInfo, _ := rekorClient.Tlog.GetLogInfo(nil)
    return *logInfo.Payload.TreeSize
}
