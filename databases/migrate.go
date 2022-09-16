package databases

import (
	"fmt"
	"go-backend-crowdfunding/schemas"
)

func DBMigrate() {
	DB.AutoMigrate(&schemas.Users{})
	fmt.Println("DB Migration Success")
}
