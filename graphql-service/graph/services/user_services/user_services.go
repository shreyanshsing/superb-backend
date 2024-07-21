package userservices

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"spuerb.com/v0/graphql_service/graph/model"
)

type Response struct {
    Data model.User `json:"data"`
}

func getUrl() string {
	userServiceBaseUrl := os.Getenv("USER_SERVICE_BASE")
	userSerivcePort := os.Getenv("USER_SERVICE_PORT")
	urlSuffix := os.Getenv("URL_SUFFIX")

	userService := urlSuffix + userSerivcePort + userServiceBaseUrl

	return userService
}

func GetUserById(userId int) (*model.User, error) {
	log.Print("calling user service to get user by id")
	
	userService := getUrl()

	res, err := http.Get(userService + strconv.Itoa(userId))

	if err != nil {
		log.Printf("Error calling user service: %v", err)
		return nil, err
	}
	defer res.Body.Close()
	
	// Log the status code and the response body for better visibility
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}
	bodyString := string(bodyBytes)
	log.Printf("User service response status: %s, body: %s", res.Status, bodyString)
	
	var user model.User
	var response Response
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		log.Printf("Error unmarshalling user data: %v", err)
		return nil, err
	}

	user = response.Data
	return &user, nil
}