package getvideo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Define API Parts
const apiKey = "<API_KEY_HERE>"
const channelId = "<CHANNEL_ID_HERE>"
const maxResults = "1"
const apiUrl = "https://www.googleapis.com/youtube/v3/search?key=" + apiKey + "&channelId=" + channelId + "&part=snippet,id&order=date" + "&maxResults=" + maxResults

// Define Structure of Json file
type Items struct {
	Items []Item `json:"items"`
}
type Item struct {
	Snippet Snippet `json:"snippet"`
	Id      Id      `json:"id"`
}
type Id struct {
	VideoId string `json:"videoId"`
}
type Snippet struct {
	Title     string    `json:"title"`
	Thumbnail Thumbnail `json:"thumbnails"`
}
type Thumbnail struct {
	High High `json:"high"`
}
type High struct {
	Url string `json:"url"`
}

func GetVideoThumb() string {
	//Get the high quality thumbnail URl from the JSON file
	fileContent, err := os.Open("latest-videos.json")

	if err != nil {
		GetNewVideos()
	}

	fmt.Println("Getting Thumbnail...")

	defer fileContent.Close()

	byteResult, _ := io.ReadAll(fileContent)
	var items Items
	json.Unmarshal(byteResult, &items)

	return items.Items[0].Snippet.Thumbnail.High.Url
}
func GetVideoTitle() string {
	//Get the title of the video from the JSON file
	fileContent, err := os.Open("latest-videos.json")

	if err != nil {
		GetNewVideos()
	}

	fmt.Println("Getting Title...")

	defer fileContent.Close()

	byteResult, _ := io.ReadAll(fileContent)
	var items Items
	json.Unmarshal(byteResult, &items)

	return items.Items[0].Snippet.Title
}
func GetVideoUrl() string {
	//Get the URL to the video from the JSON file
	fileContent, err := os.Open("latest-videos.json")

	if err != nil {
		GetNewVideos()
	}

	fmt.Println("Getting URL...")

	defer fileContent.Close()

	byteResult, _ := io.ReadAll(fileContent)
	var items Items
	json.Unmarshal(byteResult, &items)
	urlstructure := "https://www.youtube.com/watch?v="
	return urlstructure + items.Items[0].Id.VideoId
}
func GetNewVideos() {
	fullURLFile := apiUrl
	fileName := "latest-videos.json"

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	//Printout the file name and size
	fmt.Printf("File [%s] Size [%d]", fileName, size)
}
