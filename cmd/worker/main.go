package main

import (
	"emailn/internal/infrastructure/database"
	"fmt"
	"log"

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
	campaigns, err := repository.GetCampaignsToBeSent()

	if err != nil {
		println(err.Error())
	}

	println(len(campaigns))

	for _, campaign := range campaigns {
		println(campaign.ID)
	}

	//for {
	//	campaigns, _ := repository.GetCampaignsToBeSent()

	//	for _, campaign := range campaigns {
	//		println(campaign.ID)
	//	}

	//}
}
