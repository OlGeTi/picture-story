// api
package google

import (
	"encoding/xml"
	"fmt"
	"github.com/stvp/go-toml-config"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	GOOGLE_USER_URL  = config.String("google.user.url", "https://picasaweb.google.com/data/feed/api/user/%s")
	GOOGLE_ALBUM_URL = config.String("google.album.url", "https://picasaweb.google.com/data/feed/api/user/%s/albumid/%s")
)

type AlbumQuery struct {
	AlbumList []Album `xml:"entry"`
}

func (aq AlbumQuery) String() string {

	result := ""

	for _, album := range aq.AlbumList {
		result += fmt.Sprintf("%s\n\n", album.String())
	}

	return result
}

type Album struct {
	ID      string `xml:"gphoto:id"`
	Title   string `xml:"title"`
	Summary string `xml:"summary"`
}

func (a Album) String() string {
	return fmt.Sprintf("ID: %s\nTitle: %s\nSummary: %s", a.ID, a.Title, a.Summary)
}

func Albums(userID string) (*[]Album, error) {

	resp, err := http.Get(fmt.Sprintf(*GOOGLE_USER_URL, userID))

	if err != nil {
		log.Printf("Albums(%s): %s", userID, err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	var query AlbumQuery

	xml.Unmarshal(body, &query)

	fmt.Println(query)

	return &query.AlbumList, nil
}
