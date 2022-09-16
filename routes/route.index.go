package routes

import (
	"fmt"
	"go-backend-crowdfunding/configs"
	"log"
	"net/http"
)

func LaunchApp() {
	WebServer := fmt.Sprintf("Web Server is running on port %s", configs.WEB_PORT)
	log.Println(WebServer)
	log.Fatal(http.ListenAndServe(configs.WEB_PORT, nil))
}
