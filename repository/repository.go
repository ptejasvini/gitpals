package repository

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ptejasvini/gitpals/httpclient"
	"github.com/ptejasvini/gitpals/model"
)

func GetUserProfile(username string) (*model.GitHubUser, error) {

	getUserData, err := GetUser(username)
	if err != nil {
		return nil, err
	}
	fmt.Println("Raw User Profile Response:", string(getUserData))

	var user model.GitHubUser

	if err := json.Unmarshal(getUserData, &user); err != nil {
		log.Println("Error unmarshalling data : %v", err)
		return nil, err
	}
	return &user, nil
}

func GetUserActivity(username string) ([]model.GitHubEvent, error) {
	getUserData, err := GetUser(username)
	if err != nil {
		return nil, err
	}

	var events []model.GitHubEvent
	if err := json.Unmarshal(getUserData, &events); err != nil {
		log.Println("error unmarshalling the Event Data %v", err)
		return nil, err
	}
	return events, nil
}

func GetUser(username string) ([]byte, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	data, err := httpclient.Get(url)

	if err != nil {
		return nil, err
	}

	return data, nil
}
