package helios

import (
	"github.com/mikeydunn/helios/internal/pkg/database"
)

// selectUsers : get all users in our database
func selectUsers() []user {
	var (
		id       int
		username string
		userList []user
	)

	// Execute the query
	rows, _ := database.DBCon.Query("SELECT account_id, username FROM accounts")
	defer rows.Close()

	// scan returned rows
	for rows.Next() {
		rows.Scan(&id, &username)
		userEl := user{id, username}
		userList = append(userList, userEl)
	}

	return userList
}
