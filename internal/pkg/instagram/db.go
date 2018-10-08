package instagram

import (
	"fmt"
	"log"
	"time"

	"github.com/mikeydunn/helios/internal/pkg/database"
)

func insertData(id int, data instagram) {
	// insertAccountFollower(id, data)
	insertAccountPostStats(id, data)
}

func insertAccountFollower(id int, data instagram) {
	user := data.EntryData.ProfilePage[0].Graphql.User
	accountID := id
	followerCount := user.EdgeFollowedBy.Count
	followingCount := user.EdgeFollow.Count
	postCount := user.EdgeOwnerToTimelineMedia.Count
	insert, _ := database.DBCon.Prepare(`
    INSERT INTO account_follower_counts
    (account_id, follower_count, following_count, post_count)
    VALUES (?, ?, ?, ?)`)
	insert.Exec(accountID, followerCount, followingCount, postCount)
}

func insertAccountPostStats(id int, data instagram) {
	accountID := id
	postID := 1
	postDate := time.Now().Format(time.RFC3339)
	commentCount := 1
	likeCount := 1
	insert, err := database.DBCon.Prepare(`
    INSERT INTO account_post_stats
    (account_id, post_id, post_date, comment_count, like_count)
    VALUES (?, ?, ?, ?, ?)
    ON DUPLICATE KEY
      UPDATE comment_count=?, like_count=?`)
	if err != nil {
		fmt.Println("insert")
		log.Fatal(err)
	}
	_, err2 := insert.Exec(accountID, postID, postDate, commentCount, likeCount, commentCount, likeCount)
	if err2 != nil {
		fmt.Println("execute")
		log.Fatal(err2)
	}
}
