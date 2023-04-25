package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func beep() {
	dt := time.Now()
	fmt.Println("Log line", dt.Format("2006-01-02 15:04:05"))
	f, err := os.OpenFile("C:\\Users\\dan\\Projects\\014-APR-2023\\003-go-service\\program_logs.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(dt.Format("2006-01-02 15:04:05") + "\n"); err != nil {
		log.Println(err)
	}
}

// var (
// 	beepFunc = syscall.MustLoadDLL("user32.dll").MustFindProc("MessageBeep")
// )

// func beep() {
// 	beepFunc.Call(0xffffffff)
// }
