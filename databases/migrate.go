package databases

import (
	"fmt"
	"go-backend-crowdfunding/schemas"
)

func DBMigrate() {
	DB.AutoMigrate(&schemas.Users{}, &schemas.Campaigns{}, &schemas.CampaignImages{}, &schemas.CampaignImages{}, &schemas.Transactions{})
	fmt.Println("DB Migration Success")
}
