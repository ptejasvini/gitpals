package service

import (
	"fmt"

	"github.com/ptejasvini/gitpals/model"
	"github.com/ptejasvini/gitpals/repository"
)

// GetUserProfile retrieves the profile data for a given GitHub username
func GetUserProfile(username string) (*model.GitHubUser, error) {
	// Call repository to fetch the profile
	profile, err := repository.GetUserProfile(username)
	if err != nil {
		return nil, fmt.Errorf("error fetching user profile: %v", err)
	}

	return profile, nil
}

// GetUserActivity retrieves the activity data for a given GitHub username
func GetUserActivity(username string) ([]model.GitHubEvent, error) {
	// Call repository to fetch user activity (events)
	activities, err := repository.GetUserActivity(username)
	if err != nil {
		return nil, fmt.Errorf("error fetching user activity: %v", err)
	}

	return activities, nil
}
