package entities

import "time"

// Book is the entity holds the structure of book
type Book struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Title     string     `gorm:"type:varchar(50);not null" json:"title"`
	Author    string     `gorm:"type:varchar(50);not null" json:"author"`
}
