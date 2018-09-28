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

// Init : Start instagram scraping job
func Init(username string) {
	scrapeUser(username)
}

// ScrapeUser : requests instagram user page and store results
func scrapeUser(username string) {
	url := fmt.Sprintf("http://instagram.com/%s", username)
	body := requestHTML(url)
	data := parseHTML(body)
	instagramUser := instagram{}

	json.Unmarshal(data, &instagramUser)
	insertData(instagramUser)
}

// ScrapeMedia : requests instagram user page and store results
func scrapeMedia(mediaID string) {
	url := fmt.Sprintf("http://instagram.com/p/%s", mediaID)
	body := requestHTML(url)
	data := parseHTML(body)
	instagramMedia := instagram{}

	json.Unmarshal(data, &instagramMedia)
	insertData(instagramMedia)
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
