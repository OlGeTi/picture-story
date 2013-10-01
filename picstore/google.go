package picstore

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

type gAlbumQuery struct {
	AlbumList []gAlbum `xml:"entry"`
}

type gAlbum struct {
	UserID  UserID  `xml:"gphoto:user"`
	ID      AlbumID `xml:"gphoto:id"`
	Title   string  `xml:"title"`
	Summary string  `xml:"summary"`
}

func GAlbums(userID string) (*[]Album, error) {

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

	var query gAlbumQuery

	xml.Unmarshal(body, &query)

	fmt.Println(query)

	return convertToAlbumList(&query.AlbumList), nil
}

func convertToAlbum(gAlbum gAlbum) Album {

	return Album{
		UserID:  gAlbum.UserID,
		AlbumID: gAlbum.ID,
		Title:   gAlbum.Title,
		Summary: gAlbum.Summary,
	}
}

func convertToAlbumList(gAlbums *[]gAlbum) *[]Album {

	albums := make([]Album, len(*gAlbums))

	for i, gAlbum := range *gAlbums {
		albums[i] = convertToAlbum(gAlbum)
	}

	return &albums
}
