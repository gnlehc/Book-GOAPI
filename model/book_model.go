package model

type Book struct {
	BookID int    `json:"BookID" gorm:"primaryKey"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
	ISBN   string `json:"ISBN"`
}
