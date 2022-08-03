package gotool

import (
	"log"

	"github.com/robfig/cron"
)

// TenMinutesTask - 每日排程
func TenMinutesTask() {

	mi := "0,10,20,30,40,50"
	ho := "0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23"

	c := cron.New()
	c.AddFunc("0 "+mi+" "+ho+" * * *", func() {
		//	do something
	})
	c.Start()
	log.Printf("Initialize HourTask Success!\n")
	log.Printf("Initialize http://127.0.0.1:9090/ \n")

}
