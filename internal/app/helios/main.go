package helios

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // initialize mysql driver
	"github.com/mikeydunn/helios/internal/pkg/database"
	"github.com/mikeydunn/helios/internal/pkg/instagram"
)

// Init manages scrapping routine
func Init() {
	// Open database connection
	database.DBCon, _ = sql.Open("mysql", "root:@tcp(helios-mysql:3306)/helios")

	// Get users list
	usersList := selectUsers()

	// Iterater over users with scraping packages
	for _, user := range usersList {
		// verify user has instagram handle and scrape
		if user.InstagramHandle != "" {
			instagram.Init(user.ID, user.InstagramHandle)
		}
		// other social scraping ...
		// if user.twitterHandle !== "" {
		//  twitter.ScrapeUser(user.twitterHandle)
		// }
	}
}
