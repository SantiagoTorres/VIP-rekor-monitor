package monitoring

import (
	"time"
	"fmt"
	"database/sql"
    "github.com/sigstore/rekor/pkg/client"
	"github.com/sigstore/rekor/pkg/generated/models"
	_ "github.com/go-sql-driver/mysql"

)

// TODO: no error handling?!
// TODO: hardcoded parameter (URL)
/*func GetLogSize()  {
    rekorClient, _ := client.GetRekorClient("https://rekor.sigstore.dev")
    logInfo, _ := rekorClient.Tlog.GetLogInfo(nil)
    return *logInfo.Payload.TreeSize
}*/
const timeFormat = "2010-01-01 20:01:00" 

func GetLogInfoResp() (*models.LogInfo, error) {

	rekorClient, err := client.GetRekorClient("https://rekor.sigstore.dev")
	if err != nil {
		return nil, err
	}
	logInfo, err := rekorClient.Tlog.GetLogInfo(nil)
	if err != nil {
		return nil, err
	}
	return logInfo.Payload, nil
}

func InsertIntoDB(logInfoResp *models.LogInfo) error {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/monitor")
	if err != nil {
		return err
	}

	 stmt := "INSERT INTO tree_verify VALUES (?, ?, ?, ?, ?)"
	 result, err := db.Exec(stmt, *logInfoResp.TreeID, time.Now().Format(timeFormat), 
							*logInfoResp.SignedTreeHead, *logInfoResp.TreeSize, *logInfoResp.RootHash)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
} 
