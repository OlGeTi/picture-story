package picstore

import (
	"net/url"
)

type UserID string
type AlbumID string
type PictureID string

type Album struct {
	UserID  UserID
	AlbumID AlbumID
	Title   string
	Summary string
}

type Picture struct {
	UserID    UserID
	AlbumID   AlbumID
	PictureID PictureID
	URL       url.URL
}
