package models

type Photo struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"not null"`
	URL    string `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"`
}
