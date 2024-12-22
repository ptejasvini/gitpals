// model/user.go
package model

// GitHubUser represents a GitHub user profile.
type GitHubUser struct {
    Login     string           `json:"login"`     // GitHub username
    ID        int              `json:"id"`        // User ID
    AvatarURL string           `json:"avatar_url"`// URL to the user's avatar
    Links     GitHubUserLinks `json:"links"`      // User profile and URL links
    Stats     UserStats        `json:"stats"`      // User statistics (followers, public repos, etc.)
}

// GitHubUserLinks holds the links related to the user.
type GitHubUserLinks struct {
    URL string `json:"url"` // User's GitHub profile URL
}

// UserStats holds statistical data about the user's account.
type UserStats struct {
    PublicRepos int `json:"public_repos"` // Number of public repositories
    Followers   int `json:"followers"`    // Number of followers
    Following   int `json:"following"`    // Number of users they follow
}
