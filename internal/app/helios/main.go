package helios

import (
	"github.com/mikeydunn/helios/internal/pkg/instagram"
)

// Init manages scrapping routine
func Init() {
	usersList := selectUsers()

	for _, user := range usersList.Users {
		// verify user has instagram handle and scrape
		if user.InstagramHandle != "" {
			instagram.Init(user.InstagramHandle)
		}
		// other social scraping ...
		// if user.twitterHandle !== "" {
		//  twitter.ScrapeUser(user.twitterHandle)
		// }
	}
}
