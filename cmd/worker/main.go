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

	database.NewDb()
	//repository := database.CampaignRepository{Db: db}

	//for {
	//	campaigns, _ := repository.GetCampaignsToBeSent()

	//	for _, campaign := range campaigns {
	//		println(campaign.ID)
	//	}

	//}
}
