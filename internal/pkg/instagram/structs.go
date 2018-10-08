package instagram

// Instagram : parent wrapper for instagram requests
type instagram struct {
	EntryData struct {
		PostPage    []postPage    `json:"PostPage"`
		ProfilePage []profilePage `json:"ProfilePage"`
	} `json:"entry_data"`
}

// PostPage : wrapper for post requests
type postPage struct {
	Graphql struct {
		ShortcodeMedia edgeMedia `json:"shortcode_media"`
	} `json:"graphql"`
}

// ProfilePage : wrapper for profile requests
type profilePage struct {
	Graphql struct {
		User edgeUser `json:"user"`
	} `json:"graphql"`
}

// EdgeUser : user object
type edgeUser struct {
	EdgeFollowedBy           edgeFollowedBy           `json:"edge_followed_by"`
	EdgeFollow               edgeFollow               `json:"edge_follow"`
	ProfilePicURLHd          string                   `json:"profile_pic_url_hd"`
	Username                 string                   `json:"username"`
	ConnectedFbPage          interface{}              `json:"connected_fb_page"`
	EdgeOwnerToTimelineMedia edgeOwnerToTimelineMedia `json:"edge_owner_to_timeline_media"`
}

// EdgeFollowedBy : followed by object
type edgeFollowedBy struct {
	Count int `json:"count"`
}

// EdgeFollowedBy : following object
type edgeFollow struct {
	Count int `json:"count"`
}

// EdgeOwnerToTimelineMedia : collection of recent posts
type edgeOwnerToTimelineMedia struct {
	Count int `json:"count"`
	Edges []struct {
		Node edgeMedia `json:"node"`
	} `json:"edges"`
}

// EdgeMedia : media object
type edgeMedia struct {
	DisplayURL           string               `json:"display_url"`
	EdgeMediaToComment   edgeMediaToComment   `json:"edge_media_to_comment"`
	EdgeMediaPreviewLike edgeMediaPreviewLike `json:"edge_media_preview_like"`
}

// EdgeMediaToComment : comment object
type edgeMediaToComment struct {
	Count int `json:"count"`
}

// EdgeMediaPreviewLike : preview object
type edgeMediaPreviewLike struct {
	Count int `json:"count"`
}
