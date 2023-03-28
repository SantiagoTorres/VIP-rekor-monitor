package main

import (
    "fmt"
	"database/sql"
    "github.com/SantiagoTorres/vip-rekor-monitor/pkg/monitoring" // TODO: how to import getlogsize
	_ "github.com/go-sql-driver/mysql"
	//"time"
)

func main() {

    db, err := sql.Open("mysql", "root@tcp(localhost:3306)/monitor")
    if err != nil {
        panic(err)
    }
    fmt.Println(db)
    logsize := monitoring.GetLogSize()
    stmt := "INSERT INTO tree_verify VALUES (?, ?, ?, ?, ?, ?)"
    result, err := db.Exec(stmt, 3, "2003-12-31 00:00:00",
        "lol", logsize, "lol2", 31)
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
    //fmt.Printf("LogInfo Ok\nTree Size: %v!\n", logsize)
}


// See "Important settings" section.
//db.SetConnMaxLifetime(time.Minute * 3)
//db.SetMaxOpenConns(10)
//db.SetMaxIdleConns(10)
