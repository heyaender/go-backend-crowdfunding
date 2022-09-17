package databases

import (
	"fmt"
	"go-backend-crowdfunding/campaignimages"
	"go-backend-crowdfunding/campaigns"
	"go-backend-crowdfunding/transactions"
	"go-backend-crowdfunding/users"
)

// It will create a table for each of the structs in the models folder
func DBMigrate() {
	DB.AutoMigrate(&users.Users{}, &campaigns.Campaigns{}, &campaignimages.CampaignImages{}, &transactions.Transactions{})
	fmt.Println("DB Migration Success")
}
