package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
    //define the logger
    sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some program!")

    if err != nil {
    	log.Fatal(err)
    } else {
    	log.SetOutput(sysLog)
    }

    log.Panic(sysLog)
    fmt.Println("Will you see this?")
}