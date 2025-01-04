package models

// album represents data about a record album.
// Album (A) is exported while album (a) not exported in Go.
type Album struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Artist    string  `json:"artist"`
	Price     float64 `json:"price"`
	Image     string  `json:"image"`
	IsDeleted bool    `json:"isdeleted"`
}
