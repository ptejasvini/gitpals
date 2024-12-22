package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ptejasvini/gitpals/service"
)

func main() {
	// Define the flag for the GitHub username
	username := flag.String("username", "", "GitHub UserName")
	flag.Parse()

	// Check if username is provided
	if *username == "" {
		log.Fatal("Please provide a valid username using -username flag")
	}

	// Fetch and display user profile information
	fmt.Println("Fetching User Profile", *username)
	profile, err := service.GetUserProfile(*username)
	if err != nil {
		log.Fatal(err)
	}

	// Print profile information
	fmt.Println("Avatar URL: ", profile.AvatarURL)
	fmt.Println("Followers: ", profile.Stats.Followers)
	fmt.Println("Following: ", profile.Stats.Following)
	fmt.Println("Public Repos: ", profile.Stats.PublicRepos)

	// Fetch and display user activity
	fmt.Printf("\nFetching activity for: %s\n", *username)
	activities, err := service.GetUserActivity(*username)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Activity", activities[0].Repo.Name)
	//Loop through activities and display them based on the type
	for _, activity := range activities {
		switch activity.Type {
		case "PushEvent":
			// Format and print push event details
			fmt.Printf("Pushed %d commits to %s\n", len(activity.Repo.Name), activity.Repo.Name)
		case "IssuesEvent":
			// Format and print issue event details
			fmt.Printf("Opened an issue in %s\n", activity.Repo.Name)
		case "WatchEvent":
			// Format and print watch/star event details
			fmt.Printf("Starred %s\n", activity.Repo.Name)
		default:
			// For other event types, print a generic message
			fmt.Printf("Event: %s in %s\n", activity.Type, activity.Repo.Name)
		}
	}
}
