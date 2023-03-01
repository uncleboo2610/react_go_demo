package models

type Post struct {
	PostID int `gorm:"primaryKey;autoIncrement:true"`
	Title  string
	Body   string
	UserID int
	User   User `gorm:"references:ID"`
}

type User struct {
	ID   int
	Code string
	Name string
}

//go run migration/migrate.go
