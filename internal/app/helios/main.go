package helios

import (
	"encoding/json"

	"github.com/mikeydunn/helios/internal/pkg/instagram"
)

type users struct {
	Users []user `json:"users"`
}

type user struct {
	ID              int    `json:"id"`
	InstagramHandle string `json:"instagram_handle"`
}

// Init manages scrapping routine
func Init() {
	usersList := getUsers()

	for _, user := range usersList.Users {
		// verify user has instagram handle and scrape
		if user.InstagramHandle != "" {
			instagram.ScrapeUser(user.InstagramHandle)
		}
		// other social scraping ...
		// if user.twitterHandle !== "" {
		//  twitter.ScrapeUser(user.twitterHandle)
		// }
	}
}

func getUsers() users {
	mockUsersByte := []byte(`{"users":[{"id":0,"instagram_handle":"dallasdevs"},{"id":1,"instagram_handle":""}]}`)
	mockUsers := users{}

	json.Unmarshal(mockUsersByte, &mockUsers)
	return mockUsers
}
