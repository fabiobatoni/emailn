package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/infrastructure/mail"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Started worker")

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDb()
	repository := database.CampaignRepository{Db: db}
	campaignService := campaign.ServiceImp{
		Repository: &repository,
		SendMail: mail.SendMail,
	}

	for {
		campaigns, err := repository.GetCampaignsToBeSent()

		if err != nil {
			println(err.Error())
		}

		println("Amount of campaigns: ",len(campaigns))

		for _, campaign := range campaigns {
			campaignService.SendEmailAndUpdateStatus(&campaign)
			println("Campaign sent: ", campaign.ID)
		}

		time.Sleep(10 * time.Second)
	}
}
