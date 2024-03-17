package model

import "Book-GOAPI/database"

type Book struct {
	BookID int    `gorm:"primaryKey"` // gausah define di DB kalo udah pake gorm
	Title  string `json:"Title"`
	Author string `json:"Author"`
	ISBN   string `json:"ISBN"`
}

func (book *Book) StoreBookRecord() error {
	result := database.GlobalDB.Create(&book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
