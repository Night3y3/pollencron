package main

import (
	"log"
	"time"

	"github.com/Night3y3/pollencron/allergy_api"
	"github.com/Night3y3/pollencron/utils"
	"github.com/robfig/cron"
)

func main() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	checkNilError(err)
	cronJob := cron.NewWithLocation(loc)

	cronJob.AddFunc("0 0 8 * * *", func() {
		log.Println("Every day at 8:00:00 AM")
		dailyAverageMessage, err := allergy_api.GetHourlyLoadData()
		checkNilError(err)

		log.Printf("1")

		historicalDataMessage, err := allergy_api.GetCurrentChartData()
		checkNilError(err)

		log.Printf("2")

		slackMessage := *dailyAverageMessage + "\n" + *historicalDataMessage

		log.Printf("3")

		err = utils.SendSlackMessage(slackMessage)
		checkNilError(err)

		log.Println("Message sent to Slack : ", slackMessage)
	})

	cronJob.Start()

	select {}
}

func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}	
}