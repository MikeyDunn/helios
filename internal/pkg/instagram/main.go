package instagram

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type instagram struct {
	EntryData struct {
		PostPage    []postPage    `json:"PostPage"`
		ProfilePage []profilePage `json:"ProfilePage"`
	} `json:"entry_data"`
}

type postPage struct {
	Graphql struct {
		ShortcodeMedia edgeMedia `json:"shortcode_media"`
	} `json:"graphql"`
}

type profilePage struct {
	Graphql struct {
		User edgeUser `json:"user"`
	} `json:"graphql"`
}

type edgeUser struct {
	EdgeFollowedBy           edgeFollowedBy           `json:"edge_followed_by"`
	ProfilePicURLHd          string                   `json:"profile_pic_url_hd"`
	RequestedByViewer        bool                     `json:"requested_by_viewer"`
	Username                 string                   `json:"username"`
	ConnectedFbPage          interface{}              `json:"connected_fb_page"`
	EdgeOwnerToTimelineMedia edgeOwnerToTimelineMedia `json:"edge_owner_to_timeline_media"`
}

type edgeFollowedBy struct {
	Count int `json:"count"`
}

type edgeOwnerToTimelineMedia struct {
	Count int `json:"count"`
	Edges []struct {
		Node edgeMedia `json:"node"`
	} `json:"edges"`
}

type edgeMedia struct {
	DisplayURL           string               `json:"display_url"`
	EdgeMediaToComment   edgeMediaToComment   `json:"edge_media_to_comment"`
	EdgeMediaPreviewLike edgeMediaPreviewLike `json:"edge_media_preview_like"`
}

type edgeMediaToComment struct {
	Count int `json:"count"`
}

type edgeMediaPreviewLike struct {
	Count int `json:"count"`
}

// ScrapeUser requests instagram user page and store results
func ScrapeUser(username string) {
	url := fmt.Sprintf("http://instagram.com/%s", username)
	body := requestHTML(url)
	data := parseHTML(body)
	instagramUser := instagram{}

	json.Unmarshal(data, &instagramUser)
	storeData(instagramUser)
}

// ScrapeMedia requests instagram user page and store results
func ScrapeMedia(mediaID string) {
	url := fmt.Sprintf("http://instagram.com/p/%s", mediaID)
	body := requestHTML(url)
	data := parseHTML(body)
	instagramMedia := instagram{}

	json.Unmarshal(data, &instagramMedia)
	storeData(instagramMedia)
}

func requestHTML(url string) io.ReadCloser {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return response.Body
}

func parseHTML(body io.ReadCloser) []byte {
	defer body.Close()
	contents, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	htmlString := string(contents)
	scriptStringPartial := strings.Split(htmlString, "_sharedData = ")[1]
	scriptString := strings.Split(scriptStringPartial, ";</script>")[0]

	return []byte(scriptString)
}

func storeData(data instagram) {
	dataJSON, _ := json.Marshal(data)

	fmt.Println(string(dataJSON))
}
