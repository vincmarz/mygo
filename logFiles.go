package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
    programName := filepath.Base(os.Args[0])
    //define the logger
    sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7,programName)
    //set loca7 logging facility using the info logging level
    //err variable is returned
    if err != nil {
    	log.Fatal(err)
    } else {
    	log.SetOutput(sysLog)
    }
    log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")
    sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program!")
    if err != nil {
    	log.Fatal(err)
    } else {
    	log.SetOutput(sysLog)
    }
    log.Println("LOG_MAIL: Logging in Go!")
    fmt.Println("Will you see this?")
}